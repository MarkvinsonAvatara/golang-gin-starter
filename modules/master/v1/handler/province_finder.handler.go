package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/master/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MasterFinderHandler is a handler for master finder
type ProvinceMasterFinderHandler struct {
	provinceMasterFinder service.ProvinceMasterFinderUseCase
}

// NewMasterFinderHandler is a constructor for MasterFinderHandler
func ProvinceNewMasterFinderHandler(
	provinceMasterFinder service.ProvinceMasterFinderUseCase,
) *ProvinceMasterFinderHandler {
	return &ProvinceMasterFinderHandler{
		provinceMasterFinder: provinceMasterFinder,
	}
}

// GetProvinces is a handler for getting all provinces
func (mf *ProvinceMasterFinderHandler) GetProvinces(c *gin.Context) {
	var request resource.GetProvinceRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	provinces, total, err := mf.provinceMasterFinder.GetProvinces(c.Request.Context(), request.Search, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	res := make([]*resource.Province, 0)
	
	for _, province := range provinces {
		res = append(res, resource.NewProvinceResponse(province))
	}
	
	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:    total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.ProvinceListResponse{
		List:  res,
		Meta:  meta,
	}))
}

func (mf *ProvinceMasterFinderHandler) GetProvinceByID(c *gin.Context) {
	var request resource.GetProvinceByIDRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	provinceInt:= strconv.Itoa(int(request.ID))
	provinceID, err := uuid.Parse(provinceInt)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	province, err := mf.provinceMasterFinder.GetProvinceByID(c, provinceID)

	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewProvinceResponse(province)))
}





