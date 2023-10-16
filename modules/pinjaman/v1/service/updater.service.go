package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/role/v1/repository"
)

// UserUpdater is a struct that contains the dependencies of UserUpdater
type UserUpdater struct {
	cfg            config.Config
	userRoleRepo   repository.UserRoleRepositoryUseCase
}

// UserUpdaterUseCase is a struct that contains the dependencies of UserUpdaterUseCase
type UserUpdaterUseCase interface {
	// UpdateUserRoles updates user roles
	UpdateUserRoles(ctx context.Context, userRole *entity.UserRole) error

}

// NewUserUpdater is a function that creates a new UserUpdater
func NewUserRoleDeleter(
	cfg config.Config,
	userRoleRepo repository.UserRoleRepositoryUseCase,
) *UserUpdater {
	return &UserUpdater{
		cfg:            cfg,
		userRoleRepo:   userRoleRepo,
	}
}

//UpdateUserRoles updates user roles
func (uu *UserUpdater) UpdateUserRoles(ctx context.Context, userRole *entity.UserRole) error {
	if err := uu.userRoleRepo.UpdateUserRole(ctx, userRole); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}


