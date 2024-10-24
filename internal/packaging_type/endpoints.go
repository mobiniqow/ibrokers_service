package packaging_type

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    PackagingTypeHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        PackagingTypeHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.PackagingTypeHandler.GetPackagingType)
        groupV1.POST("/", e.PackagingTypeHandler.CreatePackagingType)
        groupV1.GET("/:id/", e.PackagingTypeHandler.GetPackagingTypeDetails)
        groupV1.PUT("/:id/", e.PackagingTypeHandler.UpdatePackagingType)
        groupV1.PATCH("/:id/", e.PackagingTypeHandler.UpdatePackagingTypePartial)
        groupV1.DELETE("/:id/", e.PackagingTypeHandler.DeletePackagingType)
    }
}
