package controllers

import (
	"net/http"
	"strconv"
	"time"

	Domain "ShopOps/Domain"
	Usecases "ShopOps/Usecases"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportUC Usecases.ReportUseCase
}

func NewReportController(reportUC Usecases.ReportUseCase) *ReportController {
	return &ReportController{reportUC: reportUC}
}

// GetDashboard gets dashboard overview
func (c *ReportController) GetDashboard(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	data, err := c.reportUC.GetDashboardData(businessID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetSalesReport gets sales report
func (c *ReportController) GetSalesReport(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	var req Domain.ReportRequest
	req.BusinessID = businessID
	req.Type = Domain.ReportTypeSales

	// Parse period
	if period := ctx.Query("period"); period != "" {
		req.Period = Domain.PeriodType(period)
	} else {
		req.Period = Domain.PeriodTypeMonthly
	}

	// Parse dates
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			req.StartDate = &startDate
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			req.EndDate = &endDate
		}
	}

	report, err := c.reportUC.GenerateReport(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}

// GetExpensesReport gets expenses report
func (c *ReportController) GetExpensesReport(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	var req Domain.ReportRequest
	req.BusinessID = businessID
	req.Type = Domain.ReportTypeExpenses

	// Parse period
	if period := ctx.Query("period"); period != "" {
		req.Period = Domain.PeriodType(period)
	} else {
		req.Period = Domain.PeriodTypeMonthly
	}

	// Parse dates
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			req.StartDate = &startDate
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			req.EndDate = &endDate
		}
	}

	// Parse category
	if category := ctx.Query("category"); category != "" {
		req.Category = &category
	}

	report, err := c.reportUC.GenerateReport(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}

// GetProfitReport gets profit report
func (c *ReportController) GetProfitReport(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	var req Domain.ReportRequest
	req.BusinessID = businessID
	req.Type = Domain.ReportTypeProfit

	// Parse period
	if period := ctx.Query("period"); period != "" {
		req.Period = Domain.PeriodType(period)
	} else {
		req.Period = Domain.PeriodTypeMonthly
	}

	// Parse dates
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			req.StartDate = &startDate
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			req.EndDate = &endDate
		}
	}

	report, err := c.reportUC.GenerateReport(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}

// GetInventoryReport gets inventory report
// GetInventoryReport gets inventory report
func (c *ReportController) GetInventoryReport(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	var req Domain.ReportRequest
	req.BusinessID = businessID
	req.Type = Domain.ReportTypeInventory

	report, err := c.reportUC.GenerateReport(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}

// ExportReport generates CSV export
func (c *ReportController) ExportReport(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	var req Domain.ReportRequest
	req.BusinessID = businessID

	// Parse report type
	if reportType := ctx.Query("type"); reportType != "" {
		req.Type = Domain.ReportType(reportType)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Report type is required"})
		return
	}

	// Parse period
	if period := ctx.Query("period"); period != "" {
		req.Period = Domain.PeriodType(period)
	} else {
		req.Period = Domain.PeriodTypeMonthly
	}

	// Parse dates
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			req.StartDate = &startDate
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			req.EndDate = &endDate
		}
	}

	// Set format to CSV
	format := "csv"
	req.Format = &format

	data, filename, err := c.reportUC.ExportReport(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Data(http.StatusOK, "text/csv", data)
}

// GetProfitSummary gets profit summary for period
func (c *ReportController) GetProfitSummary(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	period := Domain.PeriodType(ctx.DefaultQuery("period", "monthly"))

	var startDate, endDate *time.Time
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if sd, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &sd
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if ed, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &ed
		}
	}

	report, err := c.reportUC.GetProfitSummary(businessID, period, startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}

// GetProfitTrends gets profit trends over time
func (c *ReportController) GetProfitTrends(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	period := Domain.PeriodType(ctx.DefaultQuery("period", "weekly"))
	weeks := 12
	if weeksStr := ctx.Query("weeks"); weeksStr != "" {
		if w, err := strconv.Atoi(weeksStr); err == nil && w > 0 {
			weeks = w
		}
	}

	trends, err := c.reportUC.GetProfitTrends(businessID, period, weeks)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, trends)
}
