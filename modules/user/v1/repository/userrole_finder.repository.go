package repository

import (
	"context"
	"encoding/json"
	"fmt"
	commonCache "gin-starter/common/cache"
	"gin-starter/common/constant"
	"gin-starter/common/interfaces"
	"gin-starter/entity"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// UserRoleRepository is a repository for user role
type FinderUserRoleRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// UserRoleRepositoryUseCase is a use case for user role
type FinderUserRoleRepositoryUseCase interface {
	// GetUser Role gets all user role
	GetUserRoles(ctx context.Context, search, sort, order string, limit, Page int) ([]*entity.UserRole, int64, error)
	// FindByUserID is a method for finding user role by user id
	GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error)
}

// NewUserRoleRepository is a constructor for UserRoleRepository
func FinderNewUserRoleRepository(db *gorm.DB, cache interfaces.Cacheable) *FinderUserRoleRepository {
	return &FinderUserRoleRepository{db, cache}
}

func (nc *FinderUserRoleRepository) GetUserRoles(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.UserRole, int64, error) {
	var userRoles []*entity.UserRole
	var total int64
	offsetUser:=((page - 1)*limit)
	var gormDB = nc.db.
		WithContext(ctx).
		Model(&entity.UserRole{}).
		Count(&total).
		Limit(limit).
		Offset(offsetUser)

	if search != "" {
		gormDB = gormDB.
			Where("name ILIKE ?", "%"+search+"%")
	}

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
		return nil, 0, errors.Wrap(err, "[UserRoleRepository-GetUserRoles] error when looking up all user")
	}

	return userRoles, total, nil
}

// FindByUserID is a method for finding user role by user id
func (nc *FinderUserRoleRepository) GetUserRoleByID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error) {
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




