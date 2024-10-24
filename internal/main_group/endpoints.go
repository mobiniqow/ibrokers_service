package main_group

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    MainGroupHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        MainGroupHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.MainGroupHandler.GetMainGroup)
        groupV1.POST("/", e.MainGroupHandler.CreateMainGroup)
        groupV1.GET("/:id/", e.MainGroupHandler.GetMainGroupDetails)
        groupV1.PUT("/:id/", e.MainGroupHandler.UpdateMainGroup)
        groupV1.PATCH("/:id/", e.MainGroupHandler.UpdateMainGroupPartial)
        groupV1.DELETE("/:id/", e.MainGroupHandler.DeleteMainGroup)
    }
}
