package service

import (
	// "bytes"
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/user/v1/repository"
	// "gin-starter/utils"
	// "github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
	// "html/template"
	// "log"
)

// UserUpdater is a struct that contains the dependencies of UserUpdater
type UserRoleUpdater struct {
	cfg            config.Config
	updaterUserRoleRepo   repository.UpdateUserRoleRepositoryUseCase
	// roleRepo       repository.RoleRepositoryUseCase
	// permissionRepo repository.PermissionRepositoryUseCase
	// pinjamanRepo   repository.PinjamanRepositoryUseCase
}

// UserUpdaterUseCase is a struct that contains the dependencies of UserUpdaterUseCase
type UserRoleUpdaterUseCase interface {
	UpdateUserRoles(ctx context.Context, userRole *entity.UserRole) error
}

// NewUserUpdater is a function that creates a new UserUpdater
func NewUserRoleUpdater(
	cfg config.Config,
	updaterUserRoleRepo repository.UpdateUserRoleRepositoryUseCase,
) *UserRoleUpdater {
	return &UserRoleUpdater{
		cfg:            cfg,
		updaterUserRoleRepo:   updaterUserRoleRepo,
	}
}



// UpdateUserRoles updates user roles
func (uu *UserRoleUpdater) UpdateUserRoles(ctx context.Context, userRole *entity.UserRole) error {
	if err := uu.updaterUserRoleRepo.UpdateUserRole(ctx, userRole); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}



