package broker

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

const BucketName = "broker"

var ErrBrokerNotFound = errors.New("broker not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListBrokers godoc
// @Summary      List of brokers
// @Description  Get all brokers
// @Tags         broker
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   BrokerResponse
// @Router       /broker/api/v1/ [get]
func (h *Handler) GetBroker(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    brokers, count := h.Service.GetAllBrokers(limit, page, filters.([]operators.FilterBlock))

    response := make([]BrokerResponse, len(brokers))
    for i, broker := range brokers {
        response[i] = ToBrokerResponse(broker)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetBrokerDetails godoc
// @Summary      Get broker details
// @Description  Retrieve details of a broker by its ID
// @Tags         broker
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Broker ID"
// @Success      200 {object} BrokerResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Broker not found"
// @Router       /broker/api/v1/{id} [get]
func (h *Handler) GetBrokerDetails(ctx *gin.Context) {
    brokerId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    broker, err := h.Service.Repository.FindBrokerById(brokerId)
    if errors.Is(err, ErrBrokerNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "broker not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBrokerResponse(broker)
    ctx.JSON(http.StatusOK, response)
}

// CreateBroker godoc
// @Summary      Create broker
// @Description  Create a new broker with the provided information
// @Tags         broker
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Broker name"
// @Param        age   formData  int     true  "Broker age"
// @Param        image formData  file    true  "Broker image"
// @Success      201 {object} BrokerResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /broker/api/v1/ [post]
func (h *Handler) CreateBroker(ctx *gin.Context) {
    var req  Broker 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newBroker, err := h.Service.CreateBroker(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBrokerResponse(newBroker)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateBroker godoc
// @Summary      Update broker
// @Description  Update broker details by ID
// @Tags         broker
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Broker ID"
// @Param        name  formData string  false "Broker name"
// @Param        age   formData int     false "Broker age"
// @Param        image formData file    false "Broker image"
// @Success      200 {object} BrokerResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Broker not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /broker/api/v1/{id} [put]
func (h *Handler) UpdateBroker(ctx *gin.Context) {
    brokerId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    broker, err := h.Service.Repository.FindBrokerById(brokerId)
    if errors.Is(err, ErrBrokerNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "broker not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateBrokerRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateBroker(&broker,&req)
    
    if err := h.Service.UpdateBroker(broker); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBrokerResponse(broker)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         broker
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Broker ID"
// @Param        city body    CreateBrokerRequest true "Partial Broker information"
// @Success      200 {object} BrokerResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "broker not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /broker/api/v1/{id} [patch]
func (h *Handler) UpdateBrokerPartial(ctx *gin.Context) {
    brokerId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    broker, err := h.Service.Repository.FindBrokerById(brokerId)
    if errors.Is(err, ErrBrokerNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "broker not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateBrokerRequest
    updateBroker(&broker,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateBroker(broker); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToBrokerResponse(broker)
    ctx.JSON(http.StatusOK, response)
}


// DeleteBroker godoc
// @Summary      Delete broker
// @Description  Delete a broker by its ID
// @Tags         broker
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Broker ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Broker not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /broker/api/v1/{id} [delete]
func (h *Handler) DeleteBroker(ctx *gin.Context) {
    brokerId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    broker, err := h.Service.Repository.FindBrokerById(brokerId)
    if errors.Is(err, ErrBrokerNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "broker not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteBroker(broker); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateBroker(broker *Broker, req *CreateBrokerRequest) error {
	brokerVal := reflect.ValueOf(broker).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			brokerField := brokerVal.FieldByName(reqVal.Type().Field(i).Name)
			if brokerField.IsValid() && brokerField.CanSet() {
				brokerField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
