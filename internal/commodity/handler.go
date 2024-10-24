package commodity

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

const BucketName = "commodity"

var ErrCommodityNotFound = errors.New("commodity not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListCommoditys godoc
// @Summary      List of commoditys
// @Description  Get all commoditys
// @Tags         commodity
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   CommodityResponse
// @Router       /commodity/api/v1/ [get]
func (h *Handler) GetCommodity(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    commoditys, count := h.Service.GetAllCommoditys(limit, page, filters.([]operators.FilterBlock))

    response := make([]CommodityResponse, len(commoditys))
    for i, commodity := range commoditys {
        response[i] = ToCommodityResponse(commodity)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetCommodityDetails godoc
// @Summary      Get commodity details
// @Description  Retrieve details of a commodity by its ID
// @Tags         commodity
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Commodity ID"
// @Success      200 {object} CommodityResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Commodity not found"
// @Router       /commodity/api/v1/{id} [get]
func (h *Handler) GetCommodityDetails(ctx *gin.Context) {
    commodityId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    commodity, err := h.Service.Repository.FindCommodityById(commodityId)
    if errors.Is(err, ErrCommodityNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "commodity not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCommodityResponse(commodity)
    ctx.JSON(http.StatusOK, response)
}

// CreateCommodity godoc
// @Summary      Create commodity
// @Description  Create a new commodity with the provided information
// @Tags         commodity
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Commodity name"
// @Param        age   formData  int     true  "Commodity age"
// @Param        image formData  file    true  "Commodity image"
// @Success      201 {object} CommodityResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /commodity/api/v1/ [post]
func (h *Handler) CreateCommodity(ctx *gin.Context) {
    var req  Commodity 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newCommodity, err := h.Service.CreateCommodity(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCommodityResponse(newCommodity)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateCommodity godoc
// @Summary      Update commodity
// @Description  Update commodity details by ID
// @Tags         commodity
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Commodity ID"
// @Param        name  formData string  false "Commodity name"
// @Param        age   formData int     false "Commodity age"
// @Param        image formData file    false "Commodity image"
// @Success      200 {object} CommodityResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Commodity not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /commodity/api/v1/{id} [put]
func (h *Handler) UpdateCommodity(ctx *gin.Context) {
    commodityId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    commodity, err := h.Service.Repository.FindCommodityById(commodityId)
    if errors.Is(err, ErrCommodityNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "commodity not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateCommodityRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateCommodity(&commodity,&req)
    
    if err := h.Service.UpdateCommodity(commodity); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCommodityResponse(commodity)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         commodity
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Commodity ID"
// @Param        city body    CreateCommodityRequest true "Partial Commodity information"
// @Success      200 {object} CommodityResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "commodity not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /commodity/api/v1/{id} [patch]
func (h *Handler) UpdateCommodityPartial(ctx *gin.Context) {
    commodityId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    commodity, err := h.Service.Repository.FindCommodityById(commodityId)
    if errors.Is(err, ErrCommodityNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "commodity not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateCommodityRequest
    updateCommodity(&commodity,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateCommodity(commodity); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToCommodityResponse(commodity)
    ctx.JSON(http.StatusOK, response)
}


// DeleteCommodity godoc
// @Summary      Delete commodity
// @Description  Delete a commodity by its ID
// @Tags         commodity
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Commodity ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Commodity not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /commodity/api/v1/{id} [delete]
func (h *Handler) DeleteCommodity(ctx *gin.Context) {
    commodityId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    commodity, err := h.Service.Repository.FindCommodityById(commodityId)
    if errors.Is(err, ErrCommodityNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "commodity not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteCommodity(commodity); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateCommodity(commodity *Commodity, req *CreateCommodityRequest) error {
	commodityVal := reflect.ValueOf(commodity).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			commodityField := commodityVal.FieldByName(reqVal.Type().Field(i).Name)
			if commodityField.IsValid() && commodityField.CanSet() {
				commodityField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
