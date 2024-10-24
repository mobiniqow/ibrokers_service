package buy_method

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

const BucketName = "buy_method"

var ErrBuyMethodNotFound = errors.New("buymethod not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListBuyMethods godoc
// @Summary      List of buymethods
// @Description  Get all buymethods
// @Tags         buymethod
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   BuyMethodResponse
// @Router       /buy_method/api/v1/ [get]
func (h *Handler) GetBuyMethod(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    buymethods, count := h.Service.GetAllBuyMethods(limit, page, filters.([]operators.FilterBlock))

    response := make([]BuyMethodResponse, len(buymethods))
    for i, buymethod := range buymethods {
        response[i] = ToBuyMethodResponse(buymethod)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetBuyMethodDetails godoc
// @Summary      Get buymethod details
// @Description  Retrieve details of a buymethod by its ID
// @Tags         buymethod
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "BuyMethod ID"
// @Success      200 {object} BuyMethodResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "BuyMethod not found"
// @Router       /buy_method/api/v1/{id} [get]
func (h *Handler) GetBuyMethodDetails(ctx *gin.Context) {
    buymethodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    buymethod, err := h.Service.Repository.FindBuyMethodById(buymethodId)
    if errors.Is(err, ErrBuyMethodNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "buymethod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBuyMethodResponse(buymethod)
    ctx.JSON(http.StatusOK, response)
}

// CreateBuyMethod godoc
// @Summary      Create buymethod
// @Description  Create a new buymethod with the provided information
// @Tags         buymethod
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "BuyMethod name"
// @Param        age   formData  int     true  "BuyMethod age"
// @Param        image formData  file    true  "BuyMethod image"
// @Success      201 {object} BuyMethodResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /buy_method/api/v1/ [post]
func (h *Handler) CreateBuyMethod(ctx *gin.Context) {
    var req  BuyMethod 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newBuyMethod, err := h.Service.CreateBuyMethod(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBuyMethodResponse(newBuyMethod)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateBuyMethod godoc
// @Summary      Update buymethod
// @Description  Update buymethod details by ID
// @Tags         buymethod
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "BuyMethod ID"
// @Param        name  formData string  false "BuyMethod name"
// @Param        age   formData int     false "BuyMethod age"
// @Param        image formData file    false "BuyMethod image"
// @Success      200 {object} BuyMethodResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "BuyMethod not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /buy_method/api/v1/{id} [put]
func (h *Handler) UpdateBuyMethod(ctx *gin.Context) {
    buymethodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    buymethod, err := h.Service.Repository.FindBuyMethodById(buymethodId)
    if errors.Is(err, ErrBuyMethodNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "buymethod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateBuyMethodRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateBuyMethod(&buymethod,&req)
    
    if err := h.Service.UpdateBuyMethod(buymethod); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBuyMethodResponse(buymethod)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         buymethod
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "BuyMethod ID"
// @Param        city body    CreateBuyMethodRequest true "Partial BuyMethod information"
// @Success      200 {object} BuyMethodResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "buymethod not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /buy_method/api/v1/{id} [patch]
func (h *Handler) UpdateBuyMethodPartial(ctx *gin.Context) {
    buymethodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    buymethod, err := h.Service.Repository.FindBuyMethodById(buymethodId)
    if errors.Is(err, ErrBuyMethodNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "buymethod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateBuyMethodRequest
    updateBuyMethod(&buymethod,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateBuyMethod(buymethod); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBuyMethodResponse(buymethod)
    ctx.JSON(http.StatusOK, response)
}


// DeleteBuyMethod godoc
// @Summary      Delete buymethod
// @Description  Delete a buymethod by its ID
// @Tags         buymethod
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "BuyMethod ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "BuyMethod not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /buy_method/api/v1/{id} [delete]
func (h *Handler) DeleteBuyMethod(ctx *gin.Context) {
    buymethodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    buymethod, err := h.Service.Repository.FindBuyMethodById(buymethodId)
    if errors.Is(err, ErrBuyMethodNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "buymethod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteBuyMethod(buymethod); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateBuyMethod(buymethod *BuyMethod, req *CreateBuyMethodRequest) error {
	buymethodVal := reflect.ValueOf(buymethod).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			buymethodField := buymethodVal.FieldByName(reqVal.Type().Field(i).Name)
			if buymethodField.IsValid() && buymethodField.CanSet() {
				buymethodField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
