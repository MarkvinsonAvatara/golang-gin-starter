package resource

import "gin-starter/entity"


// Province is a struct for province
type Province struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetProvinceByIDRequest struct {
	ID int64 `uri:"id" json:"id" binding:"required"`
}

// ProvinceListResponse is a struct for province list response
type ProvinceListResponse struct {
	List []*Province `json:"list"`
	Meta *Meta       `json:"meta"`
}

type GetProvinceRequest struct {
	Search string `form:"search" json:"search"`
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Page   int    `form:"page,default=0" json:"page"`
}



func NewProvinceResponse(province *entity.Province) *Province {
	return &Province{
		ID:   province.ID,
		Name: province.Name,
	}
}

// Regency is a struct for regency
type Regency struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetRegencyByIDRequest struct {
	ID int64 `uri:"id" json:"id" binding:"required"`
}

// ProvinceListResponse is a struct for province list response
type RegencyListResponse struct {
	List []*Regency `json:"list"`
	Meta *Meta      `json:"meta"`
}

type GetRegencyRequest struct {
	Search string `form:"search" json:"search"`
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Page   int    `form:"page,default=0" json:"page"`
}

// GetRegencyByProvinceIDRequest is a struct for get regency by province id request
type GetRegencyByProvinceIDRequest struct {
	ProvinceID string `uri:"province_id" json:"province_id" binding:"required"`
}

// NewRegencyResponse create new NewRegencyResponse
func NewRegencyResponse(regency *entity.Regency) *Regency {
	return &Regency{
		ID:   regency.ID,
		Name: regency.Name,
	}
}

// District is a struct for district
type District struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetDistrictRequest struct {
	Search string `form:"search" json:"search"`
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Page   int    `form:"page,default=0" json:"page"`
}

type GetDistrictByIDRequest struct {
	ID int64 `uri:"id" json:"id" binding:"required"`
}

// GetDistrictByRegencyIDRequest is a struct for get district by regency id request
type GetDistrictByRegencyIDRequest struct {
	RegencyID string `uri:"regency_id" json:"regency_id" binding:"required"`
}

// DistrictListResponse is a struct for district list response
type DistrictListResponse struct {
	List []*District `json:"list"`
	Meta *Meta       `json:"meta"`
}

// NewDistrictResponse create new NewDistrictResponse
func NewDistrictResponse(district *entity.District) *District {
	return &District{
		ID:   district.ID,
		Name: district.Name,
	}
}

// Village is a struct for village
type Village struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetVillageByIDRequest struct {
	ID int64 `uri:"id" json:"id" binding:"required"`
}

type GetVillageRequest struct {
	Search string `form:"search" json:"search"`
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Page   int    `form:"page,default=0" json:"page"`
}

// GetVillageByDistrictIDRequest is a struct for get village by district id request
type GetVillageByDistrictIDRequest struct {
	DistrictID string `uri:"district_id" json:"district_id" binding:"required"`
}

// VillageListResponse is a struct for village list response
type VillageListResponse struct {
	List []*Village `json:"list"`
	Meta *Meta      `json:"meta"`
}

// NewVillageResponse create new NewVillageResponse
func NewVillageResponse(village *entity.Village) *Village {
	return &Village{
		ID:   village.ID,
		Name: village.Name,
	}
}
