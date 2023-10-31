package service

import (
	//"bytes"
	"context"
	//"fmt"
	//"gin-starter/common/constant"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/pinjaman/v1/repository"
	// "gin-starter/utils"
	// "github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
	// "html/template"
	// "log"
)

// UserUpdater is a struct that contains the dependencies of UserUpdater
type UserUpdater struct {
	cfg            config.Config
	pinjamanRepo   repository.PinjamanRepositoryUseCase
}

// UserUpdaterUseCase is a struct that contains the dependencies of UserUpdaterUseCase
type PinjamanUpdaterUseCase interface {
	// HandledPinjaman updates pinjaman
	HandledPinjaman(ctx context.Context, pinjaman *entity.Pinjaman) error
}

// NewUserUpdater is a function that creates a new UserUpdater
func NewPinjamanUpdater(
	cfg config.Config,
	pinjamanRepo repository.PinjamanRepositoryUseCase,
) *UserUpdater {
	return &UserUpdater{
		cfg:            cfg,
		pinjamanRepo:   pinjamanRepo,
	}
}


func (uu *UserUpdater) HandledPinjaman(ctx context.Context, pinjaman *entity.Pinjaman) error {
	if err := uu.pinjamanRepo.HandledPinjaman(ctx, pinjaman); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
