package handler

import (
	"gin-starter/common/errors"
	// "gin-starter/middleware"
	"gin-starter/common/interfaces"
	"gin-starter/modules/user/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	// "gin-starter/utils"
	// "log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

// UserCreatorHandler is a handler for user finder
type UserRoleCreatorHandler struct {
	userRoleCreator  service.UserRoleCreatorUseCase
	// userFinder  service.UserFinderUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserCreatorHandler is a constructor for UserCreatorHandler
func NewUserRoleCreatorHandler(
	userRoleCreator service.UserRoleCreatorUseCase,
	// userFinder service.UserFinderUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserRoleCreatorHandler {
	return &UserRoleCreatorHandler{
		userRoleCreator:  userRoleCreator,
		// userFinder:  userFinder,
		cloudStorage: cloudStorage,
	}
}


func (uc *UserRoleCreatorHandler) CreateUserRole(c *gin.Context) {
	var request resource.CreateUserRoleRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if strings.ToLower(request.Name) == "super admin" {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "super admin hanya boleh satu!"))
		c.Abort()
		return
	}


	role, err := uc.userRoleCreator.CreateUserRole(
		c,
		request.Name,
		request.Description,
		"Super Admin",
	)
	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserRole(role)))
}

