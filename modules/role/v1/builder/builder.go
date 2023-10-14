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
func BuildUserHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
	// Cache
	cache := utils.NewClient(redisPool)

	// Repository
	urr := userRepo.NewUserRoleRepository(db, cache)

	// Cloud Storage
	cloudStorage := gcs.NewGoogleCloudStorage(cfg)
	// cloudStorage := aws.NewS3Bucket(cfg, awsSession)

	// Service
	userCreator := service.NewUserCreator(cfg, urr, cloudStorage)
	userRoleFinder := service.NewUserFinder(cfg, urr)
	userUpdater := service.NewUserUpdater(cfg, urr)
	userDeleter := service.NewUserDeleter(cfg, urr)

	// Handler
	app.UserFinderHTTPHandler(cfg, router, userRoleFinder)
	app.UserCreatorHTTPHandler(cfg, router, userCreator, userRoleFinder, cloudStorage)
	app.UserUpdaterHTTPHandler(cfg, router, userUpdater, userRoleFinder, cloudStorage)
	app.UserDeleterHTTPHandler(cfg, router, userDeleter, cloudStorage)
}
