package commodity

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    CommodityHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        CommodityHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.CommodityHandler.GetCommodity)
        groupV1.POST("/", e.CommodityHandler.CreateCommodity)
        groupV1.GET("/:id/", e.CommodityHandler.GetCommodityDetails)
        groupV1.PUT("/:id/", e.CommodityHandler.UpdateCommodity)
        groupV1.PATCH("/:id/", e.CommodityHandler.UpdateCommodityPartial)
        groupV1.DELETE("/:id/", e.CommodityHandler.DeleteCommodity)
    }
}
