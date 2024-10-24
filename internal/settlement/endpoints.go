package settlement

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    SettlementHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        SettlementHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.SettlementHandler.GetSettlement)
        groupV1.POST("/", e.SettlementHandler.CreateSettlement)
        groupV1.GET("/:id/", e.SettlementHandler.GetSettlementDetails)
        groupV1.PUT("/:id/", e.SettlementHandler.UpdateSettlement)
        groupV1.PATCH("/:id/", e.SettlementHandler.UpdateSettlementPartial)
        groupV1.DELETE("/:id/", e.SettlementHandler.DeleteSettlement)
    }
}
