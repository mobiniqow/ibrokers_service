import os
import re
from utils.pascal_to_snake import pascal_to_snake
from utils.config import LOCATION

def create_handler_file(class_name):
    package_name = pascal_to_snake(class_name)
    
    file_package_name = LOCATION+package_name
    
    if not os.path.exists(file_package_name):
        os.makedirs(file_package_name)
    handler_filename = os.path.join(LOCATION+package_name, "handler.go")

    # ایجاد محتویات فایل
    handler_content = f"""package {package_name}

import (
    "errors"
    "net/http"
    "ibrokers_service/pkg/middleware/filter/operators"
    "ibrokers_service/pkg/middleware/pagination"
    "ibrokers_service/pkg/utils/basics"
    "ibrokers_service/pkg/utils/manager"
    "strconv"
    "reflect"
    "github.com/gin-gonic/gin"
)

const BucketName = "{package_name}"

var Err{class_name}NotFound = errors.New("{class_name.lower()} not found")

type Handler struct {{
    Service     Service
    FileManager manager.FileManager
}}

// List{class_name}s godoc
// @Summary      List of {class_name.lower()}s
// @Description  Get all {class_name.lower()}s
// @Tags         {class_name.lower()}
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {{array}}   {class_name}Response
// @Router       /{package_name}/api/v1/ [get]
func (h *Handler) Get{class_name}(ctx *gin.Context) {{
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    {class_name.lower()}s, count := h.Service.GetAll{class_name}s(limit, page, filters.([]operators.FilterBlock))

    response := make([]{class_name}Response, len({class_name.lower()}s))
    for i, {class_name.lower()} := range {class_name.lower()}s {{
        response[i] = To{class_name}Response({class_name.lower()})
    }}
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}}

// Get{class_name}Details godoc
// @Summary      Get {class_name.lower()} details
// @Description  Retrieve details of a {class_name.lower()} by its ID
// @Tags         {class_name.lower()}
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "{class_name} ID"
// @Success      200 {{object}} {class_name}Response
// @Failure      400 {{object}} basics.APIError "Invalid UUID format"
// @Failure      404 {{object}} basics.APIError "{class_name} not found"
// @Router       /{package_name}/api/v1/{{id}} [get]
func (h *Handler) Get{class_name}Details(ctx *gin.Context) {{
    {class_name.lower()}Id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {{
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }}

    {class_name.lower()}, err := h.Service.Repository.Find{class_name}ById({class_name.lower()}Id)
    if errors.Is(err, Err{class_name}NotFound) {{
        basics.ErrorResponse(ctx, http.StatusNotFound, "{class_name.lower()} not found")
        return
    }} else if err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}

    response := To{class_name}Response({class_name.lower()})
    ctx.JSON(http.StatusOK, response)
}}

// Create{class_name} godoc
// @Summary      Create {class_name.lower()}
// @Description  Create a new {class_name.lower()} with the provided information
// @Tags         {class_name.lower()}
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "{class_name} name"
// @Param        age   formData  int     true  "{class_name} age"
// @Param        image formData  file    true  "{class_name} image"
// @Success      201 {{object}} {class_name}Response
// @Failure      400 {{object}} basics.APIError "Invalid request"
// @Failure      500 {{object}} basics.APIError "Internal server error"
// @Router       /{package_name}/api/v1/ [post]
func (h *Handler) Create{class_name}(ctx *gin.Context) {{
    var req  {class_name} 
    if err := ctx.ShouldBind(&req); err != nil {{
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }}  

    new{class_name}, err := h.Service.Create{class_name}(req)
    if err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}

    response := To{class_name}Response(new{class_name})
    ctx.JSON(http.StatusCreated, response)
}}

// Update{class_name} godoc
// @Summary      Update {class_name.lower()}
// @Description  Update {class_name.lower()} details by ID
// @Tags         {class_name.lower()}
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "{class_name} ID"
// @Param        name  formData string  false "{class_name} name"
// @Param        age   formData int     false "{class_name} age"
// @Param        image formData file    false "{class_name} image"
// @Success      200 {{object}} {class_name}Response
// @Failure      400 {{object}} basics.APIError "Invalid request"
// @Failure      404 {{object}} basics.APIError "{class_name} not found"
// @Failure      500 {{object}} basics.APIError "Internal server error"
// @Router       /{package_name}/api/v1/{{id}} [put]
func (h *Handler) Update{class_name}(ctx *gin.Context) {{
    {class_name.lower()}Id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {{
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }}

    {class_name.lower()}, err := h.Service.Repository.Find{class_name}ById({class_name.lower()}Id)
    if errors.Is(err, Err{class_name}NotFound) {{
        basics.ErrorResponse(ctx, http.StatusNotFound, "{class_name.lower()} not found")
        return
    }} else if err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}
    var req {class_name}
    if err := ctx.ShouldBind(&req); err != nil {{
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }}

   
    req.Id = {class_name.lower()}.Id
    if err := h.Service.Update{class_name}(req); err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}

    response := To{class_name}Response(req)
    ctx.JSON(http.StatusOK, response)
}} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         {class_name.lower()}
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "{class_name} ID"
// @Param        city body    Create{class_name}Request true "Partial {class_name} information"
// @Success      200 {{object}} {class_name}Response
// @Failure      400 {{object}} basics.APIError "Invalid request"
// @Failure      404 {{object}} basics.APIError "{class_name.lower()} not found"
// @Failure      500 {{object}} basics.APIError "Internal server error"
// @Router       /{package_name}/api/v1/{{id}} [patch]
func (h *Handler) Update{class_name}Partial(ctx *gin.Context) {{
    {class_name.lower()}Id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {{
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }}

    {class_name.lower()}, err := h.Service.Repository.Find{class_name}ById({class_name.lower()}Id)
    if errors.Is(err, Err{class_name}NotFound) {{
        basics.ErrorResponse(ctx, http.StatusNotFound, "{class_name.lower()} not found")
        return
    }} else if err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}

    var req Create{class_name}Request
    if err := ctx.ShouldBind(&req); err != nil {{
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }}
    update{class_name}(&{class_name.lower()},&req)
   
    if err := h.Service.Update{class_name}({class_name.lower()}); err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}

    response := To{class_name}Response({class_name.lower()})
    ctx.JSON(http.StatusOK, response)
}}


// Delete{class_name} godoc
// @Summary      Delete {class_name.lower()}
// @Description  Delete a {class_name.lower()} by its ID
// @Tags         {class_name.lower()}
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "{class_name} ID"
// @Success      204 "No Content"
// @Failure      400 {{object}} basics.APIError "Invalid UUID format"
// @Failure      404 {{object}} basics.APIError "{class_name} not found"
// @Failure      500 {{object}} basics.APIError "Internal server error"
// @Router       /{package_name}/api/v1/{{id}} [delete]
func (h *Handler) Delete{class_name}(ctx *gin.Context) {{
    {class_name.lower()}Id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {{
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }}

    {class_name.lower()}, err := h.Service.Repository.Find{class_name}ById({class_name.lower()}Id)
    if errors.Is(err, Err{class_name}NotFound) {{
        basics.ErrorResponse(ctx, http.StatusNotFound, "{class_name.lower()} not found")
        return
    }} else if err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}

    if err := h.Service.Repository.Delete{class_name}({class_name.lower()}); err != nil {{
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }}

    ctx.Status(http.StatusNoContent) // 204 No Content
}}

func update{class_name}({class_name.lower()} *{class_name}, req *Create{class_name}Request) error {{
	{class_name.lower()}Val := reflect.ValueOf({class_name.lower()}).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {{
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {{
			{class_name.lower()}Field := {class_name.lower()}Val.FieldByName(reqVal.Type().Field(i).Name)
			if {class_name.lower()}Field.IsValid() && {class_name.lower()}Field.CanSet() {{
				{class_name.lower()}Field.Set(reflect.Indirect(fieldVal))
			}}
		}}
	}}

	return nil
}}
"""


    # ذخیره Handler در فایل
    with open(handler_filename, 'w') as file:
        file.write(handler_content)
    
    print(f"Handler done.")

 