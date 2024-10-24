package main

import (
	"ibrokers_service/internal/broker"
	"ibrokers_service/internal/buy_method"

	// "ibrokers_service/internal/commodity"
	// "ibrokers_service/internal/contract_type"
	// "ibrokers_service/internal/currency_unit"
	// "ibrokers_service/internal/delivery_place"
	// "ibrokers_service/internal/group"
	// "ibrokers_service/internal/group_hall"
	// "ibrokers_service/internal/hall_menu_group"
	// "ibrokers_service/internal/hall_menu_sub_group"
	// "ibrokers_service/internal/main_group"
	// "ibrokers_service/internal/manufacturers"
	// "ibrokers_service/internal/measure_unit"

	// "ibrokers_service/internal/offer_mod"
	// "ibrokers_service/internal/offer_type"
	// "ibrokers_service/internal/packaging_type"
	// "ibrokers_service/internal/report"
	// "ibrokers_service/internal/settlement"
	// "ibrokers_service/internal/sub_group"
	// "ibrokers_service/internal/supplier"
	// "ibrokers_service/internal/trading_hall"
	"ibrokers_service/pkg/configs"
	"ibrokers_service/pkg/middleware/error_handler"
	"ibrokers_service/pkg/middleware/filter"
	"ibrokers_service/pkg/middleware/filter/operators"
	"ibrokers_service/pkg/middleware/logger"
	"ibrokers_service/pkg/middleware/pagination"
	"ibrokers_service/pkg/utils/manager"
	"log"
	"net/url"
	"time"

	docs "ibrokers_service/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki-client-go/pkg/urlutil"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error

	// Loki client
	lokiClient := setupLokiClient()

	// Minio
	fileManager := setupMinio()

	// GORM
	db := setupDatabase()

	// Migrations
	if err = db.AutoMigrate(&broker.Broker{}); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	app := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	// Middleware
	setupMiddleware(app, lokiClient)

	// Routing
	router := app.RouterGroup
	setupRoutes(&router, db, fileManager)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = app.Run(":5500")
	if err != nil {
		log.Fatalln(err)
	}
}

func setupLokiClient() *loki.Client {
	lokiURL, err := url.Parse("http://localhost:3100/loki/api/v1/push")
	if err != nil {
		log.Fatalf("error parsing Loki URL: %v", err)
	}

	lokiClient, err := loki.New(loki.Config{
		URL:     urlutil.URLValue{URL: lokiURL},
		Timeout: 10 * time.Second,
	})

	if err != nil {
		log.Fatalf("error creating Loki client: %v", err)
	}
	return lokiClient
}

func setupMinio() *manager.FileManager {
	_minio := configs.NewMinio(false, configs.MINIO_SECRET_KEY, configs.MINIO_ACCESS_KEY, configs.MINIO_URL)
	return manager.NewFileManager(configs.BASE_URL, *_minio)
}

func setupDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db not connected: %v", err)
	}
	return db
}

func setupMiddleware(app *gin.Engine, lokiClient *loki.Client) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     configs.CORS_URL,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app.Use(pagination.Middleware())
	gormOperator := operators.GormOperator{}
	filterMapper := filter.Mapper{Operators: gormOperator}
	app.Use(filter.QueryFilterMiddleware(filterMapper))
	app.Use(logger.Logger(lokiClient))
	app.Use(error_handler.ErrorHandlingMiddleware())
}

