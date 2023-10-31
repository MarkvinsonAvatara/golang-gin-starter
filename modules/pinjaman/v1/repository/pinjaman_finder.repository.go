package repository

import (
	"context"
	"fmt"
	"gin-starter/common/constant"
	"gin-starter/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type FinderPinjamanRepository struct {
	db *gorm.DB
}

type FinderPinjamanRepositoryUseCase interface {
	GetPinjamanList(ctx context.Context, search, filter,sort, order string, limit, page int) ([]*entity.Pinjaman, int64, error)
	GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error)
}

// NewBookRepository creates a new Book repository
func FinderNewPinjamanRepository(db *gorm.DB) *FinderPinjamanRepository {
	return &FinderPinjamanRepository{db}

}


// GetBooks returns a list of books
func (pinjamanRepository *FinderPinjamanRepository) GetPinjamanList(ctx context.Context, search, filter, sort, order string, limit, page int) ([]*entity.Pinjaman, int64, error) {
	var pinjaman []*entity.Pinjaman
	var total int64
	offsetPinjaman := ((page - 1) * limit)
	var gormDB = pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Count(&total).
		Limit(limit).
		Offset(offsetPinjaman)

	if search != "" {
		gormDB = gormDB.
			Where("CAST(tglpinjam AS TEXT) ILIKE ?", "%"+search+"%").
			Or("CAST(tglkembali AS TEXT) ILIKE ?", "%"+search+"%")
	}

	if filter == strings.ToLower("pending") {
		gormDB = gormDB.Where("status = '0'")
	} else if filter == strings.ToLower("rejected") {
		gormDB = gormDB.Where("status = '2'")
	}else if filter == strings.ToLower("approved") {
		gormDB = gormDB.Where("status = '1'")
	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "requested_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))
	if err := gormDB.Find(&pinjaman).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[PinjamanRepository-GetPinjamanList]")
	}

	return pinjaman, total, nil
}

// GetBookByID returns a book by its ID
func (pinjamanRepository *FinderPinjamanRepository) GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error) {
	pinjaman := new(entity.Pinjaman)
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Where("id = ?", id).
		First(pinjaman).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Book not found
		}
		return nil, errors.Wrap(err, "[BookRepository-GetBookByID]")
	}
	return pinjaman, nil
}

func (pinjamanRepository *FinderPinjamanRepository) GetPinjamanByUserID(ctx context.Context, userID string) (int, error) {
	models := make([]*entity.Pinjaman, 0)
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Where("REPLACE(lower(user_id), ' ', '') = ?", strings.ToLower(userID)).
		Find(&models).
		Error; err != nil {
		return 0, errors.Wrap(err, "[PinjamanRepository-GetPinjamanByUserID]")
	}
	return len(models), nil
}

func (pinjamanRepository *FinderPinjamanRepository) GetPinjamanByBukuID(ctx context.Context, bookID string) (int, error) {
	models := make([]*entity.Pinjaman, 0)
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Where("REPLACE(lower(buku_id), ' ', '') = ?", strings.ToLower(bookID)).
		Find(&models).
		Error; err != nil {
		return 0, errors.Wrap(err, "[PinjamanRepository-GetPinjamanByBukuID]")
	}
	return len(models), nil
}


