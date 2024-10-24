package offer_mod

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    OfferModHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        OfferModHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.OfferModHandler.GetOfferMod)
        groupV1.POST("/", e.OfferModHandler.CreateOfferMod)
        groupV1.GET("/:id/", e.OfferModHandler.GetOfferModDetails)
        groupV1.PUT("/:id/", e.OfferModHandler.UpdateOfferMod)
        groupV1.PATCH("/:id/", e.OfferModHandler.UpdateOfferModPartial)
        groupV1.DELETE("/:id/", e.OfferModHandler.DeleteOfferMod)
    }
}
