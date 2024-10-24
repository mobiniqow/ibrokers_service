package delivery_place

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    DeliveryPlaceHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        DeliveryPlaceHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.DeliveryPlaceHandler.GetDeliveryPlace)
        groupV1.POST("/", e.DeliveryPlaceHandler.CreateDeliveryPlace)
        groupV1.GET("/:id/", e.DeliveryPlaceHandler.GetDeliveryPlaceDetails)
        groupV1.PUT("/:id/", e.DeliveryPlaceHandler.UpdateDeliveryPlace)
        groupV1.PATCH("/:id/", e.DeliveryPlaceHandler.UpdateDeliveryPlacePartial)
        groupV1.DELETE("/:id/", e.DeliveryPlaceHandler.DeleteDeliveryPlace)
    }
}
