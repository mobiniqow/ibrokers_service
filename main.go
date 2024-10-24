package main

import (
	"ibrokers_service/internal/broker"
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
	setupRoute(&router, db, fileManager)

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

func setupBrokerRoute(router *gin.RouterGroup, db *gorm.DB, fileManager *manager.FileManager) {
	rep := broker.Repository{DB: db}
	srv := broker.Service{Repository: rep}
	broker.CreateEndpoint(srv, router, *fileManager).V1()
}
