package currency_unit

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    CurrencyUnitHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        CurrencyUnitHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.CurrencyUnitHandler.GetCurrencyUnit)
        groupV1.POST("/", e.CurrencyUnitHandler.CreateCurrencyUnit)
        groupV1.GET("/:id/", e.CurrencyUnitHandler.GetCurrencyUnitDetails)
        groupV1.PUT("/:id/", e.CurrencyUnitHandler.UpdateCurrencyUnit)
        groupV1.PATCH("/:id/", e.CurrencyUnitHandler.UpdateCurrencyUnitPartial)
        groupV1.DELETE("/:id/", e.CurrencyUnitHandler.DeleteCurrencyUnit)
    }
}
