package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/modules/role/v1/repository"
	"github.com/google/uuid"
)

// UserDeleter is a service for user
type UserDeleter struct {
	cfg      config.Config
	userRoleRepo repository.UserRoleRepositoryUseCase

}

// DeleteUser implements UserDeleterUseCase.
func (*UserDeleter) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// UserDeleterUseCase is a use case for user
type UserDeleterUseCase interface {
	// DeleteUser deletes user role
	DeleteUserRole(ctx context.Context, id uuid.UUID) error
}

// NewUserDeleter creates a new UserDeleter
func NewUserDeleter(
	cfg config.Config,
	userRoleRepo repository.UserRoleRepositoryUseCase,

) *UserDeleter {
	return &UserDeleter{
		cfg:      cfg,
		userRoleRepo: userRoleRepo,
	}
}


// DeleteUserRole deletes user role
func (ud *UserDeleter) DeleteUserRole(ctx context.Context, id uuid.UUID) error {
	if err := ud.userRoleRepo.DeleteUserRole(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
