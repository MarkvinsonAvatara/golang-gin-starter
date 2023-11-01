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

// BookRepository is a struct for Book repository
type FinderBookRepository struct {
	db *gorm.DB
}

// BookRepositoryUseCase is an interface for Book repository use case
type FinderBookRepositoryUseCase interface {
	// GetBooks returns a list of books
	GetBooks(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error)
	// GetBookByID returns a book by its ID
	GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
	// FindBookByTitle returns a book by its title
	FindBookByTitle(ctx context.Context, title string) (int, error)
	GetBookAvalaibily(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error)
}

// NewBookRepository creates a new Book repository
func FinderNewBookRepository(db *gorm.DB) *FinderBookRepository {
	return &FinderBookRepository{db}

}

// GetBooks returns a list of books
func (bookRepository *FinderBookRepository) GetBooks(ctx context.Context,search, sort, order string, limit, page int) ([]*entity.Book, int64, error) {
	var buku []*entity.Book
	var total int64
	offsetBook := ((page - 1) * limit)
	var gormDB = bookRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Select("public.book.*, public.pinjaman.status").
		Joins("LEFT JOIN public.pinjaman ON public.book.id = public.pinjaman.buku_id").
		Count(&total).
		Limit(limit).
		Offset(offsetBook)

	if search != "" {
		gormDB = gormDB.
			Where("CAST(isbn AS TEXT)ILIKE ?", "%"+search+"%").
			Or("title ILIKE ?", "%"+search+"%").
			Or("genre ILIKE ?", "%"+search+"%").
			Or("author ILIKE ?", "%"+search+"%").
			Or("publisher ILIKE ?", "%"+search+"%").
			Or("CAST(edition AS TEXT) ILIKE ?", "%"+search+"%")

	}

	// if filter!= "" {
	// 	gormDB = gormDB.Joins("LEFT JOIN public.pinjaman ON public.book.id = public.pinjaman.buku_id").
	// 	Where("public.pinjaman.status IS NULL").
	// 	Or ("public.pinjaman.status = ?", filter)
	// }

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&buku).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[BookRepository-GetBooks]")
	}

	return buku, total, nil
}

// GetBookByID returns a book by its ID
func (bookRepository *FinderBookRepository) GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
	book := new(entity.Book)
	if err := bookRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Where("id = ?", id).
		First(book).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Book not found
		}
		return nil, errors.Wrap(err, "[BookRepository-GetBookByID]")
	}
	return book, nil
}

func (bookRepository *FinderBookRepository) FindBookByTitle(ctx context.Context, title string) (int, error) {
	models := make([]*entity.Book, 0)
	if err := bookRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Where("REPLACE(lower(title), ' ', '') = ?", strings.ToLower(title)).
		Find(&models).
		Error; err != nil {
		return 0, errors.Wrap(err, "[BookRepository-FindBookByTitle]")
	}
	return len(models), nil
}

func (bookRepository *FinderBookRepository) GetBookAvalaibily(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error) {
	var buku []*entity.Book
	var total int64
	offsetBook := ((page - 1) * limit)
	var gormDB = bookRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Limit(limit).
		Select("public.book.*, public.pinjaman.status").
		Joins("LEFT JOIN public.pinjaman ON public.book.id = public.pinjaman.buku_id").
		Where("public.pinjaman.status IS NULL").
		Count(&total).
		Offset(offsetBook)



	if search != "" {
		gormDB = gormDB.
			Where("CAST(isbn AS TEXT)ILIKE ?", "%"+search+"%").
			Or("title ILIKE ?", "%"+search+"%").
			Or("genre ILIKE ?", "%"+search+"%").
			Or("author ILIKE ?", "%"+search+"%").
			Or("publisher ILIKE ?", "%"+search+"%").
			Or("CAST(edition AS TEXT) ILIKE ?", "%"+search+"%")

	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&buku).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[BookRepository-GetBooks]")
	}

	return buku, total, nil
}