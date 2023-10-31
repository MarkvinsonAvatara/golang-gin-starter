package handler

import (
	"gin-starter/common/errors"
	// "gin-starter/middleware"
	"gin-starter/modules/user/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserFinderHandler is a handler for user finder
type UserRoleFinderHandler struct {
	userRoleFinder service.UserRoleFinderUseCase
}

// NewUserFinderHandler is a constructor for UserFinderHandler
func NewUserRoleFinderHandler(
	userRoleFinder service.UserRoleFinderUseCase,
) *UserRoleFinderHandler {
	return &UserRoleFinderHandler{
		userRoleFinder: userRoleFinder,
	}
}


// GetUser Role is a hand for get list role of user
func (uf *UserRoleFinderHandler) GetUserRoles(c *gin.Context) {
	var request resource.GetUserRoleRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	userRoles, total, err := uf.userRoleFinder.GetUserRoles(c, request.Search, request.Sort, request.Order, request.Limit, request.Page)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	res := make([]*resource.UserRole, 0)

	for _, u := range userRoles {
		res = append(res, resource.NewUserRole(u))
	}

	meta := &resource.Meta{
		Total_Data:   total,
		Per_Page:     request.Limit,
		Current_Page: request.Page,
		Total_Page:    total / int64(request.Limit),
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.GetUserRoleRespone{
		List:  res,
		Meta: meta,
	}))
}

func (uf *UserRoleFinderHandler) GetUserRoleByID(c *gin.Context) {
	var request resource.GetUserRoleByID

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	roleID, err := uuid.Parse(request.ID)

	if err != nil {
		c.JSON(errors.ErrInvalidArgument.Code, response.ErrorAPIResponse(errors.ErrInvalidArgument.Code, errors.ErrInvalidArgument.Message))
		c.Abort()
		return
	}

	userRole, err := uf.userRoleFinder.GetUserRoleByID(c, roleID)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserRole(userRole)))
}

