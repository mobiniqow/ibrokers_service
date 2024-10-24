package sub_group

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

const BucketName = "sub_group"

var ErrSubGroupNotFound = errors.New("subgroup not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListSubGroups godoc
// @Summary      List of subgroups
// @Description  Get all subgroups
// @Tags         subgroup
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   SubGroupResponse
// @Router       /sub_group/api/v1/ [get]
func (h *Handler) GetSubGroup(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    subgroups, count := h.Service.GetAllSubGroups(limit, page, filters.([]operators.FilterBlock))

    response := make([]SubGroupResponse, len(subgroups))
    for i, subgroup := range subgroups {
        response[i] = ToSubGroupResponse(subgroup)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetSubGroupDetails godoc
// @Summary      Get subgroup details
// @Description  Retrieve details of a subgroup by its ID
// @Tags         subgroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "SubGroup ID"
// @Success      200 {object} SubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "SubGroup not found"
// @Router       /sub_group/api/v1/{id} [get]
func (h *Handler) GetSubGroupDetails(ctx *gin.Context) {
    subgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    subgroup, err := h.Service.Repository.FindSubGroupById(subgroupId)
    if errors.Is(err, ErrSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "subgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSubGroupResponse(subgroup)
    ctx.JSON(http.StatusOK, response)
}

// CreateSubGroup godoc
// @Summary      Create subgroup
// @Description  Create a new subgroup with the provided information
// @Tags         subgroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "SubGroup name"
// @Param        age   formData  int     true  "SubGroup age"
// @Param        image formData  file    true  "SubGroup image"
// @Success      201 {object} SubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /sub_group/api/v1/ [post]
func (h *Handler) CreateSubGroup(ctx *gin.Context) {
    var req  SubGroup 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newSubGroup, err := h.Service.CreateSubGroup(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSubGroupResponse(newSubGroup)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateSubGroup godoc
// @Summary      Update subgroup
// @Description  Update subgroup details by ID
// @Tags         subgroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "SubGroup ID"
// @Param        name  formData string  false "SubGroup name"
// @Param        age   formData int     false "SubGroup age"
// @Param        image formData file    false "SubGroup image"
// @Success      200 {object} SubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "SubGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /sub_group/api/v1/{id} [put]
func (h *Handler) UpdateSubGroup(ctx *gin.Context) {
    subgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    subgroup, err := h.Service.Repository.FindSubGroupById(subgroupId)
    if errors.Is(err, ErrSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "subgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateSubGroupRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateSubGroup(&subgroup,&req)
    
    if err := h.Service.UpdateSubGroup(subgroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSubGroupResponse(subgroup)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         subgroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "SubGroup ID"
// @Param        city body    CreateSubGroupRequest true "Partial SubGroup information"
// @Success      200 {object} SubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "subgroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /sub_group/api/v1/{id} [patch]
func (h *Handler) UpdateSubGroupPartial(ctx *gin.Context) {
    subgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    subgroup, err := h.Service.Repository.FindSubGroupById(subgroupId)
    if errors.Is(err, ErrSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "subgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateSubGroupRequest
    updateSubGroup(&subgroup,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateSubGroup(subgroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToSubGroupResponse(subgroup)
    ctx.JSON(http.StatusOK, response)
}


// DeleteSubGroup godoc
// @Summary      Delete subgroup
// @Description  Delete a subgroup by its ID
// @Tags         subgroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "SubGroup ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "SubGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /sub_group/api/v1/{id} [delete]
func (h *Handler) DeleteSubGroup(ctx *gin.Context) {
    subgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    subgroup, err := h.Service.Repository.FindSubGroupById(subgroupId)
    if errors.Is(err, ErrSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "subgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteSubGroup(subgroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateSubGroup(subgroup *SubGroup, req *CreateSubGroupRequest) error {
	subgroupVal := reflect.ValueOf(subgroup).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			subgroupField := subgroupVal.FieldByName(reqVal.Type().Field(i).Name)
			if subgroupField.IsValid() && subgroupField.CanSet() {
				subgroupField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
