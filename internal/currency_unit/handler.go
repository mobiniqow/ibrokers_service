package currency_unit

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

const BucketName = "currency_unit"

var ErrCurrencyUnitNotFound = errors.New("currencyunit not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListCurrencyUnits godoc
// @Summary      List of currencyunits
// @Description  Get all currencyunits
// @Tags         currencyunit
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   CurrencyUnitResponse
// @Router       /currency_unit/api/v1/ [get]
func (h *Handler) GetCurrencyUnit(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    currencyunits, count := h.Service.GetAllCurrencyUnits(limit, page, filters.([]operators.FilterBlock))

    response := make([]CurrencyUnitResponse, len(currencyunits))
    for i, currencyunit := range currencyunits {
        response[i] = ToCurrencyUnitResponse(currencyunit)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetCurrencyUnitDetails godoc
// @Summary      Get currencyunit details
// @Description  Retrieve details of a currencyunit by its ID
// @Tags         currencyunit
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "CurrencyUnit ID"
// @Success      200 {object} CurrencyUnitResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "CurrencyUnit not found"
// @Router       /currency_unit/api/v1/{id} [get]
func (h *Handler) GetCurrencyUnitDetails(ctx *gin.Context) {
    currencyunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    currencyunit, err := h.Service.Repository.FindCurrencyUnitById(currencyunitId)
    if errors.Is(err, ErrCurrencyUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "currencyunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCurrencyUnitResponse(currencyunit)
    ctx.JSON(http.StatusOK, response)
}

// CreateCurrencyUnit godoc
// @Summary      Create currencyunit
// @Description  Create a new currencyunit with the provided information
// @Tags         currencyunit
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "CurrencyUnit name"
// @Param        age   formData  int     true  "CurrencyUnit age"
// @Param        image formData  file    true  "CurrencyUnit image"
// @Success      201 {object} CurrencyUnitResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /currency_unit/api/v1/ [post]
func (h *Handler) CreateCurrencyUnit(ctx *gin.Context) {
    var req  CurrencyUnit 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newCurrencyUnit, err := h.Service.CreateCurrencyUnit(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCurrencyUnitResponse(newCurrencyUnit)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateCurrencyUnit godoc
// @Summary      Update currencyunit
// @Description  Update currencyunit details by ID
// @Tags         currencyunit
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "CurrencyUnit ID"
// @Param        name  formData string  false "CurrencyUnit name"
// @Param        age   formData int     false "CurrencyUnit age"
// @Param        image formData file    false "CurrencyUnit image"
// @Success      200 {object} CurrencyUnitResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "CurrencyUnit not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /currency_unit/api/v1/{id} [put]
func (h *Handler) UpdateCurrencyUnit(ctx *gin.Context) {
    currencyunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    currencyunit, err := h.Service.Repository.FindCurrencyUnitById(currencyunitId)
    if errors.Is(err, ErrCurrencyUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "currencyunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateCurrencyUnitRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateCurrencyUnit(&currencyunit,&req)
    
    if err := h.Service.UpdateCurrencyUnit(currencyunit); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCurrencyUnitResponse(currencyunit)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         currencyunit
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "CurrencyUnit ID"
// @Param        city body    CreateCurrencyUnitRequest true "Partial CurrencyUnit information"
// @Success      200 {object} CurrencyUnitResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "currencyunit not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /currency_unit/api/v1/{id} [patch]
func (h *Handler) UpdateCurrencyUnitPartial(ctx *gin.Context) {
    currencyunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    currencyunit, err := h.Service.Repository.FindCurrencyUnitById(currencyunitId)
    if errors.Is(err, ErrCurrencyUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "currencyunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateCurrencyUnitRequest
    updateCurrencyUnit(&currencyunit,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateCurrencyUnit(currencyunit); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCurrencyUnitResponse(currencyunit)
    ctx.JSON(http.StatusOK, response)
}


// DeleteCurrencyUnit godoc
// @Summary      Delete currencyunit
// @Description  Delete a currencyunit by its ID
// @Tags         currencyunit
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "CurrencyUnit ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "CurrencyUnit not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /currency_unit/api/v1/{id} [delete]
func (h *Handler) DeleteCurrencyUnit(ctx *gin.Context) {
    currencyunitId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    currencyunit, err := h.Service.Repository.FindCurrencyUnitById(currencyunitId)
    if errors.Is(err, ErrCurrencyUnitNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "currencyunit not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteCurrencyUnit(currencyunit); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateCurrencyUnit(currencyunit *CurrencyUnit, req *CreateCurrencyUnitRequest) error {
	currencyunitVal := reflect.ValueOf(currencyunit).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			currencyunitField := currencyunitVal.FieldByName(reqVal.Type().Field(i).Name)
			if currencyunitField.IsValid() && currencyunitField.CanSet() {
				currencyunitField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
