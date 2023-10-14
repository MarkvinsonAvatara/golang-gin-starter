package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/user/v1/repository"

	"github.com/google/uuid"
)

// UserFinder is a service for user
type UserRoleFinder struct {
	ufg            config.Config
	userRoleRepo   repository.UserRoleRepositoryUseCase
}

// UserFinderUseCase is a usecase for user
type UserFinderUseCase interface {
	// GetUserRoles gets all user roles
	GetUserRoles(ctx context.Context, query, order, sort string, limit, offset int) ([]*entity.UserRole, int64, error)
	// GetUserRolesByIDs gets all user roles by ids
	GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error)
}

// NewUserFinder creates a new UserFinder
func NewUserFinder(
	ufg config.Config,
	userRoleRepo repository.UserRoleRepositoryUseCase,
) *UserRoleFinder {
	return &UserRoleFinder{
		ufg:            ufg,
		userRoleRepo:   userRoleRepo,
	}
}


// GetUsers Roles gets all User Roles
func (uf *UserRoleFinder) GetUserRoles(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.UserRole, int64, error) {
	userroles, total, err := uf.userRoleRepo.GetUserRoles(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return userroles, total, nil
}

func (uf *UserRoleFinder) GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error) {
	userRole, err := uf.userRoleRepo.GetUserRoleByID(ctx, id)

	if err != nil {
		return userRole, errors.ErrInternalServerError.Error()
	}

	if userRole == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return userRole, nil
}