package app

import (
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/middleware"
	authhandlerv1 "gin-starter/modules/auth/v1/handler"
	authservicev1 "gin-starter/modules/auth/v1/service"
	masterhandlerv1 "gin-starter/modules/master/v1/handler"
	masterservicev1 "gin-starter/modules/master/v1/service"
	notificationhandlerv1 "gin-starter/modules/notification/v1/handler"
	notificationservicev1 "gin-starter/modules/notification/v1/service"
	userhandlerv1 "gin-starter/modules/user/v1/handler"
	userservicev1 "gin-starter/modules/user/v1/service"
	// userRolehandlerv1 "gin-starter/modules/role/v1/handler"
	// userRoleservicev1 "gin-starter/modules/role/v1/service"
	pinjamanHandler1 "gin-starter/modules/pinjaman/v1/handler"
	pinjamanService1 "gin-starter/modules/pinjaman/v1/service"
	"gin-starter/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeprecatedAPI is a handler for deprecated APIs
func DeprecatedAPI(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.ErrorAPIResponse(http.StatusForbidden, "this version of api is deprecated. please use another version."))
	c.Abort()
}

// DefaultHTTPHandler is a handler for default APIs
func DefaultHTTPHandler(cfg config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "invalid route"))
		c.Abort()
	})
}

// AuthHTTPHandler is a handler for auth APIs
func AuthHTTPHandler(cfg config.Config, router *gin.Engine, auc authservicev1.AuthUseCase) {
	hnd := authhandlerv1.NewAuthHandler(auc)
	v1 := router.Group("/v1")
	{
		v1.POST("/user/login", hnd.Login)
		v1.POST("/cms/login", hnd.LoginCMS)
	}
}

// NotificationFinderHTTPHandler is a handler for notification APIs
func NotificationFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationFinderUseCase, nu notificationservicev1.NotificationUpdaterUseCase) {
	hnd := notificationhandlerv1.NewNotificationFinderHandler(cf, nu)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	{
		v1.GET("/user/notifications", hnd.GetNotification)
		v1.GET("/user/notification/count", hnd.CountUnreadNotifications)
	}
}

// NotificationCreatorHTTPHandler is a handler for notification APIs
func NotificationCreatorHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationCreatorUseCase) {
	hnd := notificationhandlerv1.NewNotificationCreatorHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.POST("/cms/notification", hnd.CreateNotification)
	}
}

// NotificationUpdaterHTTPHandler is a handler for notification APIs
func NotificationUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationUpdaterUseCase) {
	hnd := notificationhandlerv1.NewNotificationUpdaterHandler(cf)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	{
		v1.PUT("/user/notification/set", hnd.RegisterUnregisterPlayerID)
		v1.PUT("/user/notification/read", hnd.UpdateReadNotification)
	}
}

func ProvinceFinderHTTPHandler(cfg config.Config, router *gin.Engine, mf masterservicev1.ProvinceMasterFinderUseCase) {
	hnd := masterhandlerv1.ProvinceNewMasterFinderHandler(mf)
	v1 := router.Group("/v1")
	{
		v1.GET("/provinces", hnd.GetProvinces)
		v1.GET("/province/:id", hnd.GetProvinceByID)
	}
}

func RegencyFinderHTTPHandler(cfg config.Config, router *gin.Engine, mf masterservicev1.RegencyMasterFinderUseCase) {
	hnd := masterhandlerv1.RegencyNewMasterFinderHandler(mf)
	v1 := router.Group("/v1")
	{
		v1.GET("/regencies", hnd.GetRegency)
		v1.GET("/regency/:id", hnd.GetRegencyByID)
	}
}

func DistrictFinderHTTPHandler(cfg config.Config, router *gin.Engine, mf masterservicev1.DistrictMasterFinderUseCase) {
	hnd := masterhandlerv1.DistrictNewMasterFinderHandler(mf)
	v1 := router.Group("/v1")
	{
		v1.GET("/districts", hnd.GetDistrict)
		v1.GET("/district/:id", hnd.GetDistrictByID)
	}
}


func VillageFinderHTTPHandler(cfg config.Config, router *gin.Engine, mf masterservicev1.VillageMasterFinderUseCase) {
	hnd := masterhandlerv1.VillageNewMasterFinderHandler(mf)
	v1 := router.Group("/v1")
	{
		v1.GET("/villages", hnd.GetVillage)
		v1.GET("/village/:id", hnd.GetVillageByID)
	}
}

// MasterFinderHTTPHandler is a handler for master APIs
func BookMasterFinderHTTPHandler(cfg config.Config, router *gin.Engine, mf masterservicev1.BookMasterFinderUseCase) {
	hnd := masterhandlerv1.NewMasterFinderHandler(mf)
	v1 := router.Group("/v1")
	{
		// v1.GET("/regencies/:province_id", hnd.GetRegenciesByProvinceID)
		// v1.GET("/districts/:regency_id", hnd.GetDistrictsByRegencyID)
		// v1.GET("/villages/:district_id", hnd.GetVillagesByDistrictID)
		v1.GET("/books", hnd.GetBooks)
		v1.GET("/book/:id", hnd.GetBookByID)
	}
}

// MasterCreatorHTTPHandler is a handler for master APIs
func BookMasterCreatorHTTPHandler(cfg config.Config, router *gin.Engine, mc masterservicev1.BookMasterCreatorUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := masterhandlerv1.NewMasterCreatorHandler(mc, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Admin(cfg))
	{
		v1.POST("/book", hnd.CreateBook)
	}
}

// MasterDeleterHTTPHandler is a handler for master APIs
func BookMasterDeleterHTTPHandler(cfg config.Config, router *gin.Engine, md masterservicev1.BookMasterDeleterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := masterhandlerv1.NewMasterDeleterHandler(md, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Admin(cfg))
	{
		v1.DELETE("/book/:id", hnd.DeleteBook)
	}
}

// MasterUpdaterHTTPHandler is a handler for master APIs
func BookMasterUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, mu masterservicev1.BookMasterUpdaterUseCase, masterFinder masterservicev1.BookMasterFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := masterhandlerv1.NewMasterUpdaterHandler(mu, masterFinder, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/book/:id", hnd.UpdateBook)
	}
}

// UserFinderHTTPHandler is a handler for user APIs
func UserFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf userservicev1.UserFinderUseCase) {
	hnd := userhandlerv1.NewUserFinderHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.GET("/user/forgot-password/profile/:token", hnd.GetUserByForgotPasswordToken)
	}

	v1.Use(middleware.Auth(cfg))
	{
		v1.GET("/user/profile", hnd.GetUserProfile)
	}

	v1.Use(middleware.Admin(cfg))
	{
		v1.GET("/cms/profile", hnd.GetAdminProfile)
		v1.GET("/cms/admin/list", hnd.GetAdminUsers)
		v1.GET("/cms/admin/detail/:id", hnd.GetAdminUserByID)
		v1.GET("/cms/user/list", hnd.GetUsers)
		v1.GET("/cms/user/detail/:id", hnd.GetUserByID)
		// v1.GET("/cms/permission", hnd.GetPermissions)
		// v1.GET("/cms/user/permission", hnd.GetUserPermissions)
		// v1.GET("/cms/role", hnd.GetUserRoles)
		// v1.GET("/cms/user/role/:id", hnd.GetUserRoleByID)
		// v1.GET("/cms/pinjaman/list", hnd.GetPinjamanList)
		// v1.GET("/cms/pinjaman/detail/:id", hnd.GetPinjamanByID)
	}
}

// UserCreatorHTTPHandler is a handler for user APIs
func UserCreatorHTTPHandler(cfg config.Config, router *gin.Engine, uc userservicev1.UserCreatorUseCase, uf userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserCreatorHandler(uc, uf, cloudStorage)
	v1 := router.Group("/v1")
	{
		v1.POST("/user/register", hnd.RegisterUser)
	}

	v1.Use(middleware.Auth(cfg))
	{
		// v1.POST("/user/pinjam", hnd.CreatePinjamanRequest)
	}

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.POST("/cms/user", hnd.CreateUser)
		v1.POST("/cms/admin/user", hnd.CreateAdmin)
		// v1.POST("/cms/role", hnd.CreateUserRole)
		// v1.POST("/cms/permission", hnd.CreatePermission)
	}
}

// UserUpdaterHTTPHandler is a handler for user APIs
func UserUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, uu userservicev1.UserUpdaterUseCase, uf userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserUpdaterHandler(uu, uf, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.PUT("/user/forgot-password/request", hnd.ForgotPasswordRequest)
		v1.PUT("/user/forgot-password", hnd.ForgotPassword)
	}

	v1.Use(middleware.Auth(cfg))
	{
		v1.PUT("/user/profile", hnd.UpdateUser)
		v1.PUT("/user/password", hnd.ChangePassword)
	}

	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/cms/profile/:id", hnd.UpdateUser)
		v1.PUT("/cms/admin/:id", hnd.UpdateAdmin)
		// v1.PUT("/cms/role/:id", hnd.UpdateUserRole)
		// v1.PUT("/cms/user/activate/:id", hnd.ActivateDeactivateUser)
		// v1.PUT("/cms/permission/:id", hnd.UpdatePermission)
		// v1.PUT("/cms/pinjaman/:id", hnd.HandledPinjaman)
	}
}

