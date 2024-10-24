package offer

import (
	"errors"
	"github.com/gin-gonic/gin"
	"ibrokers_service/pkg/middleware/filter/operators"
	"ibrokers_service/pkg/middleware/pagination"
	"ibrokers_service/pkg/utils/basics"
	"ibrokers_service/pkg/utils/manager"
	"net/http"
	"reflect"
	"strconv"
)

const BucketName = "offer"

var ErrOfferNotFound = errors.New("offer not found")

type Handler struct {
	Service     Service
	FileManager manager.FileManager
}

// ListOffers godoc
// @Summary      List of offers
// @Description  Get all offers
// @Tags         offer
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   OfferResponse
// @Router       /offer/api/v1/ [get]
func (h *Handler) GetOffer(ctx *gin.Context) {
	page := ctx.MustGet("page").(int)
	limit := ctx.MustGet("limit").(int)
	filters, _ := ctx.Get("filters")
	offers, count := h.Service.GetAllOffers(limit, page, filters.([]operators.FilterBlock))

	response := make([]OfferResponse, len(offers))
	for i, offer := range offers {
		response[i] = ToOfferResponse(offer)
	}
	paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

	ctx.JSON(http.StatusOK, paginationResponse)
}

// GetOfferDetails godoc
// @Summary      Get offer details
// @Description  Retrieve details of a offer by its ID
// @Tags         offer
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Offer ID"
// @Success      200 {object} OfferResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Offer not found"
// @Router       /offer/api/v1/{id} [get]
func (h *Handler) GetOfferDetails(ctx *gin.Context) {
	offerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	offer, err := h.Service.Repository.FindOfferById(offerId)
	if errors.Is(err, ErrOfferNotFound) {
		basics.ErrorResponse(ctx, http.StatusNotFound, "offer not found")
		return
	} else if err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := ToOfferResponse(offer)
	ctx.JSON(http.StatusOK, response)
}

// CreateOffer godoc
// @Summary      Create offer
// @Description  Create a new offer with the provided information
// @Tags         offer
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Offer name"
// @Param        age   formData  int     true  "Offer age"
// @Param        image formData  file    true  "Offer image"
// @Success      201 {object} OfferResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer/api/v1/ [post]
func (h *Handler) CreateOffer(ctx *gin.Context) {
	var req Offer
	if err := ctx.ShouldBind(&req); err != nil {
		basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	newOffer, err := h.Service.CreateOffer(req)
	if err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := ToOfferResponse(newOffer)
	ctx.JSON(http.StatusCreated, response)
}

// UpdateOffer godoc
// @Summary      Update offer
// @Description  Update offer details by ID
// @Tags         offer
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Offer ID"
// @Param        name  formData string  false "Offer name"
// @Param        age   formData int     false "Offer age"
// @Param        image formData file    false "Offer image"
// @Success      200 {object} OfferResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Offer not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer/api/v1/{id} [put]
func (h *Handler) UpdateOffer(ctx *gin.Context) {
	offerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	offer, err := h.Service.Repository.FindOfferById(offerId)
	if errors.Is(err, ErrOfferNotFound) {
		basics.ErrorResponse(ctx, http.StatusNotFound, "offer not found")
		return
	} else if err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var req CreateOfferRequest
	if err := ctx.ShouldBind(&req); err != nil {
		basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	updateOffer(&offer, &req)

	if err := h.Service.UpdateOffer(offer); err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := ToOfferResponse(offer)
	ctx.JSON(http.StatusOK, response)
}

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         offer
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Offer ID"
// @Param        city body    CreateOfferRequest true "Partial Offer information"
// @Success      200 {object} OfferResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "offer not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer/api/v1/{id} [patch]
func (h *Handler) UpdateOfferPartial(ctx *gin.Context) {
	offerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	offer, err := h.Service.Repository.FindOfferById(offerId)
	if errors.Is(err, ErrOfferNotFound) {
		basics.ErrorResponse(ctx, http.StatusNotFound, "offer not found")
		return
	} else if err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var req CreateOfferRequest
	updateOffer(&offer, &req)
	if err := ctx.ShouldBind(&req); err != nil {
		basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := h.Service.UpdateOffer(offer); err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := ToOfferResponse(offer)
	ctx.JSON(http.StatusOK, response)
}

// DeleteOffer godoc
// @Summary      Delete offer
// @Description  Delete a offer by its ID
// @Tags         offer
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Offer ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Offer not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /offer/api/v1/{id} [delete]
func (h *Handler) DeleteOffer(ctx *gin.Context) {
	offerId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	offer, err := h.Service.Repository.FindOfferById(offerId)
	if errors.Is(err, ErrOfferNotFound) {
		basics.ErrorResponse(ctx, http.StatusNotFound, "offer not found")
		return
	} else if err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Service.Repository.DeleteOffer(offer); err != nil {
		basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateOffer(offer *Offer, req *CreateOfferRequest) error {
	offerVal := reflect.ValueOf(offer).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			offerField := offerVal.FieldByName(reqVal.Type().Field(i).Name)
			if offerField.IsValid() && offerField.CanSet() {
				offerField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
