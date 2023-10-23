package repository

import (
	"context"
	"gin-starter/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"gin-starter/common/constant"
	"fmt"
)

type PinjamanRepository struct {
	db *gorm.DB
}

type PinjamanRepositoryUseCase interface {
	CreatePinjamanRequest(ctx context.Context, pinjaman *entity.Pinjaman) error
	GetPinjamanList(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Pinjaman, int64, error)
	GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error)
	HandledPinjaman(ctx context.Context, book *entity.Pinjaman) error
}

// NewBookRepository creates a new Book repository
func NewPinjamanRepository(db *gorm.DB) *PinjamanRepository {
	return &PinjamanRepository{db}

}

// CreateBook creates a new Book
func (pinjamanRepository *PinjamanRepository) CreatePinjamanRequest(ctx context.Context, pinjaman *entity.Pinjaman) error {
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Create(pinjaman).
		Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-CreatePinjamanRequest]")
	}
	return nil
}

// GetBooks returns a list of books
func (pinjamanRepository *PinjamanRepository) GetPinjamanList(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Pinjaman, int64, error) {
	var pinjaman []*entity.Pinjaman
	var total int64
	offsetPinjaman := ((page - 1) * limit)
	var gormDB= pinjamanRepository.db.
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

		if order != constant.Ascending && order != constant.Descending {
			order = constant.Descending
		}

	
		if sort == "" {
			sort = "created_at"
		}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))
	if err := gormDB.Find(&pinjaman).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[PinjamanRepository-GetPinjamanList]")
	}


	return pinjaman,total, nil
}

// GetBookByID returns a book by its ID
func (pinjamanRepository *PinjamanRepository) GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error) {
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


// // DeleteBookByID deletes a book by its ID
// func (pinjamanRepository *PinjamanRepository) DeleteBookByID(ctx context.Context, id uuid.UUID) error {
// 	if err := pinjamanRepository.db.
// 		WithContext(ctx).
// 		Where("id = ?", id).
// 		Delete(&entity.Pinjaman{}).
// 		Error; err != nil {
// 		return errors.Wrap(err, "[PinjamanRepository-DeleteBookByID]")
// 	}
// 	return nil
// }

func (pinjamanRepository *PinjamanRepository) GetPinjamanByUserID(ctx context.Context, userID string) (int, error) {
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

func (pinjamanRepository *PinjamanRepository) GetPinjamanByBukuID(ctx context.Context, bookID string) (int, error) {
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


func (pinjamanRepository *PinjamanRepository) HandledPinjaman(ctx context.Context, pinjaman *entity.Pinjaman) error {
	if err := pinjamanRepository.db.WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Where(`id = ?`, pinjaman.ID).
		Updates(pinjaman).Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-HandledPinjaman]")
	}
	return nil
}
