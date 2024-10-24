package group

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

const BucketName = "group"

var ErrGroupNotFound = errors.New("group not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListGroups godoc
// @Summary      List of groups
// @Description  Get all groups
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   GroupResponse
// @Router       /group/api/v1/ [get]
func (h *Handler) GetGroup(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    groups, count := h.Service.GetAllGroups(limit, page, filters.([]operators.FilterBlock))

    response := make([]GroupResponse, len(groups))
    for i, group := range groups {
        response[i] = ToGroupResponse(group)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetGroupDetails godoc
// @Summary      Get group details
// @Description  Retrieve details of a group by its ID
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Group ID"
// @Success      200 {object} GroupResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Group not found"
// @Router       /group/api/v1/{id} [get]
func (h *Handler) GetGroupDetails(ctx *gin.Context) {
    groupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    group, err := h.Service.Repository.FindGroupById(groupId)
    if errors.Is(err, ErrGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "group not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupResponse(group)
    ctx.JSON(http.StatusOK, response)
}

// CreateGroup godoc
// @Summary      Create group
// @Description  Create a new group with the provided information
// @Tags         group
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Group name"
// @Param        age   formData  int     true  "Group age"
// @Param        image formData  file    true  "Group image"
// @Success      201 {object} GroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group/api/v1/ [post]
func (h *Handler) CreateGroup(ctx *gin.Context) {
    var req  Group 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newGroup, err := h.Service.CreateGroup(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupResponse(newGroup)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateGroup godoc
// @Summary      Update group
// @Description  Update group details by ID
// @Tags         group
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Group ID"
// @Param        name  formData string  false "Group name"
// @Param        age   formData int     false "Group age"
// @Param        image formData file    false "Group image"
// @Success      200 {object} GroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Group not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group/api/v1/{id} [put]
func (h *Handler) UpdateGroup(ctx *gin.Context) {
    groupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    group, err := h.Service.Repository.FindGroupById(groupId)
    if errors.Is(err, ErrGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "group not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateGroupRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateGroup(&group,&req)
    
    if err := h.Service.UpdateGroup(group); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupResponse(group)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Group ID"
// @Param        city body    CreateGroupRequest true "Partial Group information"
// @Success      200 {object} GroupResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "group not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group/api/v1/{id} [patch]
func (h *Handler) UpdateGroupPartial(ctx *gin.Context) {
    groupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    group, err := h.Service.Repository.FindGroupById(groupId)
    if errors.Is(err, ErrGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "group not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateGroupRequest
    updateGroup(&group,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateGroup(group); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupResponse(group)
    ctx.JSON(http.StatusOK, response)
}


// DeleteGroup godoc
// @Summary      Delete group
// @Description  Delete a group by its ID
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Group ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Group not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group/api/v1/{id} [delete]
func (h *Handler) DeleteGroup(ctx *gin.Context) {
    groupId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    group, err := h.Service.Repository.FindGroupById(groupId)
    if errors.Is(err, ErrGroupNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "group not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteGroup(group); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateGroup(group *Group, req *CreateGroupRequest) error {
	groupVal := reflect.ValueOf(group).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			groupField := groupVal.FieldByName(reqVal.Type().Field(i).Name)
			if groupField.IsValid() && groupField.CanSet() {
				groupField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
