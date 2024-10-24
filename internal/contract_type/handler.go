package contract_type

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

const BucketName = "contract_type"

var ErrContractTypeNotFound = errors.New("contracttype not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListContractTypes godoc
// @Summary      List of contracttypes
// @Description  Get all contracttypes
// @Tags         contracttype
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   ContractTypeResponse
// @Router       /contract_type/api/v1/ [get]
func (h *Handler) GetContractType(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    contracttypes, count := h.Service.GetAllContractTypes(limit, page, filters.([]operators.FilterBlock))

    response := make([]ContractTypeResponse, len(contracttypes))
    for i, contracttype := range contracttypes {
        response[i] = ToContractTypeResponse(contracttype)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetContractTypeDetails godoc
// @Summary      Get contracttype details
// @Description  Retrieve details of a contracttype by its ID
// @Tags         contracttype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "ContractType ID"
// @Success      200 {object} ContractTypeResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "ContractType not found"
// @Router       /contract_type/api/v1/{id} [get]
func (h *Handler) GetContractTypeDetails(ctx *gin.Context) {
    contracttypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    contracttype, err := h.Service.Repository.FindContractTypeById(contracttypeId)
    if errors.Is(err, ErrContractTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "contracttype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToContractTypeResponse(contracttype)
    ctx.JSON(http.StatusOK, response)
}

// CreateContractType godoc
// @Summary      Create contracttype
// @Description  Create a new contracttype with the provided information
// @Tags         contracttype
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "ContractType name"
// @Param        age   formData  int     true  "ContractType age"
// @Param        image formData  file    true  "ContractType image"
// @Success      201 {object} ContractTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /contract_type/api/v1/ [post]
func (h *Handler) CreateContractType(ctx *gin.Context) {
    var req  ContractType 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newContractType, err := h.Service.CreateContractType(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToContractTypeResponse(newContractType)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateContractType godoc
// @Summary      Update contracttype
// @Description  Update contracttype details by ID
// @Tags         contracttype
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "ContractType ID"
// @Param        name  formData string  false "ContractType name"
// @Param        age   formData int     false "ContractType age"
// @Param        image formData file    false "ContractType image"
// @Success      200 {object} ContractTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "ContractType not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /contract_type/api/v1/{id} [put]
func (h *Handler) UpdateContractType(ctx *gin.Context) {
    contracttypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    contracttype, err := h.Service.Repository.FindContractTypeById(contracttypeId)
    if errors.Is(err, ErrContractTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "contracttype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateContractTypeRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateContractType(&contracttype,&req)
    
    if err := h.Service.UpdateContractType(contracttype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToContractTypeResponse(contracttype)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         contracttype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "ContractType ID"
// @Param        city body    CreateContractTypeRequest true "Partial ContractType information"
// @Success      200 {object} ContractTypeResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "contracttype not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /contract_type/api/v1/{id} [patch]
func (h *Handler) UpdateContractTypePartial(ctx *gin.Context) {
    contracttypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    contracttype, err := h.Service.Repository.FindContractTypeById(contracttypeId)
    if errors.Is(err, ErrContractTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "contracttype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateContractTypeRequest
    updateContractType(&contracttype,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateContractType(contracttype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToContractTypeResponse(contracttype)
    ctx.JSON(http.StatusOK, response)
}


// DeleteContractType godoc
// @Summary      Delete contracttype
// @Description  Delete a contracttype by its ID
// @Tags         contracttype
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "ContractType ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "ContractType not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /contract_type/api/v1/{id} [delete]
func (h *Handler) DeleteContractType(ctx *gin.Context) {
    contracttypeId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    contracttype, err := h.Service.Repository.FindContractTypeById(contracttypeId)
    if errors.Is(err, ErrContractTypeNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "contracttype not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteContractType(contracttype); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateContractType(contracttype *ContractType, req *CreateContractTypeRequest) error {
	contracttypeVal := reflect.ValueOf(contracttype).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			contracttypeField := contracttypeVal.FieldByName(reqVal.Type().Field(i).Name)
			if contracttypeField.IsValid() && contracttypeField.CanSet() {
				contracttypeField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
