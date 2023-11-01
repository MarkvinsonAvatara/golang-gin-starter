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
	var count int64
	if err := pinjamanRepository.db.
		WithContext(ctx).
		Joins("INNER JOIN public.user ON public.pinjaman.user_id = public.user.id").
		Where("public.user.deleted_at IS NULL").
		Joins("INNER JOIN public.book ON public.pinjaman.buku_id = public.book.id").
		Where("public.book.deleted_at IS NULL").
		Model(&entity.Pinjaman{}).
		Count(&count).Error; err != nil {
		return errors.Wrap(err, "[PinjamanRepository-CreatePinjamanRequest]")
	}

	conditionsMet := count > 0
	if !conditionsMet {
		return errors.New("Conditions not met for creating Pinjaman")
	} else {
		if err := pinjamanRepository.db.
			WithContext(ctx).
			Create(pinjaman).
			Error; err != nil {
			return errors.Wrap(err, "[PinjamanRepository-CreatePinjamanRequest]")
		}
	}
	return nil
}
