package repository

import (
	"context"
	"gin-starter/entity"
	"gin-starter/common/constant"
	"github.com/pkg/errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProvinceRepository is an struct for Province repository
type ProvinceRepository struct {
	gormDB *gorm.DB
}

// ProvinceRepositoryUseCase is an interface for Province repository use case
type ProvinceRepositoryUseCase interface {
	// FindAll returns all provinces
	GetProvinces(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Province, int64,error)
	// FindByID returns a province by its ID
	GetProvinceByID(ctx context.Context, id uuid.UUID) (*entity.Province, error)
}

// NewProvinceRepository creates a new Province repository
func NewProvinceRepository(
	db *gorm.DB,
) *ProvinceRepository {
	return &ProvinceRepository{
		gormDB: db,
	}
}

// FindAll returns all provinces
func (provinceRepo *ProvinceRepository) GetProvinces(ctx context.Context, search, sort, order string, limit, page int) ([]*entity.Province, int64, error) {
	var province []*entity.Province
	var total int64
	offsetProvince := ((page - 1) * limit)
	var gormDB = provinceRepo.gormDB.
		WithContext(ctx).
		Model(&entity.Province{}).
		Count(&total).
		Limit(limit).
		Offset(offsetProvince)

	if search != "" {
		gormDB = gormDB.Where("name ILIKE ?", "%"+search+"%").
			Or("alt_name ILIKE ?", "%"+search+"%").
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
	if err := gormDB.Find(&province).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0,errors.Wrap(err, "[ProvinceRepository-GetProvinces]")
	}
	return province, total,nil
	
}

func (provinceRepo *ProvinceRepository) GetProvinceByID(ctx context.Context, id uuid.UUID) (*entity.Province, error) {
	province := new(entity.Province)
	if err := provinceRepo.gormDB.
		WithContext(ctx).
		Model(&entity.Province{}).
		Where("id = ?", id).
		First(&province).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ProvinceRepository-GetProvinceByID]")
	}
	return province, nil
}
