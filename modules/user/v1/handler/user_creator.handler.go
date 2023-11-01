package handler

import (
	"gin-starter/common/errors"
	// "gin-starter/middleware"
	"gin-starter/common/interfaces"
	"gin-starter/modules/user/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"gin-starter/utils"
	"log"
	"net/http"
	// "strings"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

// UserCreatorHandler is a handler for user finder
type UserCreatorHandler struct {
	userCreator  service.UserCreatorUseCase
	userFinder  service.UserFinderUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserCreatorHandler is a constructor for UserCreatorHandler
func NewUserCreatorHandler(
	userCreator service.UserCreatorUseCase,
	userFinder service.UserFinderUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserCreatorHandler {
	return &UserCreatorHandler{
		userCreator:  userCreator,
		userFinder:  userFinder,
		cloudStorage: cloudStorage,
	}
}

// CreateUser is a handler for creating user
func (uc *UserCreatorHandler) CreateUser(c *gin.Context) {
	var request resource.CreateUserRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if !utils.IsValidEmail(request.Email) {
		c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "Format Email salah"))
		c.Abort()
		return
	}

	if !utils.IsValidPassword(request.Password) {
		c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "Format Email salah"))
		c.Abort()
		return
	}
	// imagePath, err := uc.cloudStorage.Upload(request.Photo, "users/user/profile")

	// if err != nil {
	// 	parseError := errors.ParseError(err)
	// 	c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
	// 	c.Abort()
	// 	return
	// }

	dob, err := utils.DateStringToTime(request.DOB)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}
	log.Println("dob", dob)

	user, err := uc.userCreator.CreateUser(
		c.Request.Context(),
		request.Name,
		request.Email,
		request.Password,
		request.RoleId,
		dob,
	)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserProfile(user)))
}

// CreateAdmin is a handler for creating admin
func (uc *UserCreatorHandler) CreateAdmin(c *gin.Context) {
	var request resource.CreateAdminRequest
	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	// imagePath, err := uc.cloudStorage.Upload(request.Photo, "users/admin/profile")

	// if err != nil {
	// 	parseError := errors.ParseError(err)
	// 	c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
	// 	c.Abort()
	// 	return
	// }

	dob, _ := utils.DateStringToTime(request.DOB)

	// if err != nil {
	// 	parseError := errors.ParseError(err)
	// 	c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
	// 	c.Abort()
	// 	return
	// }

	// roleID, _ := uuid.Parse(request.RoleId)
	log.Println("masuk")

	// if err != nil {
	// 	parseError := errors.ParseError(err)
	// 	c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
	// 	c.Abort()
	// 	return
	// }
	log.Println("RoleID: ", request.RoleId)
	user, err := uc.userCreator.CreateAdmin(
		c,
		request.Name,
		request.Email,
		request.Password,
		request.RoleId,
		dob,
		request.RoleId,
	)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserProfile(user)))
}

// CreatePermission is a handler for creating permission data
// func (uc *UserCreatorHandler) CreatePermission(c *gin.Context) {
// 	var request resource.CreatePermissionRequest
// 	if err := c.ShouldBind(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
// 		c.Abort()
// 		return
// 	}
// 
// 	permission, err := uc.userCreator.CreatePermission(c, request.Name, request.Label)
// 	if err != nil {
// 		parseError := errors.ParseError(err)
// 		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
// 		c.Abort()
// 		return
// 	}
// 
// 	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewPermissionResponse(permission)))
// }

// CreateRole is a handler for creating role data
// func (uc *UserCreatorHandler) CreateUserRole(c *gin.Context) {
// 	var request resource.CreateUserRoleRequest
// 	if err := c.ShouldBind(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
// 		c.Abort()
// 		return
// 	}
// 
// 	if strings.ToLower(request.Name) == "super admin" {
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "super admin hanya boleh satu!"))
// 		c.Abort()
// 		return
// 	}
// 
// 	// var permissionIDs []uuid.UUID
// 	// if len(request.PermissionIDs) > 0 {
// 	// 	for _, permissionID := range request.PermissionIDs {
// 	// 		valid, err := uuid.Parse(permissionID)
// 	// 		if err != nil {
// 	// 			parseError := errors.ParseError(err)
// 	// 			c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
// 	// 			c.Abort()
// 	// 			return
// 	// 		}
// 	// 		permissionIDs = append(permissionIDs, valid)
// 	// 	}
// 	// }
// 
// 	role, err := uc.userCreator.CreateUserRole(
// 		c,
// 		request.Name,
// 		request.Description,
// 		"Super Admin",
// 	)
// 	if err != nil {
// 		parseError := errors.ParseError(err)
// 		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
// 		c.Abort()
// 		return
// 	}
// 
// 	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserRole(role)))
// }

func (uc *UserCreatorHandler) RegisterUser(c *gin.Context) {
	var request resource.CreateUserRequest
	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	if !utils.IsValidEmail(request.Email) {
		c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "Format Email salah"))
		c.Abort()
		return
	}

	if !utils.IsValidPassword(request.Password) {
		c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "Format Password salah"))
		c.Abort()
		return
	}

	dob, err := utils.DateStringToTime(request.DOB)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}
	log.Println("dob", dob)

	user, err := uc.userCreator.CreateUser(
		c,
		request.Name,
		request.Email,
		request.Password,
		request.RoleId,
		dob,
	)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewUserProfile(user)))
}

// func (uc *UserCreatorHandler) CreatePinjamanRequest(c *gin.Context) {
// 	var request resource.CreatePinjamanRequest
// 	if err := c.ShouldBind(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
// 		c.Abort()
// 		return
// 	}
// 	user, _ := uc.userFinder.GetUserByID(c, middleware.UserID)
// 	userID:= user.ID
// 	userIDString:= userID.String()
// 
// 	TglPinjam, _ := utils.DateStringToTime(request.TglPinjam)
// 	TglKembali,_ := utils.DateStringToTime(request.TglKembali)
// 
// 
// 	pinjaman, err := uc.userCreator.CreatePinjamanRequest(
// 		c.Request.Context(),
// 		userIDString,
// 		request.BukuId,
// 		TglPinjam,
// 		TglKembali,
// 		"User",
// 	)
// 
// 	if err != nil {
// 		parseError := errors.ParseError(err)
// 		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
// 		return
// 	}
// 
// 	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "Pinjaman Sukses direquest", resource.NewPinjamanResponse(pinjaman)))
// 
// }
