package service

import (
	"context"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/pinjaman/v1/repository"
	"gin-starter/utils"
	"time"

	"github.com/google/uuid"
)

// UserCreator is a struct that contains all the dependencies for the User creator
type PinjamanCreator struct {
	cfg          config.Config
	pinjamanRepo repository.PinjamanRepositoryUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

type PinjamanCreatorUseCase interface {
	CreatePinjaman(ctx context.Context, tanggalPinjaman time.Time, tanggalKembali time.Time, status string) (*entity.Pinjaman, error)
}

// NewUserCreator is a constructor for the User creator
func NewPinjamanCreator(
	cfg config.Config,
	pinjamanRepo repository.PinjamanRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *PinjamanCreator {
	return &PinjamanCreator{
		cfg:          cfg,
		pinjamanRepo: pinjamanRepo,
		cloudStorage: cloudStorage,
	}
}

// CreateRole creates a User role
func (PinjamanCreator *PinjamanCreator) CreateUserRole(ctx context.Context, tglpinjam time.Time, tglkembali time.Time, status string, createdBy string) (*entity.Pinjaman, error) {
	role := entity.NewPinjaman(
		uuid.New(), 
		utils.TimeToNullTime(tglpinjam), 
		utils.TimeToNullTime(tglkembali),
		status,
		createdBy)
	if err := PinjamanCreator.pinjamanRepo.CreatePinjaman(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}
