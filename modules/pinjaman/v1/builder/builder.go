package builder

import (
	"gin-starter/app"
	"gin-starter/config"
	userRepo "gin-starter/modules/role/v1/repository"
	"gin-starter/modules/role/v1/service"
	"gin-starter/sdk/gcs"
	"gin-starter/utils"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// BuildUserHandler builds user handler
// starting from handler down to repository or tool.
func BuildUserRoleHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
	// Cache
	cache := utils.NewClient(redisPool)

	// Repository
	UserRoleRepository := userRepo.NewUserRoleRepository(db, cache)

	// Cloud Storage
	cloudStorage := gcs.NewGoogleCloudStorage(cfg)
	// cloudStorage := aws.NewS3Bucket(cfg, awsSession)

	// Service
	userRoleCreator := service.NewUserRoleCreator(cfg, UserRoleRepository, cloudStorage)
	userRoleFinder := service.NewUserRoleFinder(cfg, UserRoleRepository)
	userRoleUpdater := service.NewUserRoleUpdater(cfg, UserRoleRepository)
	userRoleDeleter := service.NewUserRoleDeleter(cfg, UserRoleRepository)

	// Handler
	app.UserRoleFinderHTTPHandler(cfg, router, userRoleFinder)
	app.UserRoleCreatorHTTPHandler(cfg, router, userRoleCreator, userRoleFinder, cloudStorage)
	app.UserRoleUpdaterHTTPHandler(cfg, router, userRoleUpdater, userRoleFinder, cloudStorage)
	app.UserRoleDeleterHTTPHandler(cfg, router, userRoleDeleter, cloudStorage)
}
