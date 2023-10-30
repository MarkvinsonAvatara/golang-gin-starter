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
type VillageMasterFinder struct {
	cfg            config.Config
	villageRepo    repository.VillageRepositoryUseCase
}

// MasterFinderUseCase is a usecase for master
type VillageMasterFinderUseCase interface {
	// GetVillages returns all villages
	GetVillages(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Village, int64,error)
	// GetVillageByID returns a village by its ID
	GetVillageByID(ctx context.Context, id uuid.UUID) (*entity.Village, error)
}

// NewMasterFinder creates a new MasterFinder
func VillageNewMasterFinder(
	cfg config.Config,
	villageRepo repository.VillageRepositoryUseCase,

) *VillageMasterFinder {
	return &VillageMasterFinder{
		cfg:            cfg,
		villageRepo:    villageRepo,
	}
}

// GetVillages returns all villages
func (masterFinder *VillageMasterFinder) GetVillages(ctx context.Context,search, sort, order string, limit, page int) ([]*entity.Village, int64,error) {
	villages, total, err := masterFinder.villageRepo.GetVillages(ctx, search, sort, order, limit, page)

	if err != nil {
		return nil, 0,errors.ErrInternalServerError.Error()
	}

	return villages, total, nil
}

func (masterFinder *VillageMasterFinder) GetVillageByID(ctx context.Context, id uuid.UUID) (*entity.Village, error) {
	village, err := masterFinder.villageRepo.GetVillageByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return village, nil
}

