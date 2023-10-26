package service

import (
	"context"
	// "gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/entity"
	// notificationService "gin-starter/modules/notification/v1/service"
	"gin-starter/modules/pinjaman/v1/repository"
	"gin-starter/utils"
	"time"

	"github.com/google/uuid"
)

// UserCreator is a struct that contains all the dependencies for the User creator
type UserCreator struct {
	cfg            config.Config
	pinjamanRepo  repository.PinjamanRepositoryUseCase
	cloudStorage   interfaces.CloudStorageUseCase
}

// UserCreatorUseCase is a use case for the User creator
type PinjamanCreatorUseCase interface {
	// Create Pinjaman Request
	CreatePinjamanRequest(ctx context.Context, userid string, bookid string, tglpinjaman time.Time, tglkembali time.Time, requestedBy string)(*entity.Pinjaman, error)

}


// NewUserCreator is a constructor for the User creator
func NewPinjamanCreator(
	cfg config.Config,
	pinjamanRepo repository.PinjamanRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserCreator {
	return &UserCreator{
		cfg:            cfg,
		pinjamanRepo:  pinjamanRepo,
		cloudStorage:   cloudStorage,
	}
}

// Create Pinjaman Request
func (uc *UserCreator) CreatePinjamanRequest(ctx context.Context, userid string, bookid string, tglpinjaman time.Time, tglkembali time.Time, requestedBy string) (*entity.Pinjaman, error) {
	pinjaman := entity.NewPinjaman(
		uuid.New(), 
		userid, 
		bookid, 
		utils.TimeToNullTime(tglpinjaman), 
		utils.TimeToNullTime(tglkembali) ,
		requestedBy, 
	)
	if err := uc.pinjamanRepo.CreatePinjamanRequest(ctx, pinjaman); err != nil {
		return nil, err
	}

	return pinjaman, nil
}