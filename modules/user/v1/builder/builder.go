package builder

import (
	"gin-starter/app"
	"gin-starter/config"
	// notificationRepo "gin-starter/modules/notification/v1/repository"
	// notification "gin-starter/modules/notification/v1/service"
	userRepo "gin-starter/modules/user/v1/repository"
	"gin-starter/modules/user/v1/service"
	"gin-starter/sdk/gcs"
	"gin-starter/utils"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// BuildUserHandler builds user handler
// starting from handler down to repository or tool.
func BuildUserHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
	// Cache
	cache := utils.NewClient(redisPool)

	// Repository
	userCreator := userRepo.CreateNewUserRepository(db)
	userFinder := userRepo.FinderNewUserRepository(db)
	userUpdater := userRepo.UpdaterNewUserRepository(db)
	userDeleter := userRepo.DeleteNewUserRepository(db)
	// rr := userRepo.NewRoleRepository(db, cache)
	// urr := userRepo.NewUserRoleRepository(db, cache)
	userRoleCreateRepo := userRepo.CreateNewUserRoleRepository(db, cache)
	userRoleDeleteRepo := userRepo.DeleteNewUserRoleRepository(db, cache)
	userRoleFinderRepo := userRepo.FinderNewUserRoleRepository(db, cache)
	userRoleUpdateRepo := userRepo.UpdateNewUserRoleRepository(db, cache)
	
	// pr := userRepo.NewPermissionRepository(db, cache)
	// pinjamanRepository := userRepo.NewPinjamanRepository(db)
	// nr := notificationRepo.NewNotificationRepository(db)

	// Cloud Storage
	cloudStorage := gcs.NewGoogleCloudStorage(cfg)
	// cloudStorage := aws.NewS3Bucket(cfg, awsSession)

	// Service
	// nc := notification.NewNotificationCreator(cfg, nr)

	userRoleCreator := service.NewUserRoleCreator(cfg, userRoleCreateRepo, cloudStorage)
	userRoleFinder := service.NewUserRoleFinder(cfg, userRoleFinderRepo)
	userRoleUpdater := service.NewUserRoleUpdater(cfg, userRoleUpdateRepo)
	userRoleDeleter := service.NewUserRoleDeleter(cfg, userRoleDeleteRepo)

	uc := service.NewUserCreator(cfg, userCreator,  cloudStorage)
	uf := service.NewUserFinder(cfg, userFinder)
	uu := service.NewUserUpdater(cfg, userUpdater,userFinder)
	ud := service.NewUserDeleter(cfg, userDeleter)
	

	// Handler

	app.UserRoleCreatorHTTPHandler(cfg, router, userRoleCreator, cloudStorage)
	app.UserRoleFinderHTTPHandler(cfg, router, userRoleFinder)
	app.UserRoleUpdaterHTTPHandler(cfg, router, userRoleUpdater, userRoleFinder, cloudStorage)
	app.UserRoleDeleterHTTPHandler(cfg, router, userRoleDeleter, cloudStorage)

	app.UserFinderHTTPHandler(cfg, router, uf)
	app.UserCreatorHTTPHandler(cfg, router, uc, uf, cloudStorage)
	app.UserUpdaterHTTPHandler(cfg, router, uu, uf, cloudStorage)
	app.UserDeleterHTTPHandler(cfg, router, ud, cloudStorage)
}
