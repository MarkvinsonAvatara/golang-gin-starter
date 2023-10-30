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
type RegencyMasterFinderHandler struct {
	regencyMasterFinder service.RegencyMasterFinderUseCase
}

// NewMasterFinderHandler is a constructor for MasterFinderHandler
func RegencyNewMasterFinderHandler(
	regencyMasterFinder service.RegencyMasterFinderUseCase,
) *RegencyMasterFinderHandler {
	return &RegencyMasterFinderHandler{
		regencyMasterFinder: regencyMasterFinder,
	}
}

func (mf *RegencyMasterFinderHandler) GetRegency(c *gin.Context) {
	var request resource.GetRegencyRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	regency, total, err := mf.regencyMasterFinder.GetRegencies(c.Request.Context(), request.Search, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	res := make([]*resource.Regency, 0)
	
	for _, regency := range regency {
		res = append(res, resource.NewRegencyResponse(regency))
	}
	
	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:    total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.RegencyListResponse{
		List:  res,
		Meta:  meta,
	}))
}

func (mf *RegencyMasterFinderHandler) GetRegencyByID(c *gin.Context) {
	var request resource.GetProvinceByIDRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	
	regencyInt:=strconv.Itoa(int(request.ID))

	regencyID, err := uuid.Parse(regencyInt)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	regency, err := mf.regencyMasterFinder.GetRegencyByID(c, regencyID)

	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewRegencyResponse(regency)))
}





