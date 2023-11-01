package service

import (
	"context"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/entity"
	// notificationService "gin-starter/modules/notification/v1/service"
	"gin-starter/modules/user/v1/repository"

	"github.com/google/uuid"
)

// UserCreator is a struct that contains all the dependencies for the User creator
type UserRoleCreator struct {
	cfg            config.Config
	createUserRoleRepo   repository.CreateUserRoleRepositoryUseCase
	// roleRepo       repository.RoleRepositoryUseCase
	// permissionRepo repository.PermissionRepositoryUseCase
	// pinjamanRepo  repository.PinjamanRepositoryUseCase
	cloudStorage   interfaces.CloudStorageUseCase
}

// UserCreatorUseCase is a use case for the User creator
type UserRoleCreatorUseCase interface {
	// CreateRole creates a role
	CreateUserRole(ctx context.Context, name string, description string,createdBy string) (*entity.UserRole, error)
	// Create Pinjaman Request

}


// NewUserCreator is a constructor for the User creator
func NewUserRoleCreator(
	cfg config.Config,
	createUserRoleRepo repository.CreateUserRoleRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserRoleCreator {
	return &UserRoleCreator{
		cfg:            cfg,
		createUserRoleRepo:   createUserRoleRepo,
		// roleRepo:       roleRepo,
		// pinjamanRepo:  pinjamanRepo,
		// permissionRepo: permissionRepo,
		cloudStorage:   cloudStorage,
	}
}

func (uc *UserRoleCreator) CreateUserRole(ctx context.Context, name string,description string,  createdBy string) (*entity.UserRole, error) {
	role := entity.NewUserRole(uuid.New(), name, description, createdBy)
	if err := uc.createUserRoleRepo.CreateUserRole(ctx, role,); err != nil {
		return nil, err
	}

	return role, nil
}
