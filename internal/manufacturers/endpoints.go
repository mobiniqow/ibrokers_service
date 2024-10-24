package manufacturers

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    ManufacturersHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        ManufacturersHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.ManufacturersHandler.GetManufacturers)
        groupV1.POST("/", e.ManufacturersHandler.CreateManufacturers)
        groupV1.GET("/:id/", e.ManufacturersHandler.GetManufacturersDetails)
        groupV1.PUT("/:id/", e.ManufacturersHandler.UpdateManufacturers)
        groupV1.PATCH("/:id/", e.ManufacturersHandler.UpdateManufacturersPartial)
        groupV1.DELETE("/:id/", e.ManufacturersHandler.DeleteManufacturers)
    }
}
