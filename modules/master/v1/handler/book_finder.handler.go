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
type BookMasterFinderHandler struct {
	bookMasterFinder service.BookMasterFinderUseCase
}

// NewMasterFinderHandler is a constructor for MasterFinderHandler
func NewMasterFinderHandler(
	bookMasterFinder service.BookMasterFinderUseCase,
) *BookMasterFinderHandler {
	return &BookMasterFinderHandler{
		bookMasterFinder: bookMasterFinder,
	}
}


// // GetRegenciesByProvinceID is a handler for getting all regencies by province id
// func (mf *BookMasterFinderHandler) GetRegenciesByProvinceID(c *gin.Context) {
// 	var req resource.GetRegencyByProvinceIDRequest
// 	if err := c.ShouldBindUri(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
// 		c.Abort()
// 		return
// 	}

// 	regencies, err := mf.bookMasterFinder.GetRegencies(c.Request.Context(), req.ProvinceID)
// 	if err != nil {
// 		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
// 		c.Abort()
// 		return
// 	}

// 	res := make([]*resource.Regency, 0)

// 	for _, regency := range regencies {
// 		res = append(res, resource.NewRegencyResponse(regency))
// 	}

// 	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.RegencyListResponse{
// 		List:  res,
// 		Meta: nil,
// 	}))
// }

// // GetDistrictsByRegencyID is a handler for getting all districts by regency id
// func (mf *BookMasterFinderHandler) GetDistrictsByRegencyID(c *gin.Context) {
// 	var req resource.GetDistrictByRegencyIDRequest
// 	if err := c.ShouldBindUri(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
// 		c.Abort()
// 		return
// 	}

// 	districts, err := mf.bookMasterFinder.GetDistricts(c.Request.Context(), req.RegencyID)
// 	if err != nil {
// 		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
// 		c.Abort()
// 		return
// 	}

// 	res := make([]*resource.District, 0)

// 	for _, district := range districts {
// 		res = append(res, resource.NewDistrictResponse(district))
// 	}

// 	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.DistrictListResponse{
// 		List:  res,
// 		Total: int64(len(res)),
// 	}))
// }

// // GetVillagesByDistrictID is a handler for getting all villages by district id
// func (mf *BookMasterFinderHandler) GetVillagesByDistrictID(c *gin.Context) {
// 	var req resource.GetVillageByDistrictIDRequest
// 	if err := c.ShouldBindUri(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
// 		c.Abort()
// 		return
// 	}

// 	villages, err := mf.bookMasterFinder.GetVillages(c.Request.Context(), req.DistrictID)
// 	if err != nil {
// 		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
// 		c.Abort()
// 		return
// 	}

// 	res := make([]*resource.Village, 0)

// 	for _, village := range villages {
// 		res = append(res, resource.NewVillageResponse(village))
// 	}

// 	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.VillageListResponse{
// 		List:  res,
// 		Total: int64(len(res)),
// 	}))
// }

// GetBooks is a handler for getting all books
func (mf *BookMasterFinderHandler) GetBooks(c *gin.Context) {
	var request resource.GetBookRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	books, total, err := mf.bookMasterFinder.GetBooks(c, request.Search, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.BookDetail, 0)

	for _, book := range books {
		res = append(res, resource.NewBookResponse(book))
	}

	// offset := (request.Page - 1) * request.Limit

	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:    total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.GetBookListResponse{
		List: res,
		Meta: meta,
	}))
}

// GetBookByID is a handler for getting book by id
func (mf *BookMasterFinderHandler) GetBookByID(c *gin.Context) {
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

	book, err := mf.bookMasterFinder.GetBookByID(c, reqID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewBookResponse(book)))
}

// GetBookAvalaibily is a handler for getting book by id
func (mf *BookMasterFinderHandler) GetBookAvalaibily(c *gin.Context) {
	var request resource.GetBookRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	books, total, err := mf.bookMasterFinder.GetBookAvalaibily(c, request.Search, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.BookDetail, 0)

	for _, book := range books {
		res = append(res, resource.NewBookResponse(book))
	}

	// offset := (request.Page - 1) * request.Limit

	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:    total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.GetBookListResponse{
		List: res,
		Meta: meta,
	}))
}