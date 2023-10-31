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
type CreateUserRoleRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// UserRoleRepositoryUseCase is a use case for user role
type CreateUserRoleRepositoryUseCase interface {
	// CreateUserRole is a method for creating user role
	CreateUserRole(ctx context.Context, role *entity.UserRole) error
}

// NewUserRoleRepository is a constructor for UserRoleRepository
func CreateNewUserRoleRepository(db *gorm.DB, cache interfaces.Cacheable) *CreateUserRoleRepository {
	return &CreateUserRoleRepository{db, cache}
}


func (nc *CreateUserRoleRepository) CreateUserRole(ctx context.Context, role *entity.UserRole) error {

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.UserRole{}).
		Create(role).
		Error; err != nil {
		return errors.Wrap(err, "[UserRepository-CreateUser] error while creating user")
	}

	return nil
}
