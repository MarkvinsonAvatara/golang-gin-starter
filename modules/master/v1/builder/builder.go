package builder

import (
	"gin-starter/app"
	"gin-starter/config"
	"gin-starter/modules/master/v1/repository"
	"gin-starter/modules/master/v1/service"
	"gin-starter/sdk/gcs"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// BuildMasterHandler builds master handler
// starting from handler down to repository or tool.
func BuildMasterHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
	// Repository
	provinceRepo := repository.NewProvinceRepository(db)
	dr := repository.NewDistrictRepository(db)
	vr := repository.NewVillageRepository(db)
	regencyRepo := repository.NewRegencyRepository(db)
	createBookRepository:=repository.CreateNewBookRepository(db)
	finderBookRepository:=repository.FinderNewBookRepository(db)
	deleteBookRepository:=repository.DeleterNewBookRepository(db)
	updaterBookRepository:=repository.UpdaterNewBookRepository(db)
	cloudStorage := gcs.NewGoogleCloudStorage(cfg)
	// cloudStorage := aws.NewS3Bucket(cfg, awsSession)

	// Service
	bookCreator := service.BookNewMasterCreator(cfg, createBookRepository, finderBookRepository,cloudStorage)
	// mc := service.(cfg, createBookRepository, finderBookRepository,cloudStorage)
	masterDeleter := service.NewMasterDeleter(cfg, deleteBookRepository, cloudStorage)
	bookFinder := service.BookNewMasterFinder(cfg,finderBookRepository)
	provinceFinder := service.ProvinceNewMasterFinder(cfg, provinceRepo)
	regencyFinder:=service.RegencyNewMasterFinder(cfg,regencyRepo)
	districtFinder:=service.DistrictNewMasterFinder(cfg,dr)
	villagerFinder:=service.VillageNewMasterFinder(cfg,vr)
	
	masterUpdater := service.NewMasterUpdater(cfg, updaterBookRepository, cloudStorage)

	// Handler
	app.BookMasterFinderHTTPHandler(cfg, router, bookFinder)
	app.BookMasterCreatorHTTPHandler(cfg, router, bookCreator, cloudStorage)
	app.BookMasterUpdaterHTTPHandler(cfg, router, masterUpdater, bookFinder, cloudStorage)
	app.BookMasterDeleterHTTPHandler(cfg, router, masterDeleter, cloudStorage)
	app.ProvinceFinderHTTPHandler(cfg, router, provinceFinder)
	app.RegencyFinderHTTPHandler(cfg, router, regencyFinder)
	app.DistrictFinderHTTPHandler(cfg, router, districtFinder)
	app.VillageFinderHTTPHandler(cfg, router, villagerFinder)
}
