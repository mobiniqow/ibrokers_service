package hall_menu_sub_group

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    HallMenuSubGroupHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        HallMenuSubGroupHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.HallMenuSubGroupHandler.GetHallMenuSubGroup)
        groupV1.POST("/", e.HallMenuSubGroupHandler.CreateHallMenuSubGroup)
        groupV1.GET("/:id/", e.HallMenuSubGroupHandler.GetHallMenuSubGroupDetails)
        groupV1.PUT("/:id/", e.HallMenuSubGroupHandler.UpdateHallMenuSubGroup)
        groupV1.PATCH("/:id/", e.HallMenuSubGroupHandler.UpdateHallMenuSubGroupPartial)
        groupV1.DELETE("/:id/", e.HallMenuSubGroupHandler.DeleteHallMenuSubGroup)
    }
}