// UserDeleterHTTPHandler is a handler for user APIs
func UserDeleterHTTPHandler(cfg config.Config, router *gin.Engine, ud userservicev1.UserDeleterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserDeleterHandler(ud, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.DELETE("/cms/user/:id", hnd.DeleteUsers)
		v1.DELETE("/cms/admin/:id", hnd.DeleteAdmin)
		// v1.DELETE("/cms/role/:id", hnd.DeleteUserRole)
	}
}

func UserRoleFinderHTTPHandler(cfg config.Config, router *gin.Engine, userRoleFinder userservicev1.UserRoleFinderUseCase) {
	hnd := userhandlerv1.NewUserRoleFinderHandler(userRoleFinder)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.GET("/cms/role", hnd.GetUserRoles)
		v1.GET("/cms/role/:id", hnd.GetUserRoleByID)
	}
}

func UserRoleCreatorHTTPHandler(cfg config.Config, router *gin.Engine, userRoleCreate userservicev1.UserRoleCreatorUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserRoleCreatorHandler(userRoleCreate, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.POST("/cms/role", hnd.CreateUserRole)
	}
}

func UserRoleUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, userRoleUpdate userservicev1.UserRoleUpdaterUseCase, userRoleFinder userservicev1.UserRoleFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserRoleUpdaterHandler(userRoleUpdate, userRoleFinder, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/cms/role/:id", hnd.UpdateUserRole)
	}
}

func UserRoleDeleterHTTPHandler(cfg config.Config, router *gin.Engine, userRoleDelete userservicev1.UserRoleDeleterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserRoleDeleterHandler(userRoleDelete, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.DELETE("/cms/role/:id", hnd.DeleteUserRole)
	}
}



// func UserRoleFinderHTTPHandler(cfg config.Config, router *gin.Engine,userRoleFinder userRoleservicev1.UserFinderUseCase) {
// 	hnd := userRolehandlerv1.NewUserFinderHandler(userRoleFinder)
// 	v1 := router.Group("/v1")
//
// 	v1.Use(middleware.Auth(cfg))
// 	v1.Use(middleware.Admin(cfg))
// 	{
// 		v1.GET("/cms/role", hnd.GetUserRoles)
// 		v1.GET("/cms/role/:id", hnd.GetUserRoleByID)
// 	}
// }

// func UserRoleCreatorHTTPHandler(cfg config.Config, router *gin.Engine, userRoleCreate userRoleservicev1.UserCreatorUseCase, userRoleFinder userRoleservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
// 	hnd := userRolehandlerv1.NewUserCreatorHandler(userRoleCreate, cloudStorage)
// 	v1 := router.Group("/v1")
//
// 	v1.Use(middleware.Auth(cfg))
// 	v1.Use(middleware.Admin(cfg))
// 	{
// 		v1.POST("/cms/role", hnd.CreateUserRole)
// 	}
// }

// func UserRoleUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, userRoleUpdate userRoleservicev1.UserUpdaterUseCase, userRoleFinder userRoleservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
// 	hnd := userRolehandlerv1.NewUserUpdaterHandler(userRoleUpdate, userRoleFinder, cloudStorage)
// 	v1 := router.Group("/v1")
//
// 	v1.Use(middleware.Auth(cfg))
// 	v1.Use(middleware.Admin(cfg))
// 	{
// 		v1.PUT("/cms/role/:id", hnd.UpdateUserRole)
// 	}
// }

// func UserRoleDeleterHTTPHandler(cfg config.Config, router *gin.Engine, userRoleDelete userRoleservicev1.UserDeleterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
// 	hnd := userRolehandlerv1.NewUserDeleterHandler(userRoleDelete, cloudStorage)
// 	v1 := router.Group("/v1")
//
// 	v1.Use(middleware.Auth(cfg))
// 	v1.Use(middleware.Admin(cfg))
// 	{
// 		v1.DELETE("/cms/role/:id", hnd.DeleteUserRole)
// 	}
// }

