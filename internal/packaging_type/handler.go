package packaging_type

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

const BucketName = "packaging_type"

var ErrPackagingTypeNotFound = errors.New("packagingtype not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListPackagingTypes godoc
// @Summary      List of packagingtypes
// @Description  Get all packagingtypes
// @Tags         packagingtype
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   PackagingTypeResponse
// @Router       /packaging_type/api/v1/ [get]
func (h *Handler) GetPackagingType(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    packagingtypes, count := h.Service.GetAllPackagingTypes(limit, page, filters.([]operators.FilterBlock))

    response := make([]PackagingTypeResponse, len(packagingtypes))
    for i, packagingtype := range packagingtypes {
        response[i] = ToPackagingTypeResponse(packagingtype)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetPackagingTypeDetails godoc
// @Summary      Get packagingtype details
// @Description  Retrieve details of a packagingtype by its ID
// @Tags         packagingtype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "PackagingType ID"
// @Success      200 {object} PackagingTypeResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "PackagingType not found"
// @Router       /packaging_type/api/v1/{id} [get]
func (h *Handler) GetPackagingTypeDetails(ctx *gin.Context) {
    packagingtypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    packagingtype, err := h.Service.Repository.FindPackagingTypeById(packagingtypeId)
    if errors.Is(err, ErrPackagingTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "packagingtype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToPackagingTypeResponse(packagingtype)
    ctx.JSON(http.StatusOK, response)
}

// CreatePackagingType godoc
// @Summary      Create packagingtype
// @Description  Create a new packagingtype with the provided information
// @Tags         packagingtype
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "PackagingType name"
// @Param        age   formData  int     true  "PackagingType age"
// @Param        image formData  file    true  "PackagingType image"
// @Success      201 {object} PackagingTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /packaging_type/api/v1/ [post]
func (h *Handler) CreatePackagingType(ctx *gin.Context) {
    var req  PackagingType 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newPackagingType, err := h.Service.CreatePackagingType(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToPackagingTypeResponse(newPackagingType)
    ctx.JSON(http.StatusCreated, response)
}

// UpdatePackagingType godoc
// @Summary      Update packagingtype
// @Description  Update packagingtype details by ID
// @Tags         packagingtype
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "PackagingType ID"
// @Param        name  formData string  false "PackagingType name"
// @Param        age   formData int     false "PackagingType age"
// @Param        image formData file    false "PackagingType image"
// @Success      200 {object} PackagingTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "PackagingType not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /packaging_type/api/v1/{id} [put]
func (h *Handler) UpdatePackagingType(ctx *gin.Context) {
    packagingtypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    packagingtype, err := h.Service.Repository.FindPackagingTypeById(packagingtypeId)
    if errors.Is(err, ErrPackagingTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "packagingtype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreatePackagingTypeRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updatePackagingType(&packagingtype,&req)
    
    if err := h.Service.UpdatePackagingType(packagingtype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToPackagingTypeResponse(packagingtype)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         packagingtype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "PackagingType ID"
// @Param        city body    CreatePackagingTypeRequest true "Partial PackagingType information"
// @Success      200 {object} PackagingTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "packagingtype not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /packaging_type/api/v1/{id} [patch]
func (h *Handler) UpdatePackagingTypePartial(ctx *gin.Context) {
    packagingtypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    packagingtype, err := h.Service.Repository.FindPackagingTypeById(packagingtypeId)
    if errors.Is(err, ErrPackagingTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "packagingtype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreatePackagingTypeRequest
    updatePackagingType(&packagingtype,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdatePackagingType(packagingtype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToPackagingTypeResponse(packagingtype)
    ctx.JSON(http.StatusOK, response)
}


// DeletePackagingType godoc
// @Summary      Delete packagingtype
// @Description  Delete a packagingtype by its ID
// @Tags         packagingtype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "PackagingType ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "PackagingType not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /packaging_type/api/v1/{id} [delete]
func (h *Handler) DeletePackagingType(ctx *gin.Context) {
    packagingtypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    packagingtype, err := h.Service.Repository.FindPackagingTypeById(packagingtypeId)
    if errors.Is(err, ErrPackagingTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "packagingtype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeletePackagingType(packagingtype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updatePackagingType(packagingtype *PackagingType, req *CreatePackagingTypeRequest) error {
	packagingtypeVal := reflect.ValueOf(packagingtype).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			packagingtypeField := packagingtypeVal.FieldByName(reqVal.Type().Field(i).Name)
			if packagingtypeField.IsValid() && packagingtypeField.CanSet() {
				packagingtypeField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
