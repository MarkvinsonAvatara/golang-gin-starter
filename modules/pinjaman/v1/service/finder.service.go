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
	pinjamanRepo   repository.PinjamanRepositoryUseCase
}

// UserFinderUseCase is a usecase for user
type UserFinderUseCase interface {
	// GetUserRoles gets all user roles
	GetPinjaman(ctx context.Context, query, order, sort string, limit, offset int) ([]*entity.Pinjaman, int64, error)
	// GetUserRolesByIDs gets all user roles by ids
	GetPinjamanID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error)
}

// NewUserFinder creates a new UserFinder
func NewUserFinder(
	ufg config.Config,
	pinjamanRepo repository.PinjamanRepositoryUseCase,
) *PinjamanFinder {
	return &PinjamanFinder{
		ufg:            ufg,
		pinjamanRepo:   pinjamanRepo,
	}
}


// GetUsers Roles gets all User Roles
func (PinjamanFinder *PinjamanFinder) GetPinjaman(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Pinjaman, int64, error) {
	pinjaman, total, err := PinjamanFinder.pinjamanRepo.GetPinjaman(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return pinjaman, total, nil
}

func (PinjamanFinder *PinjamanFinder) GetPinjamanID(ctx context.Context, id uuid.UUID) (*entity.Pinjaman, error) {
	userRole, err := PinjamanFinder.pinjamanRepo.GetPinjamanID(ctx, id)

	if err != nil {
		return userRole, errors.ErrInternalServerError.Error()
	}

	if userRole == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return userRole, nil
}