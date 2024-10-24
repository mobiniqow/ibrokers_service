package contract_type

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {
    Router      *gin.RouterGroup
    ContractTypeHandler Handler
}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {
    return &Endpoints{
        Router:      router,
        ContractTypeHandler: Handler{Service: s, FileManager: fileManager},
    }
}

func (e *Endpoints) V1() {
    groupV1 := e.Router.Group("/api/v1")
    {
        groupV1.GET("/", e.ContractTypeHandler.GetContractType)
        groupV1.POST("/", e.ContractTypeHandler.CreateContractType)
        groupV1.GET("/:id/", e.ContractTypeHandler.GetContractTypeDetails)
        groupV1.PUT("/:id/", e.ContractTypeHandler.UpdateContractType)
        groupV1.PATCH("/:id/", e.ContractTypeHandler.UpdateContractTypePartial)
        groupV1.DELETE("/:id/", e.ContractTypeHandler.DeleteContractType)
    }
}
