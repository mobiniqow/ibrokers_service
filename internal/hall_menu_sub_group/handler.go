package hall_menu_sub_group

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

const BucketName = "hall_menu_sub_group"

var ErrHallMenuSubGroupNotFound = errors.New("hallmenusubgroup not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListHallMenuSubGroups godoc
// @Summary      List of hallmenusubgroups
// @Description  Get all hallmenusubgroups
// @Tags         hallmenusubgroup
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   HallMenuSubGroupResponse
// @Router       /hall_menu_sub_group/api/v1/ [get]
func (h *Handler) GetHallMenuSubGroup(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    hallmenusubgroups, count := h.Service.GetAllHallMenuSubGroups(limit, page, filters.([]operators.FilterBlock))

    response := make([]HallMenuSubGroupResponse, len(hallmenusubgroups))
    for i, hallmenusubgroup := range hallmenusubgroups {
        response[i] = ToHallMenuSubGroupResponse(hallmenusubgroup)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetHallMenuSubGroupDetails godoc
// @Summary      Get hallmenusubgroup details
// @Description  Retrieve details of a hallmenusubgroup by its ID
// @Tags         hallmenusubgroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "HallMenuSubGroup ID"
// @Success      200 {object} HallMenuSubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "HallMenuSubGroup not found"
// @Router       /hall_menu_sub_group/api/v1/{id} [get]
func (h *Handler) GetHallMenuSubGroupDetails(ctx *gin.Context) {
    hallmenusubgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenusubgroup, err := h.Service.Repository.FindHallMenuSubGroupById(hallmenusubgroupId)
    if errors.Is(err, ErrHallMenuSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenusubgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuSubGroupResponse(hallmenusubgroup)
    ctx.JSON(http.StatusOK, response)
}

// CreateHallMenuSubGroup godoc
// @Summary      Create hallmenusubgroup
// @Description  Create a new hallmenusubgroup with the provided information
// @Tags         hallmenusubgroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "HallMenuSubGroup name"
// @Param        age   formData  int     true  "HallMenuSubGroup age"
// @Param        image formData  file    true  "HallMenuSubGroup image"
// @Success      201 {object} HallMenuSubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_sub_group/api/v1/ [post]
func (h *Handler) CreateHallMenuSubGroup(ctx *gin.Context) {
    var req  HallMenuSubGroup 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newHallMenuSubGroup, err := h.Service.CreateHallMenuSubGroup(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuSubGroupResponse(newHallMenuSubGroup)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateHallMenuSubGroup godoc
// @Summary      Update hallmenusubgroup
// @Description  Update hallmenusubgroup details by ID
// @Tags         hallmenusubgroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "HallMenuSubGroup ID"
// @Param        name  formData string  false "HallMenuSubGroup name"
// @Param        age   formData int     false "HallMenuSubGroup age"
// @Param        image formData file    false "HallMenuSubGroup image"
// @Success      200 {object} HallMenuSubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "HallMenuSubGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_sub_group/api/v1/{id} [put]
func (h *Handler) UpdateHallMenuSubGroup(ctx *gin.Context) {
    hallmenusubgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenusubgroup, err := h.Service.Repository.FindHallMenuSubGroupById(hallmenusubgroupId)
    if errors.Is(err, ErrHallMenuSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenusubgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateHallMenuSubGroupRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateHallMenuSubGroup(&hallmenusubgroup,&req)
    
    if err := h.Service.UpdateHallMenuSubGroup(hallmenusubgroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuSubGroupResponse(hallmenusubgroup)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         hallmenusubgroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "HallMenuSubGroup ID"
// @Param        city body    CreateHallMenuSubGroupRequest true "Partial HallMenuSubGroup information"
// @Success      200 {object} HallMenuSubGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "hallmenusubgroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_sub_group/api/v1/{id} [patch]
func (h *Handler) UpdateHallMenuSubGroupPartial(ctx *gin.Context) {
    hallmenusubgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenusubgroup, err := h.Service.Repository.FindHallMenuSubGroupById(hallmenusubgroupId)
    if errors.Is(err, ErrHallMenuSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenusubgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateHallMenuSubGroupRequest
    updateHallMenuSubGroup(&hallmenusubgroup,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateHallMenuSubGroup(hallmenusubgroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuSubGroupResponse(hallmenusubgroup)
    ctx.JSON(http.StatusOK, response)
}


// DeleteHallMenuSubGroup godoc
// @Summary      Delete hallmenusubgroup
// @Description  Delete a hallmenusubgroup by its ID
// @Tags         hallmenusubgroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "HallMenuSubGroup ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "HallMenuSubGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_sub_group/api/v1/{id} [delete]
func (h *Handler) DeleteHallMenuSubGroup(ctx *gin.Context) {
    hallmenusubgroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenusubgroup, err := h.Service.Repository.FindHallMenuSubGroupById(hallmenusubgroupId)
    if errors.Is(err, ErrHallMenuSubGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenusubgroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteHallMenuSubGroup(hallmenusubgroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateHallMenuSubGroup(hallmenusubgroup *HallMenuSubGroup, req *CreateHallMenuSubGroupRequest) error {
	hallmenusubgroupVal := reflect.ValueOf(hallmenusubgroup).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			hallmenusubgroupField := hallmenusubgroupVal.FieldByName(reqVal.Type().Field(i).Name)
			if hallmenusubgroupField.IsValid() && hallmenusubgroupField.CanSet() {
				hallmenusubgroupField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
