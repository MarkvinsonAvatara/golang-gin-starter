package repository

import (
	"context"
	// "encoding/json"
	// "fmt"
	// commonCache "gin-starter/common/cache"
	// "gin-starter/common/constant"
	"gin-starter/common/interfaces"
	"gin-starter/entity"
	// "log"
	// "time"

	// "github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

// UserRoleRepository is a repository for user role
type UpdateUserRoleRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// UserRoleRepositoryUseCase is a use case for user role
type UpdateUserRoleRepositoryUseCase interface {
	// UpdateUseRole is a method for updating user role
	UpdateUserRole(ctx context.Context, UserRole *entity.UserRole) error
	// Update is a method for updating user role
	// Update(ctx context.Context, userRole *entity.UserRole) error
}

// NewUserRoleRepository is a constructor for UserRoleRepository
func UpdateNewUserRoleRepository(db *gorm.DB, cache interfaces.Cacheable) *UpdateUserRoleRepository {
	return &UpdateUserRoleRepository{db, cache}
}


func (nc *UpdateUserRoleRepository) UpdateUserRole(ctx context.Context, UserRole *entity.UserRole) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.UserRole{}).
		Where(`id = ?`, UserRole.ID).
		Updates(UserRole).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}
	return nil
}

// // Update is a method for updating user role
// func (nc *UserRoleRepository) Update(ctx context.Context, userRole *entity.UserRole) error {
// 	oldTime := userRole.UpdatedAt
// 	userRole.UpdatedAt = time.Now()
// 	if err := nc.db.	
// 		WithContext(ctx).
// 		Transaction(func(tx *gorm.DB) error {
// 			sourceModel := new(entity.UserRole)
// 			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
// 				Where("user_id = ?", userRole.ID).
// 				Find(&sourceModel).Error; err != nil {
// 				log.Println("[GamPTKRepository - Update]", err)
// 				return err
// 			}
// 			if err := tx.Model(&entity.UserRole{}).
// 				Where(`user_id`, userRole.ID).
// 				UpdateColumns(sourceModel.MapUpdateFrom(userRole)).Error; err != nil {
// 				log.Println("[GamPTKRepository - Update]", err)
// 				return err
// 			}
// 			return nil
// 		}); err != nil {
// 		userRole.UpdatedAt = oldTime
// 	}
// 	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.UserRoleByUserID, "*")); err != nil {
// 		return err
// 	}
// 	return nil
// }


