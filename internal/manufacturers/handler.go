package manufacturers

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

const BucketName = "manufacturers"

var ErrManufacturersNotFound = errors.New("manufacturers not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListManufacturerss godoc
// @Summary      List of manufacturerss
// @Description  Get all manufacturerss
// @Tags         manufacturers
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   ManufacturersResponse
// @Router       /manufacturers/api/v1/ [get]
func (h *Handler) GetManufacturers(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    manufacturerss, count := h.Service.GetAllManufacturerss(limit, page, filters.([]operators.FilterBlock))

    response := make([]ManufacturersResponse, len(manufacturerss))
    for i, manufacturers := range manufacturerss {
        response[i] = ToManufacturersResponse(manufacturers)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetManufacturersDetails godoc
// @Summary      Get manufacturers details
// @Description  Retrieve details of a manufacturers by its ID
// @Tags         manufacturers
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Manufacturers ID"
// @Success      200 {object} ManufacturersResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Manufacturers not found"
// @Router       /manufacturers/api/v1/{id} [get]
func (h *Handler) GetManufacturersDetails(ctx *gin.Context) {
    manufacturersId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    manufacturers, err := h.Service.Repository.FindManufacturersById(manufacturersId)
    if errors.Is(err, ErrManufacturersNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "manufacturers not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToManufacturersResponse(manufacturers)
    ctx.JSON(http.StatusOK, response)
}

// CreateManufacturers godoc
// @Summary      Create manufacturers
// @Description  Create a new manufacturers with the provided information
// @Tags         manufacturers
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Manufacturers name"
// @Param        age   formData  int     true  "Manufacturers age"
// @Param        image formData  file    true  "Manufacturers image"
// @Success      201 {object} ManufacturersResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /manufacturers/api/v1/ [post]
func (h *Handler) CreateManufacturers(ctx *gin.Context) {
    var req  Manufacturers 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newManufacturers, err := h.Service.CreateManufacturers(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToManufacturersResponse(newManufacturers)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateManufacturers godoc
// @Summary      Update manufacturers
// @Description  Update manufacturers details by ID
// @Tags         manufacturers
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Manufacturers ID"
// @Param        name  formData string  false "Manufacturers name"
// @Param        age   formData int     false "Manufacturers age"
// @Param        image formData file    false "Manufacturers image"
// @Success      200 {object} ManufacturersResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Manufacturers not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /manufacturers/api/v1/{id} [put]
func (h *Handler) UpdateManufacturers(ctx *gin.Context) {
    manufacturersId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    manufacturers, err := h.Service.Repository.FindManufacturersById(manufacturersId)
    if errors.Is(err, ErrManufacturersNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "manufacturers not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateManufacturersRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateManufacturers(&manufacturers,&req)
    
    if err := h.Service.UpdateManufacturers(manufacturers); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToManufacturersResponse(manufacturers)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         manufacturers
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Manufacturers ID"
// @Param        city body    CreateManufacturersRequest true "Partial Manufacturers information"
// @Success      200 {object} ManufacturersResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "manufacturers not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /manufacturers/api/v1/{id} [patch]
func (h *Handler) UpdateManufacturersPartial(ctx *gin.Context) {
    manufacturersId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    manufacturers, err := h.Service.Repository.FindManufacturersById(manufacturersId)
    if errors.Is(err, ErrManufacturersNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "manufacturers not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateManufacturersRequest
    updateManufacturers(&manufacturers,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateManufacturers(manufacturers); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToManufacturersResponse(manufacturers)
    ctx.JSON(http.StatusOK, response)
}


// DeleteManufacturers godoc
// @Summary      Delete manufacturers
// @Description  Delete a manufacturers by its ID
// @Tags         manufacturers
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Manufacturers ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Manufacturers not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /manufacturers/api/v1/{id} [delete]
func (h *Handler) DeleteManufacturers(ctx *gin.Context) {
    manufacturersId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    manufacturers, err := h.Service.Repository.FindManufacturersById(manufacturersId)
    if errors.Is(err, ErrManufacturersNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "manufacturers not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteManufacturers(manufacturers); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateManufacturers(manufacturers *Manufacturers, req *CreateManufacturersRequest) error {
	manufacturersVal := reflect.ValueOf(manufacturers).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			manufacturersField := manufacturersVal.FieldByName(reqVal.Type().Field(i).Name)
			if manufacturersField.IsValid() && manufacturersField.CanSet() {
				manufacturersField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
