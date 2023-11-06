package handler

import (
	"gin-starter/common/errors"
	"gin-starter/middleware"
	// "gin-starter/middleware"
	"gin-starter/modules/pinjaman/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// UserFinderHandler is a handler for user finder
type PinjamanFinderHandler struct {
	pinjamanFinder service.FinderPinjamanFinderUseCase
}

// NewUserFinderHandler is a constructor for UserFinderHandler
func NewPinjamanFinderHandler(
	pinjamanFinder service.FinderPinjamanFinderUseCase,
) *PinjamanFinderHandler {
	return &PinjamanFinderHandler{
		pinjamanFinder: pinjamanFinder,
	}
}

func (uf *PinjamanFinderHandler) GetPinjamanList(c *gin.Context) {
	var request resource.GetPinjamanRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	pinjaman, total, err := uf.pinjamanFinder.GetPinjamanList(c, request.Search, request.Filter, request.Sort, request.Order, request.Limit, request.Page)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.PinjamanDetail, 0)
	for _, p := range pinjaman {
		res = append(res, resource.NewPinjamanResponse(p))
	}

	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:   total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.GetPinjamanListResponse{
		List: res,
		Meta: meta,
	}))
}

func (uf *PinjamanFinderHandler) GetPinjamanByID(c *gin.Context) {
	var request resource.GetPinjamanByIDRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	reqID, err := uuid.Parse(request.ID)

	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}

	pinjaman, err := uf.pinjamanFinder.GetPinjamanByID(c, reqID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewPinjamanResponse(pinjaman)))
}

func (uf *PinjamanFinderHandler) GetPinjamanByUserID(c *gin.Context) {
	userID := middleware.UserID
	userIDString := userID.String()

	pinjaman, err := uf.pinjamanFinder.GetPinjamanByUserID(c, userIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewPinjamanResponse(pinjaman)))
}

func (uf *PinjamanFinderHandler) GetAllList(c *gin.Context) {
	totalAvalaible, totalNotAvalaible, totalUser, totalUserPinjaman, err := uf.pinjamanFinder.GetAllList(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	meta := &resource.DashboardMeta{
		Total_Buku_Tersedia: totalAvalaible,
		Total_Buku_Dipinjam: totalNotAvalaible,
		Total_User:          totalUser,
		Total_User_pinjam:   totalUserPinjaman,
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.GetDashboardListResponse{
		Meta: meta,
	}))
}
