package repository

import (
	"context"
	"fmt"
	"gin-starter/common/constant"
	"gin-starter/entity"
	"strings"
	// "time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

// UserRepository is a repository for user
type FinderUserRepository struct {
	db *gorm.DB
}

// UserRepositoryUseCase is a use case for user
type FinderUserRepositoryUseCase interface {
	// GetUserByID is a function to get user by id
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	// // GetUserByForgotPasswordToken is a function to get user by forgot password token
	// GetUserByForgotPasswordToken(ctx context.Context, token string) (*entity.User, error)
	// GetUsers is a function to get users
	GetUsers(ctx context.Context, search, filter, sort, order string, limit, page int) ([]*entity.User, int64, error)
	// GetAdminUsers is a function to get admin users
	GetAdminUsers(ctx context.Context, search, sort, order string, limit, offset int) ([]*entity.User, int64, error)
}

// NewUserRepository creates a new UserRepository
func FinderNewUserRepository(db *gorm.DB) *FinderUserRepository {
	return &FinderUserRepository{db}
}

// GetUserByID is a function to get user by id
func (ur *FinderUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	result := new(entity.User)

	if err := ur.db.
		WithContext(ctx).
		// Preload("UserRole").
		// Preload("UserRole.Role").
		// Preload("Employee").
		// Preload("Employee.CustomerBranch").
		// Preload("Employee.CustomerBranch.Province").
		// Preload("Employee.CustomerBranch.Regency").
		// Preload("Employee.CustomerBranch.District").
		// Preload("Employee.CustomerBranch.Village").
		Where("id = ?", id).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByID] user not found")
	}

	return result, nil
}

// GetUsers is a function to get all users
func (ur *FinderUserRepository) GetUsers(ctx context.Context, search, filter, sort, order string, limit, page int) ([]*entity.User, int64, error) {
	var user []*entity.User
	var total int64
	offsetUser := ((page - 1) * limit)
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Where("public.user.deleted_at is NULL").
		Count(&total).
		Limit(limit).
		Offset(offsetUser)

	if search != "" {
		gormDB = gormDB.
			Where("name ILIKE ?", "%"+search+"%").
			Or("email ILIKE ?", "%"+search+"%").
			Or("Cast(dob AS TEXT) ILIKE ?", "%"+search+"%")
	}

	if filter == strings.ToLower("user") {
		gormDB = gormDB.Where("public.user.roleid = '4ad98cc1-f2c7-4dc8-b67f-56c15022b05d'")
	} else if filter == strings.ToLower("admin") {
		gormDB = gormDB.Where("public.user.roleid = '6cb0f2c1-0408-44ec-b1ad-1095b57ec544'")
	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[UserRepository-GetAdminUsers] error when looking up all user")
	}

	return user, total, nil
}

// GetAdminUsers is a function to get all admin users
func (ur *FinderUserRepository) GetAdminUsers(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.User, int64, error) {
	var user []*entity.User
	var total int64
	offsetUser := ((page - 1) * limit)
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Where("public.user.deleted_at is NULL").
		Joins("inner join public.role on public.user.roleid=public.role.id").
		Where("public.role.name = 'Admin'").
		Count(&total).
		Limit(limit).
		Offset(offsetUser)

	if search != "" {
		gormDB = gormDB.
			Where("name ILIKE ?", "%"+search+"%").
			Or("email ILIKE ?", "%"+search+"%").
			Or("phone_number ILIKE ?", "%"+search+"%")
	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[UserRepository-GetAdminUsers] error when looking up all user")
	}

	return user, total, nil
}
