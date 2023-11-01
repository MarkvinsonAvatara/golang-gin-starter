package repository

import (
	"context"
	"fmt"
	"gin-starter/common/constant"
	"gin-starter/entity"
	// "log"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FinderPinjamanRepository struct {
	db *gorm.DB
}

type FinderPinjamanRepositoryUseCase interface {
	GetPinjamanList(ctx context.Context, search, filter, sort, order string, limit, page int) ([]*entity.PinjamanDetail, int64, error)
	GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.PinjamanDetail, error)
	GetAllList(ctx context.Context) (int64, int64, int64, int64,error) 
}

// NewBookRepository creates a new Book repository
func FinderNewPinjamanRepository(db *gorm.DB) *FinderPinjamanRepository {
	return &FinderPinjamanRepository{db}

}

// GetBooks returns a list of books
func (pinjamanRepository *FinderPinjamanRepository) GetPinjamanList(ctx context.Context, search, filter, sort, order string, limit, page int) ([]*entity.PinjamanDetail, int64, error) {
	var pinjaman []*entity.PinjamanDetail
	// var User []*entity.User
	// var Book []*entity.Book
	var total int64
	offsetPinjaman := ((page - 1) * limit)
	var gormDB = pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Select("public.pinjaman.*, public.user.name, public.user.dob , public.book.isbn, public.book.title, public.book.genre, public.book.author, public.book.publisher, public.book.edition ").
		Count(&total).
		Limit(limit).
		Offset(offsetPinjaman).
		Joins("LEFT join public.user on public.pinjaman.user_id = public.user.id inner join public.book on public.pinjaman.buku_id = public.book.id")

	if search != "" {
		gormDB = gormDB.
			Where("CAST(tglpinjam AS TEXT) ILIKE ?", "%"+search+"%").
			Or("CAST(tglkembali AS TEXT) ILIKE ?", "%"+search+"%")
	}

	if filter == strings.ToLower("pending") {
		gormDB = gormDB.Where("status = '0'")
	} else if filter == strings.ToLower("rejected") {
		gormDB = gormDB.Where("status = '2'")
	} else if filter == strings.ToLower("approved") {
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

	// for _, v := range pinjaman {
	// 	log.Println(v.BukuDetail.Title)
	// }

	return pinjaman, total, nil
}

// GetBookByID returns a book by its ID
func (pinjamanRepository *FinderPinjamanRepository) GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.PinjamanDetail, error) {
	pinjaman := new(entity.PinjamanDetail)
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Select("public.pinjaman.*, public.user.name, public.user.dob , public.book.isbn, public.book.title, public.book.genre, public.book.author, public.book.publisher, public.book.edition ").
		Joins("LEFT join public.user on public.pinjaman.user_id = public.user.id inner join public.book on public.pinjaman.buku_id = public.book.id").
		Where("public.pinjaman.id = ?", id).
		First(&pinjaman).
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

func (pinjamanRepository *FinderPinjamanRepository) GetAllList(ctx context.Context) (int64, int64, int64, int64,error) {
	// var pinjaman []*entity.Pinjaman
	var totalAvalaible int64
	var totalNotAvalaible int64
	var totalUser int64
	var totalUserPinjaman int64

	var _ = pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Select("public.book.*, public.pinjaman.status").
		Joins("LEFT JOIN public.pinjaman ON public.book.id = public.pinjaman.buku_id").
		Where("public.pinjaman.status IS NULL").
		Count(&totalAvalaible).
		Joins("LEFT join public.user on public.pinjaman.user_id = public.user.id inner join public.book on public.pinjaman.buku_id = public.book.id")

	var _ = pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Select("public.book.*, public.pinjaman.status").
		Joins("LEFT JOIN public.pinjaman ON public.book.id = public.pinjaman.buku_id").
		Where("public.pinjaman.status = 1 ").
		Count(&totalNotAvalaible).
		Joins("LEFT join public.user on public.pinjaman.user_id = public.user.id inner join public.book on public.pinjaman.buku_id = public.book.id")

	var _ = pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Count(&totalUser)

	var _ = pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Distinct("public.pinjaman.user_id").
		Count(&totalUserPinjaman)

	return totalAvalaible, totalNotAvalaible, totalUser, totalUserPinjaman, nil
}
