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
type MasterFinder struct {
	cfg            config.Config
	provinceRepo   repository.ProvinceRepositoryUseCase
	regencyRepo    repository.RegencyRepositoryUseCase
	districtRepo   repository.DistrictRepositoryUseCase
	villageRepo    repository.VillageRepositoryUseCase
	bookRepository repository.BookRepositoryUseCase
}

// MasterFinderUseCase is a usecase for master
type MasterFinderUseCase interface {
	// GetProvinces returns all provinces
	GetProvinces(ctx context.Context) ([]*entity.Province, error)
	// GetRegencies returns all regencies
	GetRegencies(ctx context.Context, id int64) ([]*entity.Regency, error)
	// GetDistricts returns all districts
	GetDistricts(ctx context.Context, id int64) ([]*entity.District, error)
	// GetVillages returns all villages
	GetVillages(ctx context.Context, id int64) ([]*entity.Village, error)
	// GetBooks returns all books
	GetBooks(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error)
	// GetBookByID returns a book by its ID
	GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error)
}

// NewMasterFinder creates a new MasterFinder
func NewMasterFinder(
	cfg config.Config,
	provinceRepo repository.ProvinceRepositoryUseCase,
	regencyRepo repository.RegencyRepositoryUseCase,
	districtRepo repository.DistrictRepositoryUseCase,
	villageRepo repository.VillageRepositoryUseCase,
	bookRepository repository.BookRepositoryUseCase,

) *MasterFinder {
	return &MasterFinder{
		cfg:            cfg,
		provinceRepo:   provinceRepo,
		regencyRepo:    regencyRepo,
		districtRepo:   districtRepo,
		villageRepo:    villageRepo,
		bookRepository: bookRepository,
	}
}

// GetProvinces returns all provinces
func (masterFinder *MasterFinder) GetProvinces(ctx context.Context) ([]*entity.Province, error) {
	provinces, err := masterFinder.provinceRepo.FindAll(ctx)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return provinces, nil
}

// GetRegencies returns all regencies
func (masterFinder *MasterFinder) GetRegencies(ctx context.Context, id int64) ([]*entity.Regency, error) {
	regencies, err := masterFinder.regencyRepo.FindByProvinceID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return regencies, nil
}

// GetDistricts returns all districts
func (masterFinder *MasterFinder) GetDistricts(ctx context.Context, id int64) ([]*entity.District, error) {
	districts, err := masterFinder.districtRepo.FindByRegencyID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return districts, nil
}

// GetVillages returns all villages
func (masterFinder *MasterFinder) GetVillages(ctx context.Context, id int64) ([]*entity.Village, error) {
	villages, err := masterFinder.villageRepo.FindByDistrictID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return villages, nil
}

// GetBooks returns all books
func (masterFinder *MasterFinder) GetBooks(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Book, int64, error) {
	books, total, err := masterFinder.bookRepository.GetBooks(ctx, search, sort, order, limit, page)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return books, total, nil
}

func (masterFinder *MasterFinder) GetBookByID(ctx context.Context, id uuid.UUID) (*entity.Book, error) {
	book, err := masterFinder.bookRepository.GetBookByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return book, nil
}
