package offer

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    OfferHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        OfferHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.OfferHandler.GetOffer)
        groupV1.POST("/", e.OfferHandler.CreateOffer)
        groupV1.GET("/:id/", e.OfferHandler.GetOfferDetails)
        groupV1.PUT("/:id/", e.OfferHandler.UpdateOffer)
        groupV1.PATCH("/:id/", e.OfferHandler.UpdateOfferPartial)
        groupV1.DELETE("/:id/", e.OfferHandler.DeleteOffer)
    }
}
