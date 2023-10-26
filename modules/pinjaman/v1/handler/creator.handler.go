package handler

import (
	"gin-starter/common/errors"
	"gin-starter/middleware"
	"gin-starter/common/interfaces"
	"gin-starter/modules/pinjaman/v1/service"
	serviceUser "gin-starter/modules/user/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"gin-starter/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type PinjamanCreatorHandler struct {
	userCreator  service.PinjamanCreatorUseCase
	userFinder  serviceUser.UserFinderUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

func NewPinjamanCreatorHandler(
	userCreator service.PinjamanCreatorUseCase,
	userFinder serviceUser.UserFinderUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *PinjamanCreatorHandler {
	return &PinjamanCreatorHandler{
		userCreator:  userCreator,
		userFinder:  userFinder,
		cloudStorage: cloudStorage,
	}
}

// CreateUser is a handler for creating user

func (uc *PinjamanCreatorHandler) CreatePinjamanRequest(c *gin.Context) {
	var request resource.CreatePinjamanRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	user, _ := uc.userFinder.GetUserByID(c, middleware.UserID)
	userID:= user.ID
	userIDString:= userID.String()

	TglPinjam, _ := utils.DateStringToTime(request.TglPinjam)
	TglKembali,_ := utils.DateStringToTime(request.TglKembali)
	

	pinjaman, err := uc.userCreator.CreatePinjamanRequest(
		c.Request.Context(),
		userIDString,
		request.BukuId,
		TglPinjam,
		TglKembali,
		"User",
	)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Pinjaman Sukses direquest", resource.NewPinjamanResponse(pinjaman)))

}
