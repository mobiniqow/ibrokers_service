package broker

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    BrokerHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        BrokerHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.BrokerHandler.GetBroker)
        groupV1.POST("/", e.BrokerHandler.CreateBroker)
        groupV1.GET("/:id/", e.BrokerHandler.GetBrokerDetails)
        groupV1.PUT("/:id/", e.BrokerHandler.UpdateBroker)
        groupV1.PATCH("/:id/", e.BrokerHandler.UpdateBrokerPartial)
        groupV1.DELETE("/:id/", e.BrokerHandler.DeleteBroker)
    }
}
