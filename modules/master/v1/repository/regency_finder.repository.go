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

type RegencyRepositoryUseCase interface {
	//GetRegency returns a list of regency
	GetRegencies(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Regency, int64, error)
	//GetRegencyByID returns a regency by its ID
	GetRegencyByID(ctx context.Context, id uuid.UUID) (*entity.Regency, error)
	// FindByProvinceID finds regency by province id
	// FindByProvinceID(ctx context.Context, id int64) ([]*entity.Regency, error)
}

// RegencyRepository is an struct for Regency repository
type RegencyRepository struct {
	gormDB *gorm.DB
}

// NewRegencyRepository creates a new Regency repository
func NewRegencyRepository(
	db *gorm.DB,
) *RegencyRepository {
	return &RegencyRepository{
		gormDB: db,
	}
}

// GetRegencies returns a list of regencies
func (repo *RegencyRepository) GetRegencies(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Regency, int64, error) {
	var regency []*entity.Regency
	var total int64
	offsetRegency := ((page - 1) * limit)
	var gormDB = repo.gormDB.
		WithContext(ctx).
		Model(&entity.Regency{}).
		Count(&total).
		Limit(limit).
		Offset(offsetRegency)

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
	if err:=gormDB.Find(&regency).Error; err != nil {
		if err==gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[RegencyRepository-GetRegencies]")
	}
	return regency, total, nil
}

// GetRegencyByID returns a regency by its ID
func (repo *RegencyRepository) GetRegencyByID(ctx context.Context, id uuid.UUID) (*entity.Regency, error) {
	regency := new(entity.Regency)
	if err := repo.gormDB.
		WithContext(ctx).
		Model(&entity.Regency{}).
		Where("id = ?", id).
		Find(&regency).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[RegencyRepository-GetRegencyByID]")
	}
	return regency, nil
}

// // FindByProvinceID finds regency by province id
// func (repo *RegencyRepository) FindByProvinceID(ctx context.Context, id int64) ([]*entity.Regency, error) {
// 	models := make([]*entity.Regency, 0)
// 	if err := repo.gormDB.
// 		WithContext(ctx).
// 		Model(&entity.Regency{}).
// 		Where("province_id = ? ", id).
// 		Find(&models).
// 		Error; err != nil {
// 		return nil, errors.Wrap(err, "[RegencyRepository-FindByProvinceID]")
// 	}
// 	return models, nil
// }
