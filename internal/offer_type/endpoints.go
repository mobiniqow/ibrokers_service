package offer_type

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    OfferTypeHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        OfferTypeHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.OfferTypeHandler.GetOfferType)
        groupV1.POST("/", e.OfferTypeHandler.CreateOfferType)
        groupV1.GET("/:id/", e.OfferTypeHandler.GetOfferTypeDetails)
        groupV1.PUT("/:id/", e.OfferTypeHandler.UpdateOfferType)
        groupV1.PATCH("/:id/", e.OfferTypeHandler.UpdateOfferTypePartial)
        groupV1.DELETE("/:id/", e.OfferTypeHandler.DeleteOfferType)
    }
}
