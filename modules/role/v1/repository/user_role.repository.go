package repository

import (
	"context"
	"encoding/json"
	"fmt"
	commonCache "gin-starter/common/cache"
	"gin-starter/common/constant"
	"gin-starter/common/interfaces"
	"gin-starter/entity"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// UserRoleRepository is a repository for user role
type UserRoleRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// UserRoleRepositoryUseCase is a use case for user role
type UserRoleRepositoryUseCase interface {
	// GetUser Role gets all user role
	GetUserRoles(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.UserRole, int64, error)
	// FindByUserID is a method for finding user role by user id
	GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error)
	// UpdateUseRole is a method for updating user role
	UpdateUserRole(ctx context.Context, UserRole *entity.UserRole) error
	// Delete is a method for deleting user role
	DeleteUserRole(ctx context.Context, id uuid.UUID) error
	// CreateUserRole is a method for creating user role
	CreateUserRole(ctx context.Context, role *entity.UserRole) error
}

// NewUserRoleRepository is a constructor for UserRoleRepository
func NewUserRoleRepository(db *gorm.DB, cache interfaces.Cacheable) *UserRoleRepository {
	return &UserRoleRepository{db, cache}
}


func (nc *UserRoleRepository) CreateUserRole(ctx context.Context, role *entity.UserRole) error {

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.UserRole{}).
		Create(role).
		Error; err != nil {
		return errors.Wrap(err, "[UserRepository-CreateUser] error while creating user")
	}

	return nil
}

func (nc *UserRoleRepository) GetUserRoles(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.UserRole, int64, error) {
	var userRoles []*entity.UserRole
	var total int64
	var gormDB = nc.db.
		WithContext(ctx).
		Model(&entity.UserRole{}).
		Find(&userRoles)

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset)

	// if query != "" {
	// 	gormDB = gormDB.
	// 		Where("name ILIKE ?", "%"+query+"%").
	// 		Or("email ILIKE ?", "%"+query+"%").
	// 		Or("phone_number ILIKE ?", "%"+query+"%")
	// }

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&userRoles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[UserRepository-GetAdminUsers] error when looking up all user")
	}

	return userRoles, total, nil
}

// FindByUserID is a method for finding user role by user id
func (nc *UserRoleRepository) GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error) {
	category := &entity.UserRole{}

	bytes, _ := nc.cache.Get(fmt.Sprintf(
		commonCache.UserRoleByUserID, id.String()))

	if bytes != nil {
		if err := json.Unmarshal(bytes, &category); err != nil {
			return nil, err
		}
		return category, nil
	}

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.UserRole{}).
		Where("id = ?", id).
		First(&category).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[NewsRepository-FindByID] error while getting category category")
	}

	if err := nc.cache.Set(fmt.Sprintf(
		commonCache.UserRoleByUserID, id.String()), &category, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return category, nil
}

func (nc *UserRoleRepository) UpdateUserRole(ctx context.Context, UserRole *entity.UserRole) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.UserRole{}).
		Where(`id = ?`, UserRole.ID).
		Updates(UserRole).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}
	return nil
}

// Delete is a method for deleting user role
func (nc *UserRoleRepository) DeleteUserRole(ctx context.Context, id uuid.UUID) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.UserRole{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"updated_at": time.Now(),
				"deleted_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.UserRoleByUserID, "*")); err != nil {
		return err
	}
	return nil
}
