package hall_menu_group

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

const BucketName = "hall_menu_group"

var ErrHallMenuGroupNotFound = errors.New("hallmenugroup not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListHallMenuGroups godoc
// @Summary      List of hallmenugroups
// @Description  Get all hallmenugroups
// @Tags         hallmenugroup
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   HallMenuGroupResponse
// @Router       /hall_menu_group/api/v1/ [get]
func (h *Handler) GetHallMenuGroup(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    hallmenugroups, count := h.Service.GetAllHallMenuGroups(limit, page, filters.([]operators.FilterBlock))

    response := make([]HallMenuGroupResponse, len(hallmenugroups))
    for i, hallmenugroup := range hallmenugroups {
        response[i] = ToHallMenuGroupResponse(hallmenugroup)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetHallMenuGroupDetails godoc
// @Summary      Get hallmenugroup details
// @Description  Retrieve details of a hallmenugroup by its ID
// @Tags         hallmenugroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "HallMenuGroup ID"
// @Success      200 {object} HallMenuGroupResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "HallMenuGroup not found"
// @Router       /hall_menu_group/api/v1/{id} [get]
func (h *Handler) GetHallMenuGroupDetails(ctx *gin.Context) {
    hallmenugroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenugroup, err := h.Service.Repository.FindHallMenuGroupById(hallmenugroupId)
    if errors.Is(err, ErrHallMenuGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenugroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuGroupResponse(hallmenugroup)
    ctx.JSON(http.StatusOK, response)
}

// CreateHallMenuGroup godoc
// @Summary      Create hallmenugroup
// @Description  Create a new hallmenugroup with the provided information
// @Tags         hallmenugroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "HallMenuGroup name"
// @Param        age   formData  int     true  "HallMenuGroup age"
// @Param        image formData  file    true  "HallMenuGroup image"
// @Success      201 {object} HallMenuGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_group/api/v1/ [post]
func (h *Handler) CreateHallMenuGroup(ctx *gin.Context) {
    var req  HallMenuGroup 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newHallMenuGroup, err := h.Service.CreateHallMenuGroup(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuGroupResponse(newHallMenuGroup)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateHallMenuGroup godoc
// @Summary      Update hallmenugroup
// @Description  Update hallmenugroup details by ID
// @Tags         hallmenugroup
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "HallMenuGroup ID"
// @Param        name  formData string  false "HallMenuGroup name"
// @Param        age   formData int     false "HallMenuGroup age"
// @Param        image formData file    false "HallMenuGroup image"
// @Success      200 {object} HallMenuGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "HallMenuGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_group/api/v1/{id} [put]
func (h *Handler) UpdateHallMenuGroup(ctx *gin.Context) {
    hallmenugroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenugroup, err := h.Service.Repository.FindHallMenuGroupById(hallmenugroupId)
    if errors.Is(err, ErrHallMenuGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenugroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateHallMenuGroupRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateHallMenuGroup(&hallmenugroup,&req)
    
    if err := h.Service.UpdateHallMenuGroup(hallmenugroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuGroupResponse(hallmenugroup)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         hallmenugroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "HallMenuGroup ID"
// @Param        city body    CreateHallMenuGroupRequest true "Partial HallMenuGroup information"
// @Success      200 {object} HallMenuGroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "hallmenugroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_group/api/v1/{id} [patch]
func (h *Handler) UpdateHallMenuGroupPartial(ctx *gin.Context) {
    hallmenugroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenugroup, err := h.Service.Repository.FindHallMenuGroupById(hallmenugroupId)
    if errors.Is(err, ErrHallMenuGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenugroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateHallMenuGroupRequest
    updateHallMenuGroup(&hallmenugroup,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateHallMenuGroup(hallmenugroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToHallMenuGroupResponse(hallmenugroup)
    ctx.JSON(http.StatusOK, response)
}


// DeleteHallMenuGroup godoc
// @Summary      Delete hallmenugroup
// @Description  Delete a hallmenugroup by its ID
// @Tags         hallmenugroup
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "HallMenuGroup ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "HallMenuGroup not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /hall_menu_group/api/v1/{id} [delete]
func (h *Handler) DeleteHallMenuGroup(ctx *gin.Context) {
    hallmenugroupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    hallmenugroup, err := h.Service.Repository.FindHallMenuGroupById(hallmenugroupId)
    if errors.Is(err, ErrHallMenuGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "hallmenugroup not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteHallMenuGroup(hallmenugroup); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateHallMenuGroup(hallmenugroup *HallMenuGroup, req *CreateHallMenuGroupRequest) error {
	hallmenugroupVal := reflect.ValueOf(hallmenugroup).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			hallmenugroupField := hallmenugroupVal.FieldByName(reqVal.Type().Field(i).Name)
			if hallmenugroupField.IsValid() && hallmenugroupField.CanSet() {
				hallmenugroupField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
