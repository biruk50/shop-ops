package controllers

import (
	"net/http"
	"strconv"

	Domain "ShopOps/Domain"
	Usecases "ShopOps/Usecases"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	inventoryUC Usecases.InventoryUseCase
}

func NewInventoryController(inventoryUC Usecases.InventoryUseCase) *InventoryController {
	return &InventoryController{inventoryUC: inventoryUC}
}

// CreateProduct adds new product
func (c *InventoryController) CreateProduct(ctx *gin.Context) {
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

	var req Domain.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.inventoryUC.CreateProduct(businessID, userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// GetProducts lists all products
func (c *InventoryController) GetProducts(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	// Parse filters
	filters := Domain.ProductFilters{}

	// Category filter
	if category := ctx.Query("category"); category != "" {
		filters.Category = &category
	}

	// Status filter
	if status := ctx.Query("status"); status != "" {
		productStatus := Domain.ProductStatus(status)
		filters.Status = &productStatus
	}

	// Low stock filter
	if lowStock := ctx.Query("low_stock"); lowStock != "" {
		if lowStock == "true" {
			trueVal := true
			filters.LowStock = &trueVal
		}
	}

	// Search filter
	if search := ctx.Query("search"); search != "" {
		filters.Search = &search
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

	products, err := c.inventoryUC.GetProducts(businessID, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetProduct gets product details
func (c *InventoryController) GetProduct(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	productID := ctx.Param("productId")
	if productID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	product, err := c.inventoryUC.GetProductByID(productID, businessID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// UpdateProduct updates product
func (c *InventoryController) UpdateProduct(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	productID := ctx.Param("productId")
	if productID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req Domain.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.inventoryUC.UpdateProduct(productID, businessID, userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct deletes product
func (c *InventoryController) DeleteProduct(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	productID := ctx.Param("productId")
	if productID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if err := c.inventoryUC.DeleteProduct(productID, businessID, userID.(string)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// AdjustStock manually adjusts stock
func (c *InventoryController) AdjustStock(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	productID := ctx.Param("productId")
	if productID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req Domain.AdjustStockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.inventoryUC.AdjustStock(productID, businessID, userID.(string), req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Stock adjusted successfully"})
}

// GetLowStock gets products below threshold
func (c *InventoryController) GetLowStock(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	threshold := 0.0
	if thresholdStr := ctx.Query("threshold"); thresholdStr != "" {
		if t, err := strconv.ParseFloat(thresholdStr, 64); err == nil {
			threshold = t
		}
	}

	products, err := c.inventoryUC.GetLowStock(businessID, threshold)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetStockHistory gets stock movement history
func (c *InventoryController) GetStockHistory(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	productID := ctx.Param("productId")
	if productID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	limit := 50
	if limitStr := ctx.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	history, err := c.inventoryUC.GetStockHistory(productID, businessID, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, history)
}
