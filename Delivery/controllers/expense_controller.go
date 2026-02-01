package controllers

import (
	"net/http"
	"strconv"
	"time"

	Domain "ShopOps/Domain"
	Usecases "ShopOps/Usecases"
	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	expenseUC Usecases.ExpenseUseCase
}

func NewExpenseController(expenseUC Usecases.ExpenseUseCase) *ExpenseController {
	return &ExpenseController{expenseUC: expenseUC}
}

// CreateExpense records a new expense
func (c *ExpenseController) CreateExpense(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req Domain.CreateExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense, err := c.expenseUC.CreateExpense(businessID, userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, expense)
}

// GetExpenses lists all expenses with filtering
func (c *ExpenseController) GetExpenses(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	// Parse filters
	filters := Domain.ExpenseFilters{}

	// Date filters
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			filters.StartDate = &startDate
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			filters.EndDate = &endDate
		}
	}

	// Category filter
	if category := ctx.Query("category"); category != "" {
		expenseCategory := Domain.ExpenseCategory(category)
		filters.Category = &expenseCategory
	}

	// Status filter
	if status := ctx.Query("status"); status != "" {
		expenseStatus := Domain.ExpenseStatus(status)
		filters.Status = &expenseStatus
	}

	// Pagination
	if limitStr := ctx.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filters.Limit = limit
		}
	}

	if offsetStr := ctx.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil && offset >= 0 {
			filters.Offset = offset
		}
	}

	expenses, err := c.expenseUC.GetExpenses(businessID, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expenses)
}

// GetExpense gets expense details
func (c *ExpenseController) GetExpense(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	expenseID := ctx.Param("expenseId")
	if expenseID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Expense ID is required"})
		return
	}

	expense, err := c.expenseUC.GetExpenseByID(expenseID, businessID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

// UpdateExpense updates expense (before sync)
func (c *ExpenseController) UpdateExpense(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	expenseID := ctx.Param("expenseId")
	if expenseID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Expense ID is required"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req Domain.CreateExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense, err := c.expenseUC.UpdateExpense(expenseID, businessID, userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expense)
}

// VoidExpense voids an expense
func (c *ExpenseController) VoidExpense(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	expenseID := ctx.Param("expenseId")
	if expenseID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Expense ID is required"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if err := c.expenseUC.VoidExpense(expenseID, businessID, userID.(string)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Expense voided successfully"})
}

// GetExpenseSummary gets expense summary by category
func (c *ExpenseController) GetExpenseSummary(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	period := ctx.DefaultQuery("period", "month")

	summary, err := c.expenseUC.GetExpenseSummary(businessID, period)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, summary)
}

// GetExpenseCategories lists available expense categories
func (c *ExpenseController) GetExpenseCategories(ctx *gin.Context) {
	categories := c.expenseUC.GetExpenseCategories()
	ctx.JSON(http.StatusOK, categories)
}
