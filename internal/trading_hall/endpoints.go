package trading_hall

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    TradingHallHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        TradingHallHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.TradingHallHandler.GetTradingHall)
        groupV1.POST("/", e.TradingHallHandler.CreateTradingHall)
        groupV1.GET("/:id/", e.TradingHallHandler.GetTradingHallDetails)
        groupV1.PUT("/:id/", e.TradingHallHandler.UpdateTradingHall)
        groupV1.PATCH("/:id/", e.TradingHallHandler.UpdateTradingHallPartial)
        groupV1.DELETE("/:id/", e.TradingHallHandler.DeleteTradingHall)
    }
}
