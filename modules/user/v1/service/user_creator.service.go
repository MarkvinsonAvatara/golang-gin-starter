package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/entity"
	// notificationService "gin-starter/modules/notification/v1/service"
	"gin-starter/modules/user/v1/repository"
	"gin-starter/utils"
	"time"

	"github.com/google/uuid"
)

// UserCreator is a struct that contains all the dependencies for the User creator
type UserCreator struct {
	cfg            config.Config
	userRepo       repository.CreateUserRepositoryUseCase
	// userRoleRepo   repository.CreateUserRoleRepositoryUseCase
	// roleRepo       repository.RoleRepositoryUseCase
	// permissionRepo repository.PermissionRepositoryUseCase
	// pinjamanRepo  repository.PinjamanRepositoryUseCase
	cloudStorage   interfaces.CloudStorageUseCase
}

// UserCreatorUseCase is a use case for the User creator
type UserCreatorUseCase interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, name string, email string, password string, roleid string, dob time.Time) (*entity.User, error)
	// CreateAdmin creates a new admin
	CreateAdmin(ctx context.Context, name string, email string, password string, roleid string, dob time.Time, role string) (*entity.User, error)
	// CreatePermission creates a permission
	// CreatePermission(ctx context.Context, name, label string) (*entity.Permission, error)
	// CreateRole creates a role
	// CreateUserRole(ctx context.Context, name string, description string,createdBy string) (*entity.UserRole, error)
	// Create Pinjaman Request
	// CreatePinjamanRequest(ctx context.Context, userid string, bookid string, tglpinjaman time.Time, tglkembali time.Time, requestedBy string)(*entity.Pinjaman, error)

}


// NewUserCreator is a constructor for the User creator
func NewUserCreator(
	cfg config.Config,
	userRepo repository.CreateUserRepositoryUseCase,
	// userRoleRepo repository.CreateUserRoleRepositoryUseCase,
	// roleRepo repository.RoleRepositoryUseCase,
	// pinjamanRepo repository.PinjamanRepositoryUseCase,
	// permissionRepo repository.PermissionRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserCreator {
	return &UserCreator{
		cfg:            cfg,
		userRepo:       userRepo,
		// userRoleRepo:   userRoleRepo,
		// roleRepo:       roleRepo,
		// pinjamanRepo:  pinjamanRepo,
		// permissionRepo: permissionRepo,
		cloudStorage:   cloudStorage,
	}
}

// CreateUser creates a new user
func (uc *UserCreator) CreateUser(ctx context.Context, name string, email string, password string, roleid string, dob time.Time) (*entity.User, error) {
	user := entity.NewUser(
		uuid.New(),
		name,
		email,
		password,
		roleid,
		utils.TimeToNullTime(dob),
		"System",
	)

	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return user, nil
}

// CreateAdmin creates a new admin
func (uc *UserCreator) CreateAdmin(ctx context.Context, name string, email string, password string, roleid string, dob time.Time, role_name string) (*entity.User, error) {
	userID := uuid.New()
	user := entity.NewUser(
		userID,
		name,
		email,
		password,
		roleid,
		utils.TimeToNullTime(dob),
		"Super Admin",
	)

	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	// userRole := entity.NewUserRole(uuid.New(), userID, roleID, "system")

// 	if err := uc.userRoleRepo.CreateOrUpdate(ctx, userRole); err != nil {
// 		return nil, errors.ErrInternalServerError.Error()
// 	}

	return user, nil
}

// // CreatePermission creates a permission
//  func (uc *UserCreator) CreatePermission(ctx context.Context, name, label string) (*entity.Permission, error) {
//  	permission := entity.NewPermission(
//  		uuid.New(),
//  		name,
//  		label,
//  		"system",
//  	)
// 
// 	if err := uc.permissionRepo.Create(ctx, permission); err != nil {
// 		return nil, err
// 	}
// 
// 	return permission, nil
// }

// CreateRole creates a role
// func (uc *UserCreator) CreateUserRole(ctx context.Context, name string,description string,  createdBy string) (*entity.UserRole, error) {
// 	role := entity.NewUserRole(uuid.New(), name, description, createdBy)
// 	if err := uc.userRoleRepo.CreateUserRole(ctx, role,); err != nil {
// 		return nil, err
// 	}
// 
// 	return role, nil
// }
// 
// Create Pinjaman Request
// func (uc *UserCreator) CreatePinjamanRequest(ctx context.Context, userid string, bookid string, tglpinjaman time.Time, tglkembali time.Time, requestedBy string) (*entity.Pinjaman, error) {
// 	pinjaman := entity.NewPinjaman(
// 		uuid.New(), 
// 		userid, 
// 		bookid, 
// 		utils.TimeToNullTime(tglpinjaman), 
// 		utils.TimeToNullTime(tglkembali) ,
// 		requestedBy, 
// 	)
// 	if err := uc.pinjamanRepo.CreatePinjamanRequest(ctx, pinjaman); err != nil {
// 		return nil, err
// 	}
// 
// 	return pinjaman, nil
// }