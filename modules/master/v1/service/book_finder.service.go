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
type BookMasterFinder struct {
	cfg            config.Config
	// regencyRepo    repository.RegencyRepositoryUseCase
	// districtRepo   repository.DistrictRepositoryUseCase
	// villageRepo    repository.VillageRepositoryUseCase
	finderBookRepository repository.FinderBookRepositoryUseCase
}

// MasterFinderUseCase is a usecase for master
type BookMasterFinderUseCase interface {
	// GetRegencies returns all regencies
	// GetRegencies(ctx context.Context, id int64) ([]*entity.Regency, error)
	// GetDistricts returns all districts
	// GetDistricts(ctx context.Context, id int64) ([]*entity.District, error)
	// // GetVillages returns all villages
	// GetVillages(ctx context.Context, id int64) ([]*entity.Village, error)
	// GetBooks returns all books
	GetBooks(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error)
	// GetBookByID returns a book by its ID
	GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
	GetBookAvalaibily(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error)
}

// NewMasterFinder creates a new MasterFinder
func BookNewMasterFinder(
	cfg config.Config,
	// regencyRepo repository.RegencyRepositoryUseCase,
	// districtRepo repository.DistrictRepositoryUseCase,
	// villageRepo repository.VillageRepositoryUseCase,
	finderBookRepository repository.FinderBookRepositoryUseCase,

) *BookMasterFinder {
	return &BookMasterFinder{
		cfg:            cfg,
		// regencyRepo:    regencyRepo,
		// districtRepo:   districtRepo,
		// villageRepo:    villageRepo,
		finderBookRepository: finderBookRepository,
	}
}


// // GetRegencies returns all regencies
// func (masterFinder *BookMasterFinder) GetRegencies(ctx context.Context, id int64) ([]*entity.Regency, error) {
// 	regencies, err := masterFinder.regencyRepo.FindByProvinceID(ctx, id)

// 	if err != nil {
// 		return nil, errors.ErrInternalServerError.Error()
// 	}

// 	return regencies, nil
// }

// GetDistricts returns all districts
// func (masterFinder *BookMasterFinder) GetDistricts(ctx context.Context, id int64) ([]*entity.District, error) {
// 	districts, err := masterFinder.districtRepo.FindByRegencyID(ctx, id)

// 	if err != nil {
// 		return nil, errors.ErrInternalServerError.Error()
// 	}

// 	return districts, nil
// }

// // GetVillages returns all villages
// func (masterFinder *BookMasterFinder) GetVillages(ctx context.Context, id int64) ([]*entity.Village, error) {
// 	villages, err := masterFinder.villageRepo.FindByDistrictID(ctx, id)

// 	if err != nil {
// 		return nil, errors.ErrInternalServerError.Error()
// 	}

// 	return villages, nil
// }

// GetBooks returns all books
func (masterFinder *BookMasterFinder) GetBooks(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error) {
	books, total, err := masterFinder.finderBookRepository.GetBooks(ctx, search, sort, order, limit, page)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return books, total, nil
}

func (masterFinder *BookMasterFinder) GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
	book, err := masterFinder.finderBookRepository.GetBookByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return book, nil
}

func (masterFinder *BookMasterFinder) GetBookAvalaibily(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error) {
	books, total, err := masterFinder.finderBookRepository.GetBookAvalaibily(ctx, search, sort, order, limit, page)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return books, total, nil
}

