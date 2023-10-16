package repository

import (
	"context"
	"gin-starter/common/interfaces"
	"gin-starter/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// UserRoleRepository is a repository for user role
type PinjamanRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// UserRoleRepositoryUseCase is a use case for user role
type PinjamanRepositoryUseCase interface {
	// GetUser Role gets all user role
	GetPinjaman(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Pinjaman, int64, error)
	// FindByUserID is a method for finding user role by user id
	GetPinjamanID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error)
	// UpdateUseRole is a method for updating user role
	UpdatePinjaman(ctx context.Context, UserRole *entity.Pinjaman) error
	// Delete is a method for deleting user role
	DeletePinjaman(ctx context.Context, id uuid.UUID) error
	// CreateUserRole is a method for creating user role
	CreatePinjaman(ctx context.Context, role *entity.Pinjaman) error
}

// NewUserRoleRepository is a constructor for UserRoleRepository
func NewUserRoleRepository(db *gorm.DB, cache interfaces.Cacheable) *PinjamanRepository {
	return &PinjamanRepository{db, cache}
}


func (pinjamanRepository *PinjamanRepository) CreatePinjaman(ctx context.Context, book *entity.Pinjaman) error {
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Create(book).
		Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-CreatePinjaman]")
	}
	return nil
}

// GetBooks returns a list of books
func (pinjamanRepository *PinjamanRepository) GetPinjaman(ctx context.Context) ([]*entity.Pinjaman, error) {
	models := make([]*entity.Pinjaman, 0)
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[PinjamanRepository-GetPinjaman]")
	}
	return models, nil
}

func (pinjamanRepository *PinjamanRepository) GetPinjamanID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error) {
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
		return nil, errors.Wrap(err, "[PinjamanRepository-GetPinjamanID]")
	}
	return pinjaman, nil
}


// DeleteBookByID deletes a book by its ID
func (pinjamanRepository *PinjamanRepository) DeletePinjaman(ctx context.Context, id uuid.UUID) error {
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.Book{}).
		Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-DeletePinjaman]")
	}
	return nil
}

// func (pinjamanRepository *PinjamanRepository) FindBookByTitle(ctx context.Context, title string) (int, error) {
// 	models := make([]*entity.Book, 0)
// 	if err := bookRepository.db.
// 		WithContext(ctx).
// 		Model(&entity.Book{}).
// 		Where("REPLACE(lower(title), ' ', '') = ?", strings.ToLower(title)).
// 		Find(&models).
// 		Error; err != nil {
// 		return 0, errors.Wrap(err, "[BookRepository-FindBookByTitle]")
// 	}
// 	return len(models), nil
// }

func (pinjamanRepository *PinjamanRepository) UpdateBook(ctx context.Context, pinjaman *entity.Pinjaman) error {
	if err := pinjamanRepository.db.WithContext(ctx).
		Model(&entity.Pinjaman{}).
		Where(`id = ?`, pinjaman.ID).
		Updates(pinjaman).Error; err != nil {
		return errors.Wrap(err, "[BookBookRepositorysitory-UpdateBook]")
	}
	return nil
}