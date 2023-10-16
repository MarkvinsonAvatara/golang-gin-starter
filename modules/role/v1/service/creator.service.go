package service

import (
	"context"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/role/v1/repository"
	"github.com/google/uuid"
)

// UserCreator is a struct that contains all the dependencies for the User creator
type UserCreator struct {
	cfg          config.Config
	userRoleRepo repository.UserRoleRepositoryUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// UserCreatorUseCase is a use case for the User creator
type UserCreatorUseCase interface {
	CreateUserRole(ctx context.Context, name string, description string, createdBy string) (*entity.UserRole, error)
}

// NewUserCreator is a constructor for the User creator
func NewUserRoleCreator(
	cfg config.Config,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserCreator {
	return &UserCreator{
		cfg:          cfg,
		userRoleRepo: userRoleRepo,
		cloudStorage: cloudStorage,
	}
}

// CreateRole creates a User role
func (uc *UserCreator) CreateUserRole(ctx context.Context, name string, description string, createdBy string) (*entity.UserRole, error) {
	role := entity.NewUserRole(uuid.New(), name, description, createdBy)
	if err := uc.userRoleRepo.CreateUserRole(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}
