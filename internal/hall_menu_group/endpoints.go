package hall_menu_group

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    HallMenuGroupHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        HallMenuGroupHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.HallMenuGroupHandler.GetHallMenuGroup)
        groupV1.POST("/", e.HallMenuGroupHandler.CreateHallMenuGroup)
        groupV1.GET("/:id/", e.HallMenuGroupHandler.GetHallMenuGroupDetails)
        groupV1.PUT("/:id/", e.HallMenuGroupHandler.UpdateHallMenuGroup)
        groupV1.PATCH("/:id/", e.HallMenuGroupHandler.UpdateHallMenuGroupPartial)
        groupV1.DELETE("/:id/", e.HallMenuGroupHandler.DeleteHallMenuGroup)
    }
}
