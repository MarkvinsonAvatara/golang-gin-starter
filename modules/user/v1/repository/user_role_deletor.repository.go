package repository

import (
	"context"
	"fmt"
	commonCache "gin-starter/common/cache"
	"gin-starter/common/interfaces"
	"gin-starter/entity"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// UserRoleRepository is a repository for user role
type DeleteUserRoleRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// UserRoleRepositoryUseCase is a use case for user role
type DeleteUserRoleRepositoryUseCase interface {
	// Delete is a method for deleting user role
	DeleteUserRole(ctx context.Context, id uuid.UUID) error

}

// NewUserRoleRepository is a constructor for UserRoleRepository
func DeleteNewUserRoleRepository(db *gorm.DB, cache interfaces.Cacheable) *DeleteUserRoleRepository {
	return &DeleteUserRoleRepository{db, cache}
}



// Delete is a method for deleting user role
func (nc *DeleteUserRoleRepository) DeleteUserRole(ctx context.Context, id uuid.UUID) error {
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
