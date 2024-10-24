package sub_group

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    SubGroupHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        SubGroupHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.SubGroupHandler.GetSubGroup)
        groupV1.POST("/", e.SubGroupHandler.CreateSubGroup)
        groupV1.GET("/:id/", e.SubGroupHandler.GetSubGroupDetails)
        groupV1.PUT("/:id/", e.SubGroupHandler.UpdateSubGroup)
        groupV1.PATCH("/:id/", e.SubGroupHandler.UpdateSubGroupPartial)
        groupV1.DELETE("/:id/", e.SubGroupHandler.DeleteSubGroup)
    }
}
