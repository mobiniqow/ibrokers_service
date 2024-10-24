package offer_type

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

const BucketName = "offer_type"

var ErrOfferTypeNotFound = errors.New("offertype not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListOfferTypes godoc
// @Summary      List of offertypes
// @Description  Get all offertypes
// @Tags         offertype
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   OfferTypeResponse
// @Router       /offer_type/api/v1/ [get]
func (h *Handler) GetOfferType(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    offertypes, count := h.Service.GetAllOfferTypes(limit, page, filters.([]operators.FilterBlock))

    response := make([]OfferTypeResponse, len(offertypes))
    for i, offertype := range offertypes {
        response[i] = ToOfferTypeResponse(offertype)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetOfferTypeDetails godoc
// @Summary      Get offertype details
// @Description  Retrieve details of a offertype by its ID
// @Tags         offertype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "OfferType ID"
// @Success      200 {object} OfferTypeResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "OfferType not found"
// @Router       /offer_type/api/v1/{id} [get]
func (h *Handler) GetOfferTypeDetails(ctx *gin.Context) {
    offertypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offertype, err := h.Service.Repository.FindOfferTypeById(offertypeId)
    if errors.Is(err, ErrOfferTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offertype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferTypeResponse(offertype)
    ctx.JSON(http.StatusOK, response)
}

// CreateOfferType godoc
// @Summary      Create offertype
// @Description  Create a new offertype with the provided information
// @Tags         offertype
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "OfferType name"
// @Param        age   formData  int     true  "OfferType age"
// @Param        image formData  file    true  "OfferType image"
// @Success      201 {object} OfferTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_type/api/v1/ [post]
func (h *Handler) CreateOfferType(ctx *gin.Context) {
    var req  OfferType 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newOfferType, err := h.Service.CreateOfferType(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferTypeResponse(newOfferType)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateOfferType godoc
// @Summary      Update offertype
// @Description  Update offertype details by ID
// @Tags         offertype
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "OfferType ID"
// @Param        name  formData string  false "OfferType name"
// @Param        age   formData int     false "OfferType age"
// @Param        image formData file    false "OfferType image"
// @Success      200 {object} OfferTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "OfferType not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_type/api/v1/{id} [put]
func (h *Handler) UpdateOfferType(ctx *gin.Context) {
    offertypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offertype, err := h.Service.Repository.FindOfferTypeById(offertypeId)
    if errors.Is(err, ErrOfferTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offertype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateOfferTypeRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateOfferType(&offertype,&req)
    
    if err := h.Service.UpdateOfferType(offertype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferTypeResponse(offertype)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         offertype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "OfferType ID"
// @Param        city body    CreateOfferTypeRequest true "Partial OfferType information"
// @Success      200 {object} OfferTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "offertype not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_type/api/v1/{id} [patch]
func (h *Handler) UpdateOfferTypePartial(ctx *gin.Context) {
    offertypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offertype, err := h.Service.Repository.FindOfferTypeById(offertypeId)
    if errors.Is(err, ErrOfferTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offertype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateOfferTypeRequest
    updateOfferType(&offertype,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateOfferType(offertype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferTypeResponse(offertype)
    ctx.JSON(http.StatusOK, response)
}


// DeleteOfferType godoc
// @Summary      Delete offertype
// @Description  Delete a offertype by its ID
// @Tags         offertype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "OfferType ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "OfferType not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_type/api/v1/{id} [delete]
func (h *Handler) DeleteOfferType(ctx *gin.Context) {
    offertypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offertype, err := h.Service.Repository.FindOfferTypeById(offertypeId)
    if errors.Is(err, ErrOfferTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offertype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteOfferType(offertype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateOfferType(offertype *OfferType, req *CreateOfferTypeRequest) error {
	offertypeVal := reflect.ValueOf(offertype).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			offertypeField := offertypeVal.FieldByName(reqVal.Type().Field(i).Name)
			if offertypeField.IsValid() && offertypeField.CanSet() {
				offertypeField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
