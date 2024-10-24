package supplier

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    SupplierHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        SupplierHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.SupplierHandler.GetSupplier)
        groupV1.POST("/", e.SupplierHandler.CreateSupplier)
        groupV1.GET("/:id/", e.SupplierHandler.GetSupplierDetails)
        groupV1.PUT("/:id/", e.SupplierHandler.UpdateSupplier)
        groupV1.PATCH("/:id/", e.SupplierHandler.UpdateSupplierPartial)
        groupV1.DELETE("/:id/", e.SupplierHandler.DeleteSupplier)
    }
}
