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

type VillageMasterFinderHandler struct {
	villageMasterFinder service.VillageMasterFinderUseCase
}

func VillageNewMasterFinderHandler(
	villageMasterFinder service.VillageMasterFinderUseCase,
) *VillageMasterFinderHandler {
	return &VillageMasterFinderHandler{
		villageMasterFinder: villageMasterFinder,
	}
}

func (mf *VillageMasterFinderHandler) GetVillage(c *gin.Context) {
	var request resource.GetVillageRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	village, total, err := mf.villageMasterFinder.GetVillages(c.Request.Context(), request.Search, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	res := make([]*resource.Village, 0)

	for _, village := range village {
		res = append(res, resource.NewVillageResponse(village))
	}

	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:    total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.VillageListResponse{
		List:  res,
		Meta:  meta,
	}))

}

func (mf *VillageMasterFinderHandler) GetVillageByID(c *gin.Context) {
	var request resource.GetVillageByIDRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	
	villageInt:=strconv.Itoa(int(request.ID))

	reqID, err := uuid.Parse(villageInt)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	village, err := mf.villageMasterFinder.GetVillageByID(c,reqID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewVillageResponse(village)))


}