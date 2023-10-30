package repository

import (
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/entity"
	// "github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	// "strings"
)

// BookRepository is an struct for Book repository
type UpdaterBookRepository struct {
	db *gorm.DB
}

// BookRepositoryUseCase is an interface for Book repository use case
type UpdaterBookRepositoryUseCase interface {
	// UpdateBook updates a book
	UpdateBook(ctx context.Context, book *entity.Book) error
}

// NewBookRepository creates a new Book repository
func UpdaterNewBookRepository(db *gorm.DB) *UpdaterBookRepository {
	return &UpdaterBookRepository{db}

}

func (bookRepository *UpdaterBookRepository) UpdateBook(ctx context.Context, book *entity.Book) error {
	if err := bookRepository.db.WithContext(ctx).
		Model(&entity.Book{}).
		Where(`id = ?`, book.ID).
		Updates(book).Error; err != nil {
		return errors.Wrap(err, "[BookBookRepositorysitory-UpdateBook]")
	}
	return nil
}
