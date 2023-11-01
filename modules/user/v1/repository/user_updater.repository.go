package repository

import (
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/entity"
	// "strings"
	"time"

	// "github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

// UserRepository is a repository for user
type UpdaterUserRepository struct {
	db *gorm.DB
}

// UserRepositoryUseCase is a use case for user
type UpdaterUserRepositoryUseCase interface {
	// ChangePassword is a function to change password
	ChangePassword(ctx context.Context, user *entity.User, newPassword string) error
	UpdateUser(ctx context.Context, user *entity.User) error
	// UpdateUserStatus is a function to update user status
	// UpdateUserStatus(ctx context.Context, id uuid.UUID, status string) error
}

// NewUserRepository creates a new UserRepository
func UpdaterNewUserRepository(db *gorm.DB) *UpdaterUserRepository {
	return &UpdaterUserRepository{db}
}



// // UpdateOTP is a function to update otp
// func (ur *UserRepository) UpdateOTP(ctx context.Context, user *entity.User, otp string) error {
// 	if err := ur.db.WithContext(ctx).
// 		Model(&entity.User{}).
// 		Where(`id = ?`, user.ID).
// 		Updates(
// 			map[string]interface{}{
// 				"otp":        otp,
// 				"updated_at": time.Now(),
// 			}).Error; err != nil {
// 		return errors.Wrap(err, "[UserRepository-Update] error when updating user data")
// 	}
// 	return nil
// }
// 
// // Update is a function to update user
// // func (ur *UserRepository) Update(ctx context.Context, user *entity.User) error {
// // 	oldTime := user.UpdatedAt
// // 	user.UpdatedAt = time.Now()
// // 	if err := ur.db.
// // 		WithContext(ctx).
// // 		Transaction(func(tx *gorm.DB) error {
// // 			sourceModel := new(entity.User)
// // 			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
// // 				log.Println("[GamPTKRepository - Update]", err)
// // 				return err
// // 			}
// // 			if err := tx.Model(&entity.User{}).
// // 				Where(`id`, user.ID).
// // 				UpdateColumns(sourceModel.MapUpdateFrom(user)).Error; err != nil {
// // 				log.Println("[GamPTKRepository - Update]", err)
// // 				return err
// // 			}
// // 			return nil
// // 		}); err != nil {
// // 		user.UpdatedAt = oldTime
// // 	}
// // 	return nil
// // }

// ChangePassword is a function to change password
func (ur *UpdaterUserRepository) ChangePassword(ctx context.Context, user *entity.User, newPassword string) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, user.ID).
		Updates(
			map[string]interface{}{
				"password":   newPassword,
				"updated_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-Update] error when updating user data")
	}

	return nil
}


// UpdateUser is a function to update user
func (ur *UpdaterUserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, user.ID).
		Updates(user).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-Update] error when updating user data")
	}

	return nil
}

// UpdateUserStatus is a function to update user status
// func (ur *UserRepository) UpdateUserStatus(ctx context.Context, id uuid.UUID, status string) error {
// 	if err := ur.db.WithContext(ctx).
// 		Model(&entity.User{}).
// 		Where(`id = ?`, id).
// 		Updates(
// 			map[string]interface{}{
// 				"status":     status,
// 				"updated_at": time.Now(),
// 			}).Error; err != nil {
// 		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
// 	}
// 
// 	return nil
// }

