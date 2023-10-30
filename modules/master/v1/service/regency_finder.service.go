package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"github.com/google/uuid"
	"gin-starter/entity"
	"gin-starter/modules/master/v1/repository"
)

// MasterFinder is a service for master
type RegencyMasterFinder struct {
	cfg            config.Config
	regencyRepo    repository.RegencyRepositoryUseCase
}

// MasterFinderUseCase is a usecase for master
type RegencyMasterFinderUseCase interface {
	//GetRegency returns a list of regency
	GetRegencies(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Regency, int64, error)
	//GetRegencyByID returns a regency by its ID
	GetRegencyByID(ctx context.Context, id uuid.UUID) (*entity.Regency, error)
}

// NewMasterFinder creates a new MasterFinder
func RegencyNewMasterFinder(
	cfg config.Config,
	regencyRepo repository.RegencyRepositoryUseCase,

) *RegencyMasterFinder {
	return &RegencyMasterFinder{
		cfg:            cfg,
		regencyRepo:    regencyRepo,
	}
}


// GetRegencies returns all regencies
func (masterFinder *RegencyMasterFinder) GetRegencies(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Regency, int64,error) {
	regencies, total, err := masterFinder.regencyRepo.GetRegencies(ctx, search, sort, order, limit, page)

	if err != nil {
		return nil, 0,errors.ErrInternalServerError.Error()
	}

	return regencies, total,nil
}

func (masterFinder *RegencyMasterFinder) GetRegencyByID(ctx context.Context, id uuid.UUID) (*entity.Regency, error) {
	regency, err := masterFinder.regencyRepo.GetRegencyByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return regency, nil
}
