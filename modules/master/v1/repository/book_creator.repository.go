package repository

import (
	"context"
	// "gin-starter/common/constant"
	"gin-starter/entity"
	// "github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// BookRepository is a struct for Book repository
type CreateBookRepository struct {
	db *gorm.DB
}

// BookRepositoryUseCase is an interface for Book repository use case
type CreateBookRepositoryUseCase interface {
	// CreateBook creates a new Book
	CreateBook(ctx context.Context, book *entity.Book) error
}

// NewBookRepository creates a new Book repository
func CreateNewBookRepository(db *gorm.DB,) *CreateBookRepository {
	return &CreateBookRepository{db}

}

// CreateBook creates a new Book
func (bookRepository *CreateBookRepository) CreateBook(ctx context.Context, book *entity.Book) error {
	if err := bookRepository.db.
		WithContext(ctx).
		Create(book).
		Error; err != nil {
		return errors.Wrap(err, "[BookRepository-CreateBook]")
	}
	return nil
}
