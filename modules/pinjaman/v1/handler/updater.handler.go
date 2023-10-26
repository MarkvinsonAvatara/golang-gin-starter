package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/entity"
	// "gin-starter/middleware"
	"gin-starter/modules/pinjaman/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	// "gin-starter/utils"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserUpdaterHandler is a handler for user updater
type PinjamanUpdaterHandler struct {
	userUpdater  service.PinjamanUpdaterUseCase
	userFinder   service.PinjamanFinderUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserUpdaterHandler is a constructor for UserUpdaterHandler
func NewPinjamanUpdaterHandler(
	userUpdater service.PinjamanUpdaterUseCase,
	userFinder service.PinjamanFinderUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *PinjamanUpdaterHandler {
	return &PinjamanUpdaterHandler{
		userUpdater:  userUpdater,
		userFinder:   userFinder,
		cloudStorage: cloudStorage,
	}
}


func (uu *PinjamanUpdaterHandler) HandledPinjaman(c *gin.Context) {
	var request resource.HandledRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	pinjamanIDstr := c.Param("id")
	pinjamanID, err := uuid.Parse(pinjamanIDstr)
	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}
	_, err = uu.userFinder.GetPinjamanByID(c, pinjamanID)
	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}


	pinjaman := entity.HandledPinjaman(
		pinjamanID,
		"admin",
		request.Status,

	)

	if err := uu.userUpdater.HandledPinjaman(c, pinjaman); err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Penangan Pinjaman Success", nil))
}