func PinjamanFinderHTTPHandler(cfg config.Config, router *gin.Engine, pinjamanFinder pinjamanService1.FinderPinjamanFinderUseCase) {
	hnd := pinjamanHandler1.NewPinjamanFinderHandler(pinjamanFinder)
	v1 := router.Group("/v1")

	v1.Use(middleware.Admin(cfg))
	{
		v1.GET("/cms/pinjaman/list", hnd.GetPinjamanList)
		v1.GET("/cms/pinjaman/detail/:id", hnd.GetPinjamanByID)
	}
}

func PinjamanCreatorHTTPHandler(cfg config.Config, router *gin.Engine, pinjamanCreate pinjamanService1.CreatePinjamanCreatorUseCase, cf userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := pinjamanHandler1.NewPinjamanCreatorHandler(pinjamanCreate, cf, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	{
		v1.POST("/user/pinjam", hnd.CreatePinjamanRequest)
	}
}

func PinjamanUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, pinjamanUpdate pinjamanService1.PinjamanUpdaterUseCase, pinjamanFinder pinjamanService1.FinderPinjamanFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := pinjamanHandler1.NewPinjamanUpdaterHandler(pinjamanUpdate, pinjamanFinder, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/cms/pinjaman/:id", hnd.HandledPinjaman)
	}
}

// // UserCreatorHTTPHandler is a handler for user APIs
// func UserCreatorHTTPHandler(cfg config.Config, router *gin.Engine, uc userservicev1.UserCreatorUseCase, uf userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
// 	hnd := userhandlerv1.NewUserCreatorHandler(uc, uf,cloudStorage)
// 	v1 := router.Group("/v1")
// 	{
// 		v1.POST("/user/register", hnd.RegisterUser)
// 	}
//
// 	v1.Use(middleware.Auth(cfg))
// 	{
// 		v1.POST("/user/pinjam", hnd.CreatePinjamanRequest)
// 	}
//
// 	v1.Use(middleware.Auth(cfg))
// 	v1.Use(middleware.Admin(cfg))
// 	{
// 		v1.POST("/cms/user", hnd.CreateUser)
// 		v1.POST("/cms/admin/user", hnd.CreateAdmin)
// 		v1.POST("/cms/role", hnd.CreateUserRole)
// 		// v1.POST("/cms/permission", hnd.CreatePermission)
// 	}
// }

// // UserUpdaterHTTPHandler is a handler for user APIs
// func UserUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, uu userservicev1.UserUpdaterUseCase, uf userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
// 	hnd := userhandlerv1.NewUserUpdaterHandler(uu, uf, cloudStorage)
// 	v1 := router.Group("/v1")
//
// 	v1.Use(middleware.Auth(cfg))
//
// 	{
// 		v1.PUT("/user/forgot-password/request", hnd.ForgotPasswordRequest)
// 		v1.PUT("/user/forgot-password", hnd.ForgotPassword)
// 	}
//
// 	v1.Use(middleware.Auth(cfg))
// 	{
// 		v1.PUT("/user/profile", hnd.UpdateUser)
// 		v1.PUT("/user/password", hnd.ChangePassword)
// 	}
//
// 	v1.Use(middleware.Admin(cfg))
// 	{
// 		v1.PUT("/cms/profile/:id", hnd.UpdateUser)
// 		v1.PUT("/cms/admin/:id", hnd.UpdateAdmin)
// 		v1.PUT("/cms/role/:id", hnd.UpdateUserRole)
// 		v1.PUT("/cms/user/activate/:id", hnd.ActivateDeactivateUser)
// 		v1.PUT("/cms/permission/:id", hnd.UpdatePermission)
// 		v1.PUT("/cms/pinjaman/:id", hnd.HandledPinjaman)
// 	}
// }
//
// // UserDeleterHTTPHandler is a handler for user APIs
// func UserDeleterHTTPHandler(cfg config.Config, router *gin.Engine, ud userservicev1.UserDeleterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
// 	hnd := userhandlerv1.NewUserDeleterHandler(ud, cloudStorage)
// 	v1 := router.Group("/v1")
//
// 	v1.Use(middleware.Auth(cfg))
// 	v1.Use(middleware.Admin(cfg))
// 	{
// 		v1.DELETE("/cms/user/:id", hnd.DeleteUsers)
// 		v1.DELETE("/cms/admin/:id", hnd.DeleteAdmin)
// 		v1.DELETE("/cms/role/:id", hnd.DeleteUserRole)
// 	}
// }
