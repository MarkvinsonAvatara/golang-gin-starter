package repository

import (
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	// "strings"
)

// BookRepository is an struct for Book repository
type DeleterBookRepository struct {
	db *gorm.DB
}

// BookRepositoryUseCase is an interface for Book repository use case
type DeleterBookRepositoryUseCase interface {
	DeleteBookByID(ctx context.Context, id uuid.UUID) error
}

// NewBookRepository creates a new Book repository
func DeleterNewBookRepository(db *gorm.DB) *DeleterBookRepository {
	return &DeleterBookRepository{db}

}


// DeleteBookByID deletes a book by its ID
func (bookRepository *DeleterBookRepository) DeleteBookByID(ctx context.Context, id uuid.UUID) error {
	if err := bookRepository.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.Book{}).
		Error; err != nil {
		return errors.Wrap(err, "[BookRepository-DeleteBookByID]")
	}
	return nil
}


