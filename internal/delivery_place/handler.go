package delivery_place

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

const BucketName = "delivery_place"

var ErrDeliveryPlaceNotFound = errors.New("deliveryplace not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListDeliveryPlaces godoc
// @Summary      List of deliveryplaces
// @Description  Get all deliveryplaces
// @Tags         deliveryplace
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   DeliveryPlaceResponse
// @Router       /delivery_place/api/v1/ [get]
func (h *Handler) GetDeliveryPlace(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    deliveryplaces, count := h.Service.GetAllDeliveryPlaces(limit, page, filters.([]operators.FilterBlock))

    response := make([]DeliveryPlaceResponse, len(deliveryplaces))
    for i, deliveryplace := range deliveryplaces {
        response[i] = ToDeliveryPlaceResponse(deliveryplace)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetDeliveryPlaceDetails godoc
// @Summary      Get deliveryplace details
// @Description  Retrieve details of a deliveryplace by its ID
// @Tags         deliveryplace
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "DeliveryPlace ID"
// @Success      200 {object} DeliveryPlaceResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "DeliveryPlace not found"
// @Router       /delivery_place/api/v1/{id} [get]
func (h *Handler) GetDeliveryPlaceDetails(ctx *gin.Context) {
    deliveryplaceId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    deliveryplace, err := h.Service.Repository.FindDeliveryPlaceById(deliveryplaceId)
    if errors.Is(err, ErrDeliveryPlaceNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "deliveryplace not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToDeliveryPlaceResponse(deliveryplace)
    ctx.JSON(http.StatusOK, response)
}

// CreateDeliveryPlace godoc
// @Summary      Create deliveryplace
// @Description  Create a new deliveryplace with the provided information
// @Tags         deliveryplace
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "DeliveryPlace name"
// @Param        age   formData  int     true  "DeliveryPlace age"
// @Param        image formData  file    true  "DeliveryPlace image"
// @Success      201 {object} DeliveryPlaceResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /delivery_place/api/v1/ [post]
func (h *Handler) CreateDeliveryPlace(ctx *gin.Context) {
    var req  DeliveryPlace 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newDeliveryPlace, err := h.Service.CreateDeliveryPlace(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToDeliveryPlaceResponse(newDeliveryPlace)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateDeliveryPlace godoc
// @Summary      Update deliveryplace
// @Description  Update deliveryplace details by ID
// @Tags         deliveryplace
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "DeliveryPlace ID"
// @Param        name  formData string  false "DeliveryPlace name"
// @Param        age   formData int     false "DeliveryPlace age"
// @Param        image formData file    false "DeliveryPlace image"
// @Success      200 {object} DeliveryPlaceResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "DeliveryPlace not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /delivery_place/api/v1/{id} [put]
func (h *Handler) UpdateDeliveryPlace(ctx *gin.Context) {
    deliveryplaceId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    deliveryplace, err := h.Service.Repository.FindDeliveryPlaceById(deliveryplaceId)
    if errors.Is(err, ErrDeliveryPlaceNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "deliveryplace not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateDeliveryPlaceRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateDeliveryPlace(&deliveryplace,&req)
    
    if err := h.Service.UpdateDeliveryPlace(deliveryplace); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToDeliveryPlaceResponse(deliveryplace)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         deliveryplace
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "DeliveryPlace ID"
// @Param        city body    CreateDeliveryPlaceRequest true "Partial DeliveryPlace information"
// @Success      200 {object} DeliveryPlaceResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "deliveryplace not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /delivery_place/api/v1/{id} [patch]
func (h *Handler) UpdateDeliveryPlacePartial(ctx *gin.Context) {
    deliveryplaceId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    deliveryplace, err := h.Service.Repository.FindDeliveryPlaceById(deliveryplaceId)
    if errors.Is(err, ErrDeliveryPlaceNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "deliveryplace not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateDeliveryPlaceRequest
    updateDeliveryPlace(&deliveryplace,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateDeliveryPlace(deliveryplace); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToDeliveryPlaceResponse(deliveryplace)
    ctx.JSON(http.StatusOK, response)
}


// DeleteDeliveryPlace godoc
// @Summary      Delete deliveryplace
// @Description  Delete a deliveryplace by its ID
// @Tags         deliveryplace
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "DeliveryPlace ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "DeliveryPlace not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /delivery_place/api/v1/{id} [delete]
func (h *Handler) DeleteDeliveryPlace(ctx *gin.Context) {
    deliveryplaceId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    deliveryplace, err := h.Service.Repository.FindDeliveryPlaceById(deliveryplaceId)
    if errors.Is(err, ErrDeliveryPlaceNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "deliveryplace not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteDeliveryPlace(deliveryplace); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateDeliveryPlace(deliveryplace *DeliveryPlace, req *CreateDeliveryPlaceRequest) error {
	deliveryplaceVal := reflect.ValueOf(deliveryplace).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			deliveryplaceField := deliveryplaceVal.FieldByName(reqVal.Type().Field(i).Name)
			if deliveryplaceField.IsValid() && deliveryplaceField.CanSet() {
				deliveryplaceField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
