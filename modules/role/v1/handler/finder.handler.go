package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/role/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserFinderHandler is a handler for user finder
type UserFinderHandler struct {
	userFinder service.UserFinderUseCase
}

// NewUserFinderHandler is a constructor for UserFinderHandler
func NewUserFinderHandler(
	userFinder service.UserFinderUseCase,
) *UserFinderHandler {
	return &UserFinderHandler{
		userFinder: userFinder,
	}
}


//GetUser Role is a hand for get list role of user
func (uf *UserFinderHandler)GetUserRoles(c *gin.Context){
	var request resource.GetAdminUsersRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	userRoles, total, err := uf.userFinder.GetUserRoles(c, request.Search, request.Sort, request.Order, request.Limit, request.Page)

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
		Per_Page:     0,
		Current_Page: 0,
		Total_Page:    0,
	}


	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.GetUserRoleRespone{
		List:  res,
		Meta: meta,
	}))
}


func (uf *UserFinderHandler) GetUserRoleByID(c *gin.Context) {
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

	userRole, err := uf.userFinder.GetUserRoleByID(c, roleID)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserRole(userRole)))
}