package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/modules/user/v1/repository"

	"github.com/google/uuid"
)

// UserDeleter is a service for user
type UserRoleDeleter struct {
	cfg      config.Config
	deleteUserRoleRepo repository.DeleteUserRoleRepositoryUseCase

}



// UserDeleterUseCase is a use case for user
type UserRoleDeleterUseCase interface {
	DeleteUserRole(ctx context.Context, id uuid.UUID) error
}

// NewUserDeleter creates a new UserDeleter
func NewUserRoleDeleter(
	cfg config.Config,
	deleteUserRoleRepo repository.DeleteUserRoleRepositoryUseCase,

) *UserRoleDeleter {
	return &UserRoleDeleter{
		cfg:      cfg,
		deleteUserRoleRepo: deleteUserRoleRepo,
	}
}

// DeleteUserRole deletes user role
func (ud *UserRoleDeleter) DeleteUserRole(ctx context.Context, id uuid.UUID) error {
	if err := ud.deleteUserRoleRepo.DeleteUserRole(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
