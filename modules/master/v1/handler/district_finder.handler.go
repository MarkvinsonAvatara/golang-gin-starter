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
type DistrictMasterFinderHandler struct {
	districtMasterFinder service.DistrictMasterFinderUseCase
}

// NewMasterFinderHandler is a constructor for MasterFinderHandler
func DistrictNewMasterFinderHandler(
	districtMasterFinder service.DistrictMasterFinderUseCase,
) *DistrictMasterFinderHandler {
	return &DistrictMasterFinderHandler{
		districtMasterFinder: districtMasterFinder,
	}
}

func (mf *DistrictMasterFinderHandler) GetDistrict(c *gin.Context) {
	var request resource.GetDistrictRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	district, total, err := mf.districtMasterFinder.GetDistricts(c.Request.Context(), request.Search, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	res := make([]*resource.District, 0)

	for _, district := range district {
		res = append(res, resource.NewDistrictResponse(district))
	}

	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:    total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.DistrictListResponse{
		List:  res,
		Meta:  meta,
	}))

}

func (mf *DistrictMasterFinderHandler) GetDistrictByID(c *gin.Context) {
	var request resource.GetDistrictByIDRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	districtInt := strconv.Itoa(int(request.ID))

	districtID, err := uuid.Parse(districtInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	district, err := mf.districtMasterFinder.GetDistrictByID(c, districtID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewDistrictResponse(district)))
}
