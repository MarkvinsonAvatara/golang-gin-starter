package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/modules/role/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"
	// "strings"

	"github.com/gin-gonic/gin"
)

// UserCreatorHandler is a handler for user finder
type UserCreatorHandler struct {
	userCreator  service.UserCreatorUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserCreatorHandler is a constructor for UserCreatorHandler
func NewUserCreatorHandler(
	userCreator service.UserCreatorUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserCreatorHandler {
	return &UserCreatorHandler{
		userCreator:  userCreator,
		cloudStorage: cloudStorage,
	}
}

// CreateRole is a handler for creating role data
func (uc *UserCreatorHandler) CreateUserRole(c *gin.Context) {
	var request resource.CreateUserRoleRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	// if strings.ToLower(request.Name) == "super admin" {
	// 	c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "super admin hanya boleh satu!"))
	// 	c.Abort()
	// 	return
	// }

	role, err := uc.userCreator.CreateUserRole(
		c,
		request.Name,
		request.Description,
		"Admin",
	)
	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserRole(role)))
}
