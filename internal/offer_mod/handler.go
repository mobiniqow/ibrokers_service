package offer_mod

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

const BucketName = "offer_mod"

var ErrOfferModNotFound = errors.New("offermod not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListOfferMods godoc
// @Summary      List of offermods
// @Description  Get all offermods
// @Tags         offermod
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   OfferModResponse
// @Router       /offer_mod/api/v1/ [get]
func (h *Handler) GetOfferMod(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    offermods, count := h.Service.GetAllOfferMods(limit, page, filters.([]operators.FilterBlock))

    response := make([]OfferModResponse, len(offermods))
    for i, offermod := range offermods {
        response[i] = ToOfferModResponse(offermod)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetOfferModDetails godoc
// @Summary      Get offermod details
// @Description  Retrieve details of a offermod by its ID
// @Tags         offermod
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "OfferMod ID"
// @Success      200 {object} OfferModResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "OfferMod not found"
// @Router       /offer_mod/api/v1/{id} [get]
func (h *Handler) GetOfferModDetails(ctx *gin.Context) {
    offermodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offermod, err := h.Service.Repository.FindOfferModById(offermodId)
    if errors.Is(err, ErrOfferModNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offermod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferModResponse(offermod)
    ctx.JSON(http.StatusOK, response)
}

// CreateOfferMod godoc
// @Summary      Create offermod
// @Description  Create a new offermod with the provided information
// @Tags         offermod
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "OfferMod name"
// @Param        age   formData  int     true  "OfferMod age"
// @Param        image formData  file    true  "OfferMod image"
// @Success      201 {object} OfferModResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_mod/api/v1/ [post]
func (h *Handler) CreateOfferMod(ctx *gin.Context) {
    var req  OfferMod 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newOfferMod, err := h.Service.CreateOfferMod(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferModResponse(newOfferMod)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateOfferMod godoc
// @Summary      Update offermod
// @Description  Update offermod details by ID
// @Tags         offermod
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "OfferMod ID"
// @Param        name  formData string  false "OfferMod name"
// @Param        age   formData int     false "OfferMod age"
// @Param        image formData file    false "OfferMod image"
// @Success      200 {object} OfferModResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "OfferMod not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_mod/api/v1/{id} [put]
func (h *Handler) UpdateOfferMod(ctx *gin.Context) {
    offermodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offermod, err := h.Service.Repository.FindOfferModById(offermodId)
    if errors.Is(err, ErrOfferModNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offermod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateOfferModRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateOfferMod(&offermod,&req)
    
    if err := h.Service.UpdateOfferMod(offermod); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferModResponse(offermod)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         offermod
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "OfferMod ID"
// @Param        city body    CreateOfferModRequest true "Partial OfferMod information"
// @Success      200 {object} OfferModResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "offermod not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_mod/api/v1/{id} [patch]
func (h *Handler) UpdateOfferModPartial(ctx *gin.Context) {
    offermodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offermod, err := h.Service.Repository.FindOfferModById(offermodId)
    if errors.Is(err, ErrOfferModNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offermod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateOfferModRequest
    updateOfferMod(&offermod,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateOfferMod(offermod); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToOfferModResponse(offermod)
    ctx.JSON(http.StatusOK, response)
}


// DeleteOfferMod godoc
// @Summary      Delete offermod
// @Description  Delete a offermod by its ID
// @Tags         offermod
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "OfferMod ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "OfferMod not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer_mod/api/v1/{id} [delete]
func (h *Handler) DeleteOfferMod(ctx *gin.Context) {
    offermodId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    offermod, err := h.Service.Repository.FindOfferModById(offermodId)
    if errors.Is(err, ErrOfferModNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "offermod not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteOfferMod(offermod); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateOfferMod(offermod *OfferMod, req *CreateOfferModRequest) error {
	offermodVal := reflect.ValueOf(offermod).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			offermodField := offermodVal.FieldByName(reqVal.Type().Field(i).Name)
			if offermodField.IsValid() && offermodField.CanSet() {
				offermodField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
