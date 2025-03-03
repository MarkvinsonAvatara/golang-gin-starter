package builder

import (
	"gin-starter/app"
	"gin-starter/config"
	// notificationRepo "gin-starter/modules/notification/v1/repository"
	// notification "gin-starter/modules/notification/v1/service"
	pinjamanRepo "gin-starter/modules/pinjaman/v1/repository"
	"gin-starter/modules/pinjaman/v1/service"
	userRepo "gin-starter/modules/user/v1/repository"
	userService "gin-starter/modules/user/v1/service"
	"gin-starter/sdk/gcs"
	// "gin-starter/utils"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// BuildUserHandler builds user handler
// starting from handler down to repository or tool.
func BuildPinjamanHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
	// Cache
	// cache := utils.NewClient(redisPool)

	// Repository
	// ur := userRepo.NewUserRepository(db)
	ur := userRepo.FinderNewUserRepository(db)
	// rr := userRepo.NewRoleRepository(db, cache)
	// urr := userRepo.FinderNewUserRoleRepository(db, cache)
	// pr := userRepo.NewPermissionRepository(db, cache)
	createPinjamanRepo:= pinjamanRepo.CreateNewPinjamanRepository(db)
	finderPinjamanRepo:= pinjamanRepo.FinderNewPinjamanRepository(db)
	pinjamanRepository := pinjamanRepo.NewPinjamanRepository(db)

	// Cloud Storage
	cloudStorage := gcs.NewGoogleCloudStorage(cfg)
	// cloudStorage := aws.NewS3Bucket(cfg, awsSession)

	// Service
	userService := userService.NewUserFinder(cfg, ur)
	pinjamanFinder := service.NewPinjamanFinder(cfg, finderPinjamanRepo)
	pinjamanCreator := service.NewPinjamanCreator(cfg, createPinjamanRepo, cloudStorage)
	pinjamanUpdater := service.NewPinjamanUpdater(cfg, pinjamanRepository)

	// Handler
	app.PinjamanFinderHTTPHandler(cfg, router, pinjamanFinder)
	app.PinjamanCreatorHTTPHandler(cfg, router, pinjamanCreator, userService, cloudStorage)
	app.PinjamanUpdaterHTTPHandler(cfg, router, pinjamanUpdater, pinjamanFinder, cloudStorage)
}
