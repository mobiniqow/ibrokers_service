package measure_unit

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    MeasureUnitHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        MeasureUnitHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.MeasureUnitHandler.GetMeasureUnit)
        groupV1.POST("/", e.MeasureUnitHandler.CreateMeasureUnit)
        groupV1.GET("/:id/", e.MeasureUnitHandler.GetMeasureUnitDetails)
        groupV1.PUT("/:id/", e.MeasureUnitHandler.UpdateMeasureUnit)
        groupV1.PATCH("/:id/", e.MeasureUnitHandler.UpdateMeasureUnitPartial)
        groupV1.DELETE("/:id/", e.MeasureUnitHandler.DeleteMeasureUnit)
    }
}
