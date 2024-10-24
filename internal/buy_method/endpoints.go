package buy_method

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    BuyMethodHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        BuyMethodHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.BuyMethodHandler.GetBuyMethod)
        groupV1.POST("/", e.BuyMethodHandler.CreateBuyMethod)
        groupV1.GET("/:id/", e.BuyMethodHandler.GetBuyMethodDetails)
        groupV1.PUT("/:id/", e.BuyMethodHandler.UpdateBuyMethod)
        groupV1.PATCH("/:id/", e.BuyMethodHandler.UpdateBuyMethodPartial)
        groupV1.DELETE("/:id/", e.BuyMethodHandler.DeleteBuyMethod)
    }
}
