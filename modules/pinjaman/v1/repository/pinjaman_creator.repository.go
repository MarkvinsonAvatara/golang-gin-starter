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

// CreateBook creates a new Book
func (pinjamanRepository *CreatePinjamanRepository) CreatePinjamanRequest(ctx context.Context, pinjaman *entity.Pinjaman) error {
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Create(pinjaman).
		Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-CreatePinjamanRequest]")
	}
	return nil
}

