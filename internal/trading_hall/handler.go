package trading_hall

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

const BucketName = "trading_hall"

var ErrTradingHallNotFound = errors.New("tradinghall not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListTradingHalls godoc
// @Summary      List of tradinghalls
// @Description  Get all tradinghalls
// @Tags         tradinghall
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   TradingHallResponse
// @Router       /trading_hall/api/v1/ [get]
func (h *Handler) GetTradingHall(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    tradinghalls, count := h.Service.GetAllTradingHalls(limit, page, filters.([]operators.FilterBlock))

    response := make([]TradingHallResponse, len(tradinghalls))
    for i, tradinghall := range tradinghalls {
        response[i] = ToTradingHallResponse(tradinghall)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetTradingHallDetails godoc
// @Summary      Get tradinghall details
// @Description  Retrieve details of a tradinghall by its ID
// @Tags         tradinghall
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "TradingHall ID"
// @Success      200 {object} TradingHallResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "TradingHall not found"
// @Router       /trading_hall/api/v1/{id} [get]
func (h *Handler) GetTradingHallDetails(ctx *gin.Context) {
    tradinghallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    tradinghall, err := h.Service.Repository.FindTradingHallById(tradinghallId)
    if errors.Is(err, ErrTradingHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "tradinghall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToTradingHallResponse(tradinghall)
    ctx.JSON(http.StatusOK, response)
}

// CreateTradingHall godoc
// @Summary      Create tradinghall
// @Description  Create a new tradinghall with the provided information
// @Tags         tradinghall
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "TradingHall name"
// @Param        age   formData  int     true  "TradingHall age"
// @Param        image formData  file    true  "TradingHall image"
// @Success      201 {object} TradingHallResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /trading_hall/api/v1/ [post]
func (h *Handler) CreateTradingHall(ctx *gin.Context) {
    var req  TradingHall 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newTradingHall, err := h.Service.CreateTradingHall(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToTradingHallResponse(newTradingHall)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateTradingHall godoc
// @Summary      Update tradinghall
// @Description  Update tradinghall details by ID
// @Tags         tradinghall
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "TradingHall ID"
// @Param        name  formData string  false "TradingHall name"
// @Param        age   formData int     false "TradingHall age"
// @Param        image formData file    false "TradingHall image"
// @Success      200 {object} TradingHallResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "TradingHall not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /trading_hall/api/v1/{id} [put]
func (h *Handler) UpdateTradingHall(ctx *gin.Context) {
    tradinghallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    tradinghall, err := h.Service.Repository.FindTradingHallById(tradinghallId)
    if errors.Is(err, ErrTradingHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "tradinghall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateTradingHallRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateTradingHall(&tradinghall,&req)
    
    if err := h.Service.UpdateTradingHall(tradinghall); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToTradingHallResponse(tradinghall)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         tradinghall
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "TradingHall ID"
// @Param        city body    CreateTradingHallRequest true "Partial TradingHall information"
// @Success      200 {object} TradingHallResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "tradinghall not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /trading_hall/api/v1/{id} [patch]
func (h *Handler) UpdateTradingHallPartial(ctx *gin.Context) {
    tradinghallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    tradinghall, err := h.Service.Repository.FindTradingHallById(tradinghallId)
    if errors.Is(err, ErrTradingHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "tradinghall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateTradingHallRequest
    updateTradingHall(&tradinghall,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateTradingHall(tradinghall); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToTradingHallResponse(tradinghall)
    ctx.JSON(http.StatusOK, response)
}


// DeleteTradingHall godoc
// @Summary      Delete tradinghall
// @Description  Delete a tradinghall by its ID
// @Tags         tradinghall
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "TradingHall ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "TradingHall not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /trading_hall/api/v1/{id} [delete]
func (h *Handler) DeleteTradingHall(ctx *gin.Context) {
    tradinghallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    tradinghall, err := h.Service.Repository.FindTradingHallById(tradinghallId)
    if errors.Is(err, ErrTradingHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "tradinghall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteTradingHall(tradinghall); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateTradingHall(tradinghall *TradingHall, req *CreateTradingHallRequest) error {
	tradinghallVal := reflect.ValueOf(tradinghall).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			tradinghallField := tradinghallVal.FieldByName(reqVal.Type().Field(i).Name)
			if tradinghallField.IsValid() && tradinghallField.CanSet() {
				tradinghallField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
