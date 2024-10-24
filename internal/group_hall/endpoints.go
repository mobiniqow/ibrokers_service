package group_hall

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    GroupHallHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        GroupHallHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.GroupHallHandler.GetGroupHall)
        groupV1.POST("/", e.GroupHallHandler.CreateGroupHall)
        groupV1.GET("/:id/", e.GroupHallHandler.GetGroupHallDetails)
        groupV1.PUT("/:id/", e.GroupHallHandler.UpdateGroupHall)
        groupV1.PATCH("/:id/", e.GroupHallHandler.UpdateGroupHallPartial)
        groupV1.DELETE("/:id/", e.GroupHallHandler.DeleteGroupHall)
    }
}
