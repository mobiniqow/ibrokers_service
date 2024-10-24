package settlement

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

const BucketName = "settlement"

var ErrSettlementNotFound = errors.New("settlement not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListSettlements godoc
// @Summary      List of settlements
// @Description  Get all settlements
// @Tags         settlement
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   SettlementResponse
// @Router       /settlement/api/v1/ [get]
func (h *Handler) GetSettlement(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    settlements, count := h.Service.GetAllSettlements(limit, page, filters.([]operators.FilterBlock))

    response := make([]SettlementResponse, len(settlements))
    for i, settlement := range settlements {
        response[i] = ToSettlementResponse(settlement)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetSettlementDetails godoc
// @Summary      Get settlement details
// @Description  Retrieve details of a settlement by its ID
// @Tags         settlement
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Settlement ID"
// @Success      200 {object} SettlementResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Settlement not found"
// @Router       /settlement/api/v1/{id} [get]
func (h *Handler) GetSettlementDetails(ctx *gin.Context) {
    settlementId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    settlement, err := h.Service.Repository.FindSettlementById(settlementId)
    if errors.Is(err, ErrSettlementNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "settlement not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSettlementResponse(settlement)
    ctx.JSON(http.StatusOK, response)
}

// CreateSettlement godoc
// @Summary      Create settlement
// @Description  Create a new settlement with the provided information
// @Tags         settlement
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Settlement name"
// @Param        age   formData  int     true  "Settlement age"
// @Param        image formData  file    true  "Settlement image"
// @Success      201 {object} SettlementResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /settlement/api/v1/ [post]
func (h *Handler) CreateSettlement(ctx *gin.Context) {
    var req  Settlement 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newSettlement, err := h.Service.CreateSettlement(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSettlementResponse(newSettlement)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateSettlement godoc
// @Summary      Update settlement
// @Description  Update settlement details by ID
// @Tags         settlement
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Settlement ID"
// @Param        name  formData string  false "Settlement name"
// @Param        age   formData int     false "Settlement age"
// @Param        image formData file    false "Settlement image"
// @Success      200 {object} SettlementResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Settlement not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /settlement/api/v1/{id} [put]
func (h *Handler) UpdateSettlement(ctx *gin.Context) {
    settlementId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    settlement, err := h.Service.Repository.FindSettlementById(settlementId)
    if errors.Is(err, ErrSettlementNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "settlement not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateSettlementRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateSettlement(&settlement,&req)
    
    if err := h.Service.UpdateSettlement(settlement); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSettlementResponse(settlement)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         settlement
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Settlement ID"
// @Param        city body    CreateSettlementRequest true "Partial Settlement information"
// @Success      200 {object} SettlementResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "settlement not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /settlement/api/v1/{id} [patch]
func (h *Handler) UpdateSettlementPartial(ctx *gin.Context) {
    settlementId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    settlement, err := h.Service.Repository.FindSettlementById(settlementId)
    if errors.Is(err, ErrSettlementNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "settlement not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateSettlementRequest
    updateSettlement(&settlement,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateSettlement(settlement); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSettlementResponse(settlement)
    ctx.JSON(http.StatusOK, response)
}


// DeleteSettlement godoc
// @Summary      Delete settlement
// @Description  Delete a settlement by its ID
// @Tags         settlement
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Settlement ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Settlement not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /settlement/api/v1/{id} [delete]
func (h *Handler) DeleteSettlement(ctx *gin.Context) {
    settlementId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    settlement, err := h.Service.Repository.FindSettlementById(settlementId)
    if errors.Is(err, ErrSettlementNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "settlement not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteSettlement(settlement); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateSettlement(settlement *Settlement, req *CreateSettlementRequest) error {
	settlementVal := reflect.ValueOf(settlement).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			settlementField := settlementVal.FieldByName(reqVal.Type().Field(i).Name)
			if settlementField.IsValid() && settlementField.CanSet() {
				settlementField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
