package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/pinjaman/v1/repository"

	"github.com/google/uuid"
)

// UserFinder is a service for user
type PinjamanFinder struct {
	ufg            config.Config
	pinjamanRepo   repository.FinderPinjamanRepositoryUseCase
}

// UserFinderUseCase is a usecase for user
type FinderPinjamanFinderUseCase interface {
	// GetPinjamanList gets all pinjaman
	GetPinjamanList(ctx context.Context, search, filter, order, sort string, limit, page int) ([]*entity.Pinjaman, int64, error)
	// GetPinjamanByID gets a pinjaman by ID
	GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error)
}

// NewUserFinder creates a new UserFinder
func NewPinjamanFinder(
	ufg config.Config,
	pinjamanRepo repository.FinderPinjamanRepositoryUseCase,
) *PinjamanFinder {
	return &PinjamanFinder{
		ufg:            ufg,
		pinjamanRepo:   pinjamanRepo,
	}
}

// GetPinjamanList gets all pinjaman
func (uf *PinjamanFinder) GetPinjamanList(ctx context.Context, search, filter, sort, order string, limit, page int) ([]*entity.Pinjaman, int64, error) {
	pinjaman, total, err := uf.pinjamanRepo.GetPinjamanList(ctx, search, filter, sort, order, limit, page)
	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}
	return pinjaman, total, nil

}

// GetPinjamanByID gets a pinjaman by ID
func (uf *PinjamanFinder) GetPinjamanByID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error) {
	pinjaman, err := uf.pinjamanRepo.GetPinjamanByID(ctx, id)

	if err != nil {
		return pinjaman, errors.ErrInternalServerError.Error()
	}

	if pinjaman == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return pinjaman, nil
}
