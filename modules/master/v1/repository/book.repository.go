package repository

import (
	"context"
	"gin-starter/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

// BookRepository is an struct for Book repository
type BookRepository struct {
	db *gorm.DB
}

// BookRepositoryUseCase is an interface for Book repository use case
type BookRepositoryUseCase interface {
	// FindAll returns all Books
	FindAll(ctx context.Context) ([]*entity.Book, error)
	// CreateBook creates a new Book
	CreateBook(ctx context.Context, book *entity.Book) error
	// GetBooks returns a list of books
	GetBooks(ctx context.Context) ([]*entity.Book, error)
	// GetBookByID returns a book by its ID
	GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
	DeleteBookByID(ctx context.Context, id uuid.UUID) error
	// FindBookByTitle returns a book by its title
	FindBookByTitle(ctx context.Context, title string) (int, error)
	// UpdateBook updates a book
	UpdateBook(ctx context.Context, book *entity.Book) error
}

// NewBookRepository creates a new Book repository
func NewBookRepository (db *gorm.DB) *BookRepository{
	return &BookRepository{db}
	
}

// CreateBook creates a new Book
func (bookRepository *BookRepository) CreateBook(ctx context.Context, book *entity.Book) error {
	if err := bookRepository.db.
		WithContext(ctx).
		Create(book).
		Error; err != nil {
		return errors.Wrap(err, "[BookRepository-CreateBook]")
	}
	return nil
}

// GetBooks returns a list of books
func (bookRepository *BookRepository) GetBooks(ctx context.Context) ([]*entity.Book, error) {
	models := make([]*entity.Book, 0)
	if err := bookRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[BookRepository-GetBooks]")
	}
	return models, nil
}

// GetBookByID returns a book by its ID
func (bookRepository *BookRepository) GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
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

// FindAll returns all Books
func (bookRepository *BookRepository) FindAll(ctx context.Context) ([]*entity.Book, error) {
	models := make([]*entity.Book, 0)
	if err := bookRepository.db.
		WithContext(ctx).
		Model(&entity.Book{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[BookRepository-FindAll]")
	}
	return models, nil
}

// DeleteBookByID deletes a book by its ID
func (bookRepository *BookRepository) DeleteBookByID(ctx context.Context, id uuid.UUID) error {
	if err := bookRepository.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.Book{}).
		Error; err != nil {
		return errors.Wrap(err, "[BookRepository-DeleteBookByID]")
	}
	return nil
}

func (bookRepository *BookRepository) FindBookByTitle(ctx context.Context, title string) (int, error) {
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

func (bookRepository *BookRepository) UpdateBook(ctx context.Context, book *entity.Book) error {
	if err := bookRepository.db.WithContext(ctx).
		Model(&entity.Book{}).
		Where(`id = ?`, book.ID).
		Updates(book).Error; err != nil {
		return errors.Wrap(err, "[BookBookRepositorysitory-UpdateBook]")
	}
	return nil
}
