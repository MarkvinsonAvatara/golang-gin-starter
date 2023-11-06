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

type CreatePinjamanRepository struct {
	db *gorm.DB
}

type CreatePinjamanRepositoryUseCase interface {
	CreatePinjamanRequest(ctx context.Context, pinjaman *entity.Pinjaman) error
}

// NewBookRepository creates a new Book repository
func CreateNewPinjamanRepository(db *gorm.DB) *CreatePinjamanRepository {
	return &CreatePinjamanRepository{db}

}

func (pinjamanRepository *CreatePinjamanRepository) CreatePinjamanRequest(ctx context.Context, pinjaman *entity.Pinjaman) error {

	var user entity.User
	var book entity.Book
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", pinjaman.UserID).
		First(&user).Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-CreatePinjamanRequest] User not found or soft-deleted")
	}

	if err := pinjamanRepository.db.
		WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", pinjaman.BukuID).
		First(&book).Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-CreatePinjamanRequest] Book not found or soft-deleted")
	}

	if err := pinjamanRepository.db.
		WithContext(ctx).
		Create(pinjaman).
		Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-CreatePinjamanRequest]")
	}
	return nil
}