func setupRoutes(router *gin.RouterGroup, db *gorm.DB, fileManager *manager.FileManager) {
	{
		rep := broker.Repository{DB: db}
		srv := broker.Service{Repository: rep}
		broker.CreateEndpoint(srv, router.Group("/broker"), *fileManager).V1()
	}
	{
		rep := buy_method.Repository{DB: db}
		srv := buy_method.Service{Repository: rep}
		buy_method.CreateEndpoint(srv, router.Group("/buy-method"), *fileManager).V1()
	}
	// {
	// 	rep := commodity.Repository{DB: db}
	// 	srv := commodity.Service{Repository: rep}
	// 	commodity.CreateEndpoint(srv, router.Group("/commodity"), *fileManager).V1()
	// }
	// {
	// 	rep := contract_type.Repository{DB: db}
	// 	srv := contract_type.Service{Repository: rep}
	// 	contract_type.CreateEndpoint(srv, router.Group("/contract-type"), *fileManager).V1()
	// }
	// {
	// 	rep := currency_unit.Repository{DB: db}
	// 	srv := currency_unit.Service{Repository: rep}
	// 	currency_unit.CreateEndpoint(srv, router.Group("/contract-unit"), *fileManager).V1()
	// }
	// {
	// 	rep := delivery_place.Repository{DB: db}
	// 	srv := delivery_place.Service{Repository: rep}
	// 	delivery_place.CreateEndpoint(srv, router.Group("/delivery-place"), *fileManager).V1()
	// }
	// {
	// 	rep := main_group.Repository{DB: db}
	// 	srv := main_group.Service{Repository: rep}
	// 	main_group.CreateEndpoint(srv, router.Group("/menu-group"), *fileManager).V1()
	// }
	// {
	// 	rep := group.Repository{DB: db}
	// 	srv := group.Service{Repository: rep}
	// 	group.CreateEndpoint(srv, router.Group("/group"), *fileManager).V1()
	// }
	// {
	// 	rep := sub_group.Repository{DB: db}
	// 	srv := sub_group.Service{Repository: rep}
	// 	sub_group.CreateEndpoint(srv, router.Group("/sub-group"), *fileManager).V1()
	// }
	// {
	// 	rep := group_hall.Repository{DB: db}
	// 	srv := group_hall.Service{Repository: rep}
	// 	group_hall.CreateEndpoint(srv, router.Group("/group-hall"), *fileManager).V1()
	// }
	// {
	// 	rep := hall_menu_group.Repository{DB: db}
	// 	srv := hall_menu_group.Service{Repository: rep}
	// 	hall_menu_group.CreateEndpoint(srv, router.Group("/hall-menu-group"), *fileManager).V1()
	// }
	// {
	// 	rep := trading_hall.Repository{DB: db}
	// 	srv := trading_hall.Service{Repository: rep}
	// 	trading_hall.CreateEndpoint(srv, router.Group("/trading-hall"), *fileManager).V1()
	// }
	// {
	// 	rep := hall_menu_sub_group.Repository{DB: db}
	// 	srv := hall_menu_sub_group.Service{Repository: rep}
	// 	hall_menu_sub_group.CreateEndpoint(srv, router.Group("/hall-menu-sub-group"), *fileManager).V1()
	// }
	// {
	// 	rep := manufacturers.Repository{DB: db}
	// 	srv := manufacturers.Service{Repository: rep}
	// 	manufacturers.CreateEndpoint(srv, router.Group("/manufacturers"), *fileManager).V1()
	// }
	// {
	// 	rep := measure_unit.Repository{DB: db}
	// 	srv := measure_unit.Service{Repository: rep}
	// 	measure_unit.CreateEndpoint(srv, router.Group("/measure-unit"), *fileManager).V1()
	// }
	// {
	// 	rep := offer.Repository{DB: db}
	// 	srv := offer.Service{Repository: rep}
	// 	offer.CreateEndpoint(srv, router.Group("/offer"), *fileManager).V1()
	// }
	// {
	// 	rep := offer_mod.Repository{DB: db}
	// 	srv := offer_mod.Service{Repository: rep}
	// 	offer_mod.CreateEndpoint(srv, router.Group("/offer-mod"), *fileManager).V1()
	// }
	// {
	// 	rep := offer_type.Repository{DB: db}
	// 	srv := offer_type.Service{Repository: rep}
	// 	offer_type.CreateEndpoint(srv, router.Group("/offer-type"), *fileManager).V1()
	// }
	// {
	// 	rep := packaging_type.Repository{DB: db}
	// 	srv := packaging_type.Service{Repository: rep}
	// 	packaging_type.CreateEndpoint(srv, router.Group("/packaging-type"), *fileManager).V1()
	// }
	// {
	// 	rep := report.Repository{DB: db}
	// 	srv := report.Service{Repository: rep}
	// 	report.CreateEndpoint(srv, router.Group("/report"), *fileManager).V1()
	// }
	// {
	// 	rep := settlement.Repository{DB: db}
	// 	srv := settlement.Service{Repository: rep}
	// 	settlement.CreateEndpoint(srv, router.Group("/settlement"), *fileManager).V1()
	// }
	// {
	// 	rep := supplier.Repository{DB: db}
	// 	srv := supplier.Service{Repository: rep}
	// 	supplier.CreateEndpoint(srv, router.Group("/supplier"), *fileManager).V1()
	// }
}
