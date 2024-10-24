package report

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

const BucketName = "report"

var ErrReportNotFound = errors.New("report not found")

type Handler struct {
    Service     Service
    FileManager manager.FileManager
}

// ListReports godoc
// @Summary      List of reports
// @Description  Get all reports
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        name    query     string  false  "Search by name"
// @Param        age     query     string  false  "Search by age"
// @Param        page    query     int     false  "page number"
// @Param        limit   query     int     false  "page size"
// @Success      200     {array}   ReportResponse
// @Router       /report/api/v1/ [get]
func (h *Handler) GetReport(ctx *gin.Context) {
    page := ctx.MustGet("page").(int)
    limit := ctx.MustGet("limit").(int)
    filters, _ := ctx.Get("filters")
    reports, count := h.Service.GetAllReports(limit, page, filters.([]operators.FilterBlock))

    response := make([]ReportResponse, len(reports))
    for i, report := range reports {
        response[i] = ToReportResponse(report)
    }
    paginationResponse := pagination.GenerateResponse(limit, page, count, ctx, response)

    ctx.JSON(http.StatusOK, paginationResponse)
}

// GetReportDetails godoc
// @Summary      Get report details
// @Description  Retrieve details of a report by its ID
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Report ID"
// @Success      200 {object} ReportResponse
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Report not found"
// @Router       /report/api/v1/{id} [get]
func (h *Handler) GetReportDetails(ctx *gin.Context) {
    reportId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    report, err := h.Service.Repository.FindReportById(reportId)
    if errors.Is(err, ErrReportNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "report not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToReportResponse(report)
    ctx.JSON(http.StatusOK, response)
}

// CreateReport godoc
// @Summary      Create report
// @Description  Create a new report with the provided information
// @Tags         report
// @Accept       multipart/form-data
// @Produce      json
// @Param        name  formData  string  true  "Report name"
// @Param        age   formData  int     true  "Report age"
// @Param        image formData  file    true  "Report image"
// @Success      201 {object} ReportResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /report/api/v1/ [post]
func (h *Handler) CreateReport(ctx *gin.Context) {
    var req  Report 
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }  

    newReport, err := h.Service.CreateReport(req)
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToReportResponse(newReport)
    ctx.JSON(http.StatusCreated, response)
}

// UpdateReport godoc
// @Summary      Update report
// @Description  Update report details by ID
// @Tags         report
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path    string  true  "Report ID"
// @Param        name  formData string  false "Report name"
// @Param        age   formData int     false "Report age"
// @Param        image formData file    false "Report image"
// @Success      200 {object} ReportResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "Report not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /report/api/v1/{id} [put]
func (h *Handler) UpdateReport(ctx *gin.Context) {
    reportId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    report, err := h.Service.Repository.FindReportById(reportId)
    if errors.Is(err, ErrReportNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "report not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }
    var req CreateReportRequest
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }

    updateReport(&report,&req)
    
    if err := h.Service.UpdateReport(report); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToReportResponse(report)
    ctx.JSON(http.StatusOK, response)
} 

// UpdateCityPartial godoc
// @Summary      Update city partially
// @Description  Update specific fields of a city by ID
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Report ID"
// @Param        city body    CreateReportRequest true "Partial Report information"
// @Success      200 {object} ReportResponse
// @Failure      400 {object} basics.APIError "Invalid request"
// @Failure      404 {object} basics.APIError "report not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /report/api/v1/{id} [patch]
func (h *Handler) UpdateReportPartial(ctx *gin.Context) {
    reportId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    report, err := h.Service.Repository.FindReportById(reportId)
    if errors.Is(err, ErrReportNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "report not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    var req CreateReportRequest
    updateReport(&report,&req)
    if err := ctx.ShouldBind(&req); err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request")
        return
    }
   
    if err := h.Service.UpdateReport(report); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    response := ToReportResponse(report)
    ctx.JSON(http.StatusOK, response)
}


// DeleteReport godoc
// @Summary      Delete report
// @Description  Delete a report by its ID
// @Tags         report
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "Report ID"
// @Success      204 "No Content"
// @Failure      400 {object} basics.APIError "Invalid UUID format"
// @Failure      404 {object} basics.APIError "Report not found"
// @Failure      500 {object} basics.APIError "Internal server error"
// @Router       /report/api/v1/{id} [delete]
func (h *Handler) DeleteReport(ctx *gin.Context) {
    reportId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        basics.ErrorResponse(ctx, http.StatusBadRequest, "Invalid UUID format")
        return
    }

    report, err := h.Service.Repository.FindReportById(reportId)
    if errors.Is(err, ErrReportNotFound) {
        basics.ErrorResponse(ctx, http.StatusNotFound, "report not found")
        return
    } else if err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    if err := h.Service.Repository.DeleteReport(report); err != nil {
        basics.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
        return
    }

    ctx.Status(http.StatusNoContent) // 204 No Content
}

func updateReport(report *Report, req *CreateReportRequest) error {
	reportVal := reflect.ValueOf(report).Elem()
	reqVal := reflect.ValueOf(req).Elem()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		if !fieldVal.IsNil() {
			reportField := reportVal.FieldByName(reqVal.Type().Field(i).Name)
			if reportField.IsValid() && reportField.CanSet() {
				reportField.Set(reflect.Indirect(fieldVal))
			}
		}
	}

	return nil
}
