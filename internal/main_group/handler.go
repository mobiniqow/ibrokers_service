package main_group

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

const BucketName = "main_group"

var ErrMainGroupNotFound = errors.New("maingroup not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListMainGroups godoc
// @Summary      List of maingroups
// @Description  Get all maingroups
// @Tags         maingroup
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   MainGroupResponse
// @Router       /main_group/api/v1/ [get]
func (h *Handler) GetMainGroup(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    maingroups, count := h.Service.GetAllMainGroups(limit, page, filters.([]operators.FilterBlock))

    response := make([]MainGroupResponse, len(maingroups))
    for i, maingroup := range maingroups {
        response[i] = ToMainGroupResponse(maingroup)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetMainGroupDetails godoc
// @Summary      Get maingroup details
// @Description  Retrieve details of a maingroup by its ID
// @Tags         maingroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "MainGroup ID"
// @Success      200 {object} MainGroupResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "MainGroup not found"
// @Router       /main_group/api/v1/{id} [get]
func (h *Handler) GetMainGroupDetails(ctx *gin.Context) {
    maingroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    maingroup, err := h.Service.Repository.FindMainGroupById(maingroupId)
    if errors.Is(err, ErrMainGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "maingroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMainGroupResponse(maingroup)
    ctx.JSON(http.StatusOK, response)
}

// CreateMainGroup godoc
// @Summary      Create maingroup
// @Description  Create a new maingroup with the provided information
// @Tags         maingroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "MainGroup name"
// @Param        age   formData  int     true  "MainGroup age"
// @Param        image formData  file    true  "MainGroup image"
// @Success      201 {object} MainGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /main_group/api/v1/ [post]
func (h *Handler) CreateMainGroup(ctx *gin.Context) {
    var req  MainGroup 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newMainGroup, err := h.Service.CreateMainGroup(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMainGroupResponse(newMainGroup)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateMainGroup godoc
// @Summary      Update maingroup
// @Description  Update maingroup details by ID
// @Tags         maingroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "MainGroup ID"
// @Param        name  formData string  false "MainGroup name"
// @Param        age   formData int     false "MainGroup age"
// @Param        image formData file    false "MainGroup image"
// @Success      200 {object} MainGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "MainGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /main_group/api/v1/{id} [put]
func (h *Handler) UpdateMainGroup(ctx *gin.Context) {
    maingroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    maingroup, err := h.Service.Repository.FindMainGroupById(maingroupId)
    if errors.Is(err, ErrMainGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "maingroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateMainGroupRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateMainGroup(&maingroup,&req)
    
    if err := h.Service.UpdateMainGroup(maingroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMainGroupResponse(maingroup)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         maingroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "MainGroup ID"
// @Param        city body    CreateMainGroupRequest true "Partial MainGroup information"
// @Success      200 {object} MainGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "maingroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /main_group/api/v1/{id} [patch]
func (h *Handler) UpdateMainGroupPartial(ctx *gin.Context) {
    maingroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    maingroup, err := h.Service.Repository.FindMainGroupById(maingroupId)
    if errors.Is(err, ErrMainGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "maingroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateMainGroupRequest
    updateMainGroup(&maingroup,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateMainGroup(maingroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMainGroupResponse(maingroup)
    ctx.JSON(http.StatusOK, response)
}


// DeleteMainGroup godoc
// @Summary      Delete maingroup
// @Description  Delete a maingroup by its ID
// @Tags         maingroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "MainGroup ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "MainGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /main_group/api/v1/{id} [delete]
func (h *Handler) DeleteMainGroup(ctx *gin.Context) {
    maingroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    maingroup, err := h.Service.Repository.FindMainGroupById(maingroupId)
    if errors.Is(err, ErrMainGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "maingroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteMainGroup(maingroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateMainGroup(maingroup *MainGroup, req *CreateMainGroupRequest) error {
	maingroupVal := reflect.ValueOf(maingroup).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			maingroupField := maingroupVal.FieldByName(reqVal.Type().Field(i).Name)
			if maingroupField.IsValid() && maingroupField.CanSet() {
				maingroupField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
