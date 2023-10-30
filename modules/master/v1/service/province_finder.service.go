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
type ProvinceMasterFinder struct {
	cfg            config.Config
	provinceRepo   repository.ProvinceRepositoryUseCase
}

// MasterFinderUseCase is a usecase for master
type ProvinceMasterFinderUseCase interface {
	GetProvinces(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Province, int64, error)
	GetProvinceByID(ctx context.Context, id uuid.UUID) (*entity.Province, error)

}

// NewMasterFinder creates a new MasterFinder
func ProvinceNewMasterFinder(
	cfg config.Config,
	provinceRepo repository.ProvinceRepositoryUseCase,
) *ProvinceMasterFinder {
	return &ProvinceMasterFinder{
		cfg:            cfg,
		provinceRepo:   provinceRepo,
	}
}

// GetProvinces returns all provinces
func (masterFinder *ProvinceMasterFinder) GetProvinces(ctx context.Context,search, sort, order string, limit, page int) ([]*entity.Province, int64,error) {
	provinces, total, err := masterFinder.provinceRepo.GetProvinces(ctx, search, sort, order, limit, page)

	if err != nil {
		return nil, total,errors.ErrInternalServerError.Error()
	}

	return provinces, total, nil
}

// GetProvincesByID returns a province by its ID
func (masterFinder *ProvinceMasterFinder) GetProvinceByID(ctx context.Context, id uuid.UUID) (*entity.Province, error) {
	province, err := masterFinder.provinceRepo.GetProvinceByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return province, nil
}

