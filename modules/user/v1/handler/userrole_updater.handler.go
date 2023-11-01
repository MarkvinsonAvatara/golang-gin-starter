package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/entity"
	// "gin-starter/middleware"
	"gin-starter/modules/user/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	// "gin-starter/utils"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserUpdaterHandler is a handler for user updater
type UserRoleUpdaterHandler struct {
	userRoleUpdater  service.UserRoleUpdaterUseCase
	userRoleFinder   service.UserRoleFinderUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserUpdaterHandler is a constructor for UserUpdaterHandler
func NewUserRoleUpdaterHandler(
	userRoleUpdater service.UserRoleUpdaterUseCase,
	userRoleFinder service.UserRoleFinderUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserRoleUpdaterHandler {
	return &UserRoleUpdaterHandler{
		userRoleUpdater:  userRoleUpdater,
		userRoleFinder:   userRoleFinder,
		cloudStorage: cloudStorage,
	}
}


func (uu *UserRoleUpdaterHandler) UpdateUserRole(c *gin.Context) {
	var request resource.UpdateUserRoleRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	userRoleIDstr := c.Param("id")
	userRoleID, err := uuid.Parse(userRoleIDstr)
	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}

	_, err = uu.userRoleFinder.GetUserRoleByID(c, userRoleID)

	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}
	userRole := entity.UpdateUserRole(
		userRoleID,
		request.Name,
		request.Description,
		"Admin",
	)

	if err := uu.userRoleUpdater.UpdateUserRoles(c, userRole); err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Update Success", nil))

}

