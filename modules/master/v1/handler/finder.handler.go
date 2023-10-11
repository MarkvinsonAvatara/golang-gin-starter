package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/master/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MasterFinderHandler is a handler for master finder
type MasterFinderHandler struct {
	masterFinder service.MasterFinderUseCase
}

// NewMasterFinderHandler is a constructor for MasterFinderHandler
func NewMasterFinderHandler(
	masterFinder service.MasterFinderUseCase,
) *MasterFinderHandler {
	return &MasterFinderHandler{
		masterFinder: masterFinder,
	}
}

// GetProvinces is a handler for getting all provinces
func (mf *MasterFinderHandler) GetProvinces(c *gin.Context) {
	provinces, err := mf.masterFinder.GetProvinces(c.Request.Context())
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.Province, 0)

	for _, province := range provinces {
		res = append(res, resource.NewProvinceResponse(province))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.ProvinceListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}

// GetRegenciesByProvinceID is a handler for getting all regencies by province id
func (mf *MasterFinderHandler) GetRegenciesByProvinceID(c *gin.Context) {
	var req resource.GetRegencyByProvinceIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	regencies, err := mf.masterFinder.GetRegencies(c.Request.Context(), req.ProvinceID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.Regency, 0)

	for _, regency := range regencies {
		res = append(res, resource.NewRegencyResponse(regency))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.RegencyListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}

// GetDistrictsByRegencyID is a handler for getting all districts by regency id
func (mf *MasterFinderHandler) GetDistrictsByRegencyID(c *gin.Context) {
	var req resource.GetDistrictByRegencyIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	districts, err := mf.masterFinder.GetDistricts(c.Request.Context(), req.RegencyID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.District, 0)

	for _, district := range districts {
		res = append(res, resource.NewDistrictResponse(district))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.DistrictListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}

// GetVillagesByDistrictID is a handler for getting all villages by district id
func (mf *MasterFinderHandler) GetVillagesByDistrictID(c *gin.Context) {
	var req resource.GetVillageByDistrictIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	villages, err := mf.masterFinder.GetVillages(c.Request.Context(), req.DistrictID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.Village, 0)

	for _, village := range villages {
		res = append(res, resource.NewVillageResponse(village))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.VillageListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}

// GetBooks is a handler for getting all books
func (mf *MasterFinderHandler) GetBooks(c *gin.Context) {
	books, err := mf.masterFinder.GetBooks(c.Request.Context())
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.BookDetail, 0)

	for _, book := range books {
		res = append(res, resource.NewBookResponse(book))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.GetBookListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}

// GetBookByID is a handler for getting book by id
func (mf *MasterFinderHandler) GetBookByID(c *gin.Context) {
	var request resource.GetBookByIDRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	reqID, err := uuid.Parse(request.ID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	book, err := mf.masterFinder.GetBookByID(c, reqID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewBookResponse(book)))
}
