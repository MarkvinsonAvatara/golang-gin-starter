package service

import (
	"context"
	"gin-starter/common/errors"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/master/v1/repository"
	"github.com/google/uuid"
)

// MasterFinder is a service for master
type DistrictMasterFinder struct {
	cfg            config.Config
	districtRepo   repository.DistrictRepositoryUseCase
}

// MasterFinderUseCase is a usecase for master
type DistrictMasterFinderUseCase interface {
	//GetDistricts returns a list of districts
	GetDistricts(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.District, int64, error)
	//GetDistrictByID returns a district by its ID
	GetDistrictByID(ctx context.Context, id uuid.UUID) (*entity.District, error)
}

// NewMasterFinder creates a new MasterFinder
func DistrictNewMasterFinder(
	cfg config.Config,
	districtRepo repository.DistrictRepositoryUseCase,

) *DistrictMasterFinder {
	return &DistrictMasterFinder{
		cfg:            cfg,
		districtRepo:   districtRepo,
	}
}

// GetDistricts returns a list of districts
func (masterfinder *DistrictMasterFinder) GetDistricts(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.District, int64, error) {
	districts, total, err := masterfinder.districtRepo.GetDistricts(ctx, search, sort, order, limit, page)
	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}
	return districts, total, nil
}

func (masterfinder *DistrictMasterFinder) GetDistrictByID(ctx context.Context, id uuid.UUID) (*entity.District, error) {
	district, err := masterfinder.districtRepo.GetDistrictByID(ctx, id)
	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}
	return district, nil
}



