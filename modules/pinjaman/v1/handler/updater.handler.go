package handler

import (
	"gin-starter/common/errors"
	"gin-starter/common/interfaces"
	"gin-starter/entity"
	"gin-starter/modules/role/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserUpdaterHandler is a handler for user updater
type UserUpdaterHandler struct {
	userUpdater  service.UserUpdaterUseCase
	userFinder   service.UserFinderUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserUpdaterHandler is a constructor for UserUpdaterHandler
func NewUserUpdaterHandler(
	userUpdater service.UserUpdaterUseCase,
	userFinder service.UserFinderUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserUpdaterHandler {
	return &UserUpdaterHandler{
		userUpdater:  userUpdater,
		userFinder:   userFinder,
		cloudStorage: cloudStorage,
	}
}

func (uu *UserUpdaterHandler) UpdateUserRole(c *gin.Context){
	var request resource.UpdateUserRoleRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}
	userRoleIDstr:= c.Param("id")
	userRoleID, err := uuid.Parse(userRoleIDstr)
	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}

	_, err = uu.userFinder.GetUserRoleByID(c, userRoleID)
	
	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}
	userRole := entity.NewUserRole(
		userRoleID,
		request.Name,
		request.Description,
		"Super Admin",
	)

	if err := uu.userUpdater.UpdateUserRoles(c, userRole); err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))	
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Update Success", nil))

}


