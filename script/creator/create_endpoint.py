import os
import re
from utils.pascal_to_snake import pascal_to_snake
from utils.config import LOCATION

def create_endpoints_file(class_name):
    package_name = pascal_to_snake(class_name)
    file_package_name = LOCATION + package_name
    
    if not os.path.exists(file_package_name):
        os.makedirs(file_package_name)
    endpoints_filename = os.path.join(file_package_name, "endpoints.go")
    # Create the content of the endpoints file
    endpoints_content = f"""package {package_name}

import (
    "github.com/gin-gonic/gin"
    "ibrokers_service/pkg/utils/manager"
)

type Endpoints struct {{
    Router      *gin.RouterGroup
    {class_name}Handler Handler
}}

func CreateEndpoint(s Service, router *gin.RouterGroup, fileManager manager.FileManager) *Endpoints {{
    return &Endpoints{{
        Router:      router,
        {class_name}Handler: Handler{{Service: s, FileManager: fileManager}},
    }}
}}

func (e *Endpoints) V1() {{
    groupV1 := e.Router.Group("/api/v1")
    {{
        groupV1.GET("/", e.{class_name}Handler.Get{class_name})
        groupV1.POST("/", e.{class_name}Handler.Create{class_name})
        groupV1.GET("/:id/", e.{class_name}Handler.Get{class_name}Details)
        groupV1.PUT("/:id/", e.{class_name}Handler.Update{class_name})
        groupV1.PATCH("/:id/", e.{class_name}Handler.Update{class_name}Partial)
        groupV1.DELETE("/:id/", e.{class_name}Handler.Delete{class_name})
    }}
}}
"""
    # Save the Endpoints in the file
    with open(endpoints_filename, 'w') as file:
        file.write(endpoints_content)
    
    print(f"EndPoints done.")

 