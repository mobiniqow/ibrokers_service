package measure_unit

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

const BucketName = "measure_unit"

var ErrMeasureUnitNotFound = errors.New("measureunit not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListMeasureUnits godoc
// @Summary      List of measureunits
// @Description  Get all measureunits
// @Tags         measureunit
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   MeasureUnitResponse
// @Router       /measure_unit/api/v1/ [get]
func (h *Handler) GetMeasureUnit(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    measureunits, count := h.Service.GetAllMeasureUnits(limit, page, filters.([]operators.FilterBlock))

    response := make([]MeasureUnitResponse, len(measureunits))
    for i, measureunit := range measureunits {
        response[i] = ToMeasureUnitResponse(measureunit)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetMeasureUnitDetails godoc
// @Summary      Get measureunit details
// @Description  Retrieve details of a measureunit by its ID
// @Tags         measureunit
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "MeasureUnit ID"
// @Success      200 {object} MeasureUnitResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "MeasureUnit not found"
// @Router       /measure_unit/api/v1/{id} [get]
func (h *Handler) GetMeasureUnitDetails(ctx *gin.Context) {
    measureunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    measureunit, err := h.Service.Repository.FindMeasureUnitById(measureunitId)
    if errors.Is(err, ErrMeasureUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "measureunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMeasureUnitResponse(measureunit)
    ctx.JSON(http.StatusOK, response)
}

// CreateMeasureUnit godoc
// @Summary      Create measureunit
// @Description  Create a new measureunit with the provided information
// @Tags         measureunit
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "MeasureUnit name"
// @Param        age   formData  int     true  "MeasureUnit age"
// @Param        image formData  file    true  "MeasureUnit image"
// @Success      201 {object} MeasureUnitResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /measure_unit/api/v1/ [post]
func (h *Handler) CreateMeasureUnit(ctx *gin.Context) {
    var req  MeasureUnit 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newMeasureUnit, err := h.Service.CreateMeasureUnit(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMeasureUnitResponse(newMeasureUnit)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateMeasureUnit godoc
// @Summary      Update measureunit
// @Description  Update measureunit details by ID
// @Tags         measureunit
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "MeasureUnit ID"
// @Param        name  formData string  false "MeasureUnit name"
// @Param        age   formData int     false "MeasureUnit age"
// @Param        image formData file    false "MeasureUnit image"
// @Success      200 {object} MeasureUnitResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "MeasureUnit not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /measure_unit/api/v1/{id} [put]
func (h *Handler) UpdateMeasureUnit(ctx *gin.Context) {
    measureunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    measureunit, err := h.Service.Repository.FindMeasureUnitById(measureunitId)
    if errors.Is(err, ErrMeasureUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "measureunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateMeasureUnitRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateMeasureUnit(&measureunit,&req)
    
    if err := h.Service.UpdateMeasureUnit(measureunit); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMeasureUnitResponse(measureunit)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         measureunit
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "MeasureUnit ID"
// @Param        city body    CreateMeasureUnitRequest true "Partial MeasureUnit information"
// @Success      200 {object} MeasureUnitResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "measureunit not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /measure_unit/api/v1/{id} [patch]
func (h *Handler) UpdateMeasureUnitPartial(ctx *gin.Context) {
    measureunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    measureunit, err := h.Service.Repository.FindMeasureUnitById(measureunitId)
    if errors.Is(err, ErrMeasureUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "measureunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateMeasureUnitRequest
    updateMeasureUnit(&measureunit,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateMeasureUnit(measureunit); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToMeasureUnitResponse(measureunit)
    ctx.JSON(http.StatusOK, response)
}


// DeleteMeasureUnit godoc
// @Summary      Delete measureunit
// @Description  Delete a measureunit by its ID
// @Tags         measureunit
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "MeasureUnit ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "MeasureUnit not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /measure_unit/api/v1/{id} [delete]
func (h *Handler) DeleteMeasureUnit(ctx *gin.Context) {
    measureunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    measureunit, err := h.Service.Repository.FindMeasureUnitById(measureunitId)
    if errors.Is(err, ErrMeasureUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "measureunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteMeasureUnit(measureunit); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateMeasureUnit(measureunit *MeasureUnit, req *CreateMeasureUnitRequest) error {
	measureunitVal := reflect.ValueOf(measureunit).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			measureunitField := measureunitVal.FieldByName(reqVal.Type().Field(i).Name)
			if measureunitField.IsValid() && measureunitField.CanSet() {
				measureunitField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
