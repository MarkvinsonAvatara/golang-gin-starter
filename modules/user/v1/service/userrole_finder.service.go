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
	finderUserRoleRepo   repository.FinderUserRoleRepositoryUseCase
	// finderPinjamanRepo   repository.PinjamanRepositoryUseCase
	// roleRepo       repository.RoleRepositoryUseCase
	// permissionRepo repository.PermissionRepositoryUseCase
}

// UserFinderUseCase is a usecase for user
type UserRoleFinderUseCase interface {
	GetUserRoles(ctx context.Context, search, order, sort string, limit, offset int) ([]*entity.UserRole, int64, error)
	// GetUserRolesByIDs gets all user roles by ids
	GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error)
}

// NewUserFinder creates a new UserFinder
func NewUserRoleFinder(
	ufg config.Config,
	finderUserRoleRepo repository.FinderUserRoleRepositoryUseCase,
	// pinjamanRepo repository.PinjamanRepositoryUseCase,
	// permissionRepo repository.PermissionRepositoryUseCase,
) *UserRoleFinder {
	return &UserRoleFinder{
		ufg:            ufg,
		finderUserRoleRepo:   finderUserRoleRepo,
		// permissionRepo: permissionRepo,
	}
}

// GetUsers Roles gets all User Roles
func (uf *UserRoleFinder) GetUserRoles(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.UserRole, int64, error) {
	userroles, total, err := uf.finderUserRoleRepo.GetUserRoles(ctx, search, sort, order, limit, page)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return userroles, total, nil
}

func (uf *UserRoleFinder) GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error) {
	userRole, err := uf.finderUserRoleRepo.GetUserRoleByID(ctx, id)

	if err != nil {
		return userRole, errors.ErrInternalServerError.Error()
	}

	if userRole == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return userRole, nil
}


