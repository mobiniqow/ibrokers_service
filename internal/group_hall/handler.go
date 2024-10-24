package group_hall

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

const BucketName = "group_hall"

var ErrGroupHallNotFound = errors.New("grouphall not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListGroupHalls godoc
// @Summary      List of grouphalls
// @Description  Get all grouphalls
// @Tags         grouphall
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   GroupHallResponse
// @Router       /group_hall/api/v1/ [get]
func (h *Handler) GetGroupHall(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    grouphalls, count := h.Service.GetAllGroupHalls(limit, page, filters.([]operators.FilterBlock))

    response := make([]GroupHallResponse, len(grouphalls))
    for i, grouphall := range grouphalls {
        response[i] = ToGroupHallResponse(grouphall)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetGroupHallDetails godoc
// @Summary      Get grouphall details
// @Description  Retrieve details of a grouphall by its ID
// @Tags         grouphall
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "GroupHall ID"
// @Success      200 {object} GroupHallResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "GroupHall not found"
// @Router       /group_hall/api/v1/{id} [get]
func (h *Handler) GetGroupHallDetails(ctx *gin.Context) {
    grouphallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    grouphall, err := h.Service.Repository.FindGroupHallById(grouphallId)
    if errors.Is(err, ErrGroupHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "grouphall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupHallResponse(grouphall)
    ctx.JSON(http.StatusOK, response)
}

// CreateGroupHall godoc
// @Summary      Create grouphall
// @Description  Create a new grouphall with the provided information
// @Tags         grouphall
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "GroupHall name"
// @Param        age   formData  int     true  "GroupHall age"
// @Param        image formData  file    true  "GroupHall image"
// @Success      201 {object} GroupHallResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group_hall/api/v1/ [post]
func (h *Handler) CreateGroupHall(ctx *gin.Context) {
    var req  GroupHall 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newGroupHall, err := h.Service.CreateGroupHall(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupHallResponse(newGroupHall)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateGroupHall godoc
// @Summary      Update grouphall
// @Description  Update grouphall details by ID
// @Tags         grouphall
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "GroupHall ID"
// @Param        name  formData string  false "GroupHall name"
// @Param        age   formData int     false "GroupHall age"
// @Param        image formData file    false "GroupHall image"
// @Success      200 {object} GroupHallResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "GroupHall not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group_hall/api/v1/{id} [put]
func (h *Handler) UpdateGroupHall(ctx *gin.Context) {
    grouphallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    grouphall, err := h.Service.Repository.FindGroupHallById(grouphallId)
    if errors.Is(err, ErrGroupHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "grouphall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateGroupHallRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateGroupHall(&grouphall,&req)
    
    if err := h.Service.UpdateGroupHall(grouphall); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupHallResponse(grouphall)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         grouphall
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "GroupHall ID"
// @Param        city body    CreateGroupHallRequest true "Partial GroupHall information"
// @Success      200 {object} GroupHallResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "grouphall not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group_hall/api/v1/{id} [patch]
func (h *Handler) UpdateGroupHallPartial(ctx *gin.Context) {
    grouphallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    grouphall, err := h.Service.Repository.FindGroupHallById(grouphallId)
    if errors.Is(err, ErrGroupHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "grouphall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateGroupHallRequest
    updateGroupHall(&grouphall,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateGroupHall(grouphall); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToGroupHallResponse(grouphall)
    ctx.JSON(http.StatusOK, response)
}


// DeleteGroupHall godoc
// @Summary      Delete grouphall
// @Description  Delete a grouphall by its ID
// @Tags         grouphall
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "GroupHall ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "GroupHall not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /group_hall/api/v1/{id} [delete]
func (h *Handler) DeleteGroupHall(ctx *gin.Context) {
    grouphallId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    grouphall, err := h.Service.Repository.FindGroupHallById(grouphallId)
    if errors.Is(err, ErrGroupHallNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "grouphall not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteGroupHall(grouphall); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateGroupHall(grouphall *GroupHall, req *CreateGroupHallRequest) error {
	grouphallVal := reflect.ValueOf(grouphall).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			grouphallField := grouphallVal.FieldByName(reqVal.Type().Field(i).Name)
			if grouphallField.IsValid() && grouphallField.CanSet() {
				grouphallField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
