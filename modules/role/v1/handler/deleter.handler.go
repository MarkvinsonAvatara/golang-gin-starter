package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/modules/role/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// UserDeleterHandler is a handler for user finder
type UserDeleterHandler struct {
	userDeleter  service.UserDeleterUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserDeleterHandler is a constructor for UserDeleterHandler
func NewUserDeleterHandler(
	userDeleter service.UserDeleterUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserDeleterHandler {
	return &UserDeleterHandler{
		userDeleter:  userDeleter,
		cloudStorage: cloudStorage,
	}
}


//DeleteUserRole is a handler for delete user role
func (ud *UserDeleterHandler) DeleteUserRole(c*gin.Context){
	var request resource.DeleteUserRoleRequest

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

	if err := ud.userDeleter.DeleteUserRole(c, reqID); err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success role berhasil terhapus", nil))
}
