package report

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    ReportHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        ReportHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.ReportHandler.GetReport)
        groupV1.POST("/", e.ReportHandler.CreateReport)
        groupV1.GET("/:id/", e.ReportHandler.GetReportDetails)
        groupV1.PUT("/:id/", e.ReportHandler.UpdateReport)
        groupV1.PATCH("/:id/", e.ReportHandler.UpdateReportPartial)
        groupV1.DELETE("/:id/", e.ReportHandler.DeleteReport)
    }
}
