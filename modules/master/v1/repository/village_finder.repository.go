package repository

import (
	"context"
	"gin-starter/entity"
	"github.com/pkg/errors"
	"gin-starter/common/constant"
	"github.com/google/uuid"
	"fmt"
	"gorm.io/gorm"
)

type VillageRepositoryUseCase interface {
	// FindByDistrictID finds village by district id
	// FindByDistrictID(ctx context.Context, id int64) ([]*entity.Village, error)
	GetVillages(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Village, int64, error)
	GetVillageByID(ctx context.Context, id uuid.UUID) (*entity.Village, error)

	
}

// VillageRepository is an struct for Village repository
type VillageRepository struct {
	gormDB *gorm.DB
}

// NewVillageRepository creates a new Village repository
func NewVillageRepository(
	db *gorm.DB,
) *VillageRepository {
	return &VillageRepository{
		gormDB: db,
	}
}



// // FindByDistrictID finds village by district id
// func (repo *VillageRepository) FindByDistrictID(ctx context.Context, id int64) ([]*entity.Village, error) {
// 	models := make([]*entity.Village, 0)
// 	if err := repo.gormDB.
// 		WithContext(ctx).
// 		Model(&entity.Village{}).
// 		Where("district_id = ? ", id).
// 		Find(&models).
// 		Error; err != nil {
// 		return nil, errors.Wrap(err, "[VillageRepository-FindByDistrictID]")
// 	}
// 	return models, nil
// }

// GetVillages returns a list of villages
func (repo *VillageRepository) GetVillages(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Village, int64, error) {
	var village []*entity.Village
	var total int64
	offsetVillage := ((page - 1) * limit)
	var gormDB = repo.gormDB.
		WithContext(ctx).
		Model(&entity.Village{}).
		Count(&total).
		Limit(limit).
		Offset(offsetVillage)

	if search != "" {
		gormDB = gormDB.Where("name ILIKE ?", "%"+search+"%").
			Or("alt_name LIKE ?", "%"+search+"%").
			Or("CAST(latitude AS TEXT) ILIKE ?", "%"+search+"%").
			Or("CAST(longitude AS TEXT) ILIKE ?", "%"+search+"%")
	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "id"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))
		

	if err := gormDB.Find(&village).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[VillageRepository-GetVillages]")
	}

	return village, total, nil
}

// GetVillageByID returns a village by its ID
func (repo *VillageRepository)GetVillageByID(ctx context.Context, id uuid.UUID) (*entity.Village, error) {
	village := new(entity.Village)
	if err := repo.gormDB.
		WithContext(ctx).
		Model(&entity.Village{}).
		Where("id = ?", id).
		Find(&village).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[VillageRepository-GetVillageByID]")
	}
	return village, nil
}



