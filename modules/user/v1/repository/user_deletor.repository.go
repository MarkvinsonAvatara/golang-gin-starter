package repository

import (
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/entity"
	// "strings"
	// "time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

// UserRepository is a repository for user
type DeleteUserRepository struct {
	db *gorm.DB
}

// UserRepositoryUseCase is a use case for user
type DeleteUserRepositoryUseCase interface {
	// DeleteAdmin is a function to delete admin user
	DeleteAdmin(ctx context.Context, id uuid.UUID) error
	// DeleteUsers is a function to delete user
	DeleteUsers(ctx context.Context, id uuid.UUID) error
}

// NewUserRepository creates a new UserRepository
func DeleteNewUserRepository(db *gorm.DB) *DeleteUserRepository {
	return &DeleteUserRepository{db}
}

// DeleteAdmin is a function to delete admin user
func (ur *DeleteUserRepository) DeleteAdmin(ctx context.Context, id uuid.UUID) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, id).
		Delete(&entity.User{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeleteAdmin] error when updating user data")
	}

	return nil
}

// DeleteUsers is a function to delete user
func (ur *DeleteUserRepository) DeleteUsers(ctx context.Context, id uuid.UUID) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, id).
		Delete(&entity.User{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeleteUsers] error when updating user data")
	}

	return nil
}
