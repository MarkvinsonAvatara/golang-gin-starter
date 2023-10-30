package repository

import (
	"context"
	"gin-starter/entity"
	"github.com/pkg/errors"
	"fmt"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type DistrictRepositoryUseCase interface {
	//GetDistrict returns a list of district
	GetDistricts(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.District, int64, error)
	//GetDistrictByID returns a district by its ID
	GetDistrictByID(ctx context.Context, id uuid.UUID) (*entity.District, error)
	// FindByRegencyID finds district by regency id
	// FindByRegencyID(ctx context.Context, id int64) ([]*entity.District, error)
}

// DistrictRepository is an struct for District repository
type DistrictRepository struct {
	gormDB *gorm.DB
}

// NewDistrictRepository creates a new District repository
func NewDistrictRepository(
	db *gorm.DB,
) *DistrictRepository {
	return &DistrictRepository{
		gormDB: db,
	}
}

// GetDistricts returns a list of districts
func (repo *DistrictRepository) GetDistricts(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.District, int64, error) {
	var district []*entity.District
	var total int64
	offsetDistrict := ((page - 1) * limit)
	var gormDB = repo.gormDB.
		WithContext(ctx).
		Model(&entity.District{}).
		Count(&total).
		Limit(limit).
		Offset(offsetDistrict)

	if search != "" {
		gormDB = gormDB.Where("name ILIKE ?", "%"+search+"%").
			Or("alt_name ILIKE ?", "%"+search+"%").
			Or("CAST(latitude AS TEXT) ILIKE ?", "%"+search+"%").
			Or("CAST(longitude AS TEXT) ILIKE ?", "%"+search+"%")
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	if sort == "" {
		sort = "id"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&district).Error; err != nil {
		return nil, 0, errors.Wrap(err, "[DistrictRepository-GetDistricts]")
	}

	return district, total, nil
}

// GetDistrictByID returns a district by its ID
func (repo *DistrictRepository) GetDistrictByID(ctx context.Context, id uuid.UUID) (*entity.District, error) {
	district := new(entity.District)
	if err := repo.gormDB.
		WithContext(ctx).
		Model(&entity.District{}).
		Where("id = ?", id).
		Find(&district).
		Error; err != nil {
		return nil, errors.Wrap(err, "[DistrictRepository-GetDistrictByID]")
	}
	return district, nil
}



// // FindByRegencyID finds district by regency id
// func (repo *DistrictRepository) FindByRegencyID(ctx context.Context, id int64) ([]*entity.District, error) {
// 	models := make([]*entity.District, 0)
// 	if err := repo.gormDB.
// 		WithContext(ctx).
// 		Model(&entity.District{}).
// 		Where("regency_id = ? ", id).
// 		Find(&models).
// 		Error; err != nil {
// 		return nil, errors.Wrap(err, "[DistrictRepository-FindByRegencyID]")
// 	}
// 	return models, nil
// }
