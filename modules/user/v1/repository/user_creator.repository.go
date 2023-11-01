package repository

import (
	"context"
	// "fmt"
	// "gin-starter/common/constant"
	"gin-starter/entity"
	// "strings"
	// "time"

	// "github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

// UserRepository is a repository for user
type CreateUserRepository struct {
	db *gorm.DB
}

// UserRepositoryUseCase is a use case for user
type CreateUserRepositoryUseCase interface {
	// CreateUser is a function to create user
	CreateUser(ctx context.Context, user *entity.User) error
}

// NewUserRepository creates a new UserRepository
func CreateNewUserRepository(db *gorm.DB) *CreateUserRepository {
	return &CreateUserRepository{db}
}

// CreateUser is a function to create user
func (ur *CreateUserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Create(user).
		Error; err != nil {
		return errors.Wrap(err, "[UserRepository-CreateUser] error while creating user")
	}

	return nil
}
