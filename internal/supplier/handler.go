package supplier

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

const BucketName = "supplier"

var ErrSupplierNotFound = errors.New("supplier not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListSuppliers godoc
// @Summary      List of suppliers
// @Description  Get all suppliers
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   SupplierResponse
// @Router       /supplier/api/v1/ [get]
func (h *Handler) GetSupplier(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    suppliers, count := h.Service.GetAllSuppliers(limit, page, filters.([]operators.FilterBlock))

    response := make([]SupplierResponse, len(suppliers))
    for i, supplier := range suppliers {
        response[i] = ToSupplierResponse(supplier)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetSupplierDetails godoc
// @Summary      Get supplier details
// @Description  Retrieve details of a supplier by its ID
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Supplier ID"
// @Success      200 {object} SupplierResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Supplier not found"
// @Router       /supplier/api/v1/{id} [get]
func (h *Handler) GetSupplierDetails(ctx *gin.Context) {
    supplierId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    supplier, err := h.Service.Repository.FindSupplierById(supplierId)
    if errors.Is(err, ErrSupplierNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "supplier not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSupplierResponse(supplier)
    ctx.JSON(http.StatusOK, response)
}

// CreateSupplier godoc
// @Summary      Create supplier
// @Description  Create a new supplier with the provided information
// @Tags         supplier
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Supplier name"
// @Param        age   formData  int     true  "Supplier age"
// @Param        image formData  file    true  "Supplier image"
// @Success      201 {object} SupplierResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /supplier/api/v1/ [post]
func (h *Handler) CreateSupplier(ctx *gin.Context) {
    var req  Supplier 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newSupplier, err := h.Service.CreateSupplier(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSupplierResponse(newSupplier)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateSupplier godoc
// @Summary      Update supplier
// @Description  Update supplier details by ID
// @Tags         supplier
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Supplier ID"
// @Param        name  formData string  false "Supplier name"
// @Param        age   formData int     false "Supplier age"
// @Param        image formData file    false "Supplier image"
// @Success      200 {object} SupplierResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Supplier not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /supplier/api/v1/{id} [put]
func (h *Handler) UpdateSupplier(ctx *gin.Context) {
    supplierId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    supplier, err := h.Service.Repository.FindSupplierById(supplierId)
    if errors.Is(err, ErrSupplierNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "supplier not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateSupplierRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateSupplier(&supplier,&req)
    
    if err := h.Service.UpdateSupplier(supplier); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSupplierResponse(supplier)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Supplier ID"
// @Param        city body    CreateSupplierRequest true "Partial Supplier information"
// @Success      200 {object} SupplierResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "supplier not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /supplier/api/v1/{id} [patch]
func (h *Handler) UpdateSupplierPartial(ctx *gin.Context) {
    supplierId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    supplier, err := h.Service.Repository.FindSupplierById(supplierId)
    if errors.Is(err, ErrSupplierNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "supplier not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateSupplierRequest
    updateSupplier(&supplier,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateSupplier(supplier); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSupplierResponse(supplier)
    ctx.JSON(http.StatusOK, response)
}


// DeleteSupplier godoc
// @Summary      Delete supplier
// @Description  Delete a supplier by its ID
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Supplier ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Supplier not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /supplier/api/v1/{id} [delete]
func (h *Handler) DeleteSupplier(ctx *gin.Context) {
    supplierId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    supplier, err := h.Service.Repository.FindSupplierById(supplierId)
    if errors.Is(err, ErrSupplierNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "supplier not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteSupplier(supplier); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateSupplier(supplier *Supplier, req *CreateSupplierRequest) error {
	supplierVal := reflect.ValueOf(supplier).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			supplierField := supplierVal.FieldByName(reqVal.Type().Field(i).Name)
			if supplierField.IsValid() && supplierField.CanSet() {
				supplierField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
