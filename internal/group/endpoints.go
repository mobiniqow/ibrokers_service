package group

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    GroupHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        GroupHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.GroupHandler.GetGroup)
        groupV1.POST("/", e.GroupHandler.CreateGroup)
        groupV1.GET("/:id/", e.GroupHandler.GetGroupDetails)
        groupV1.PUT("/:id/", e.GroupHandler.UpdateGroup)
        groupV1.PATCH("/:id/", e.GroupHandler.UpdateGroupPartial)
        groupV1.DELETE("/:id/", e.GroupHandler.DeleteGroup)
    }
}
