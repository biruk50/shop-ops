package controllers

import (
	"net/http"

	Domain "ShopOps/Domain"
	Usecases "ShopOps/Usecases"
	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	businessUC Usecases.BusinessUseCase
}

func NewBusinessController(businessUC Usecases.BusinessUseCase) *BusinessController {
	return &BusinessController{businessUC: businessUC}
}

// CreateBusiness creates a new business
func (c *BusinessController) CreateBusiness(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req Domain.CreateBusinessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	business, err := c.businessUC.CreateBusiness(userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

// GetBusinesses lists user's businesses
func (c *BusinessController) GetBusinesses(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	businesses, err := c.businessUC.GetUserBusinesses(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, businesses)
}

// GetBusiness gets business details
func (c *BusinessController) GetBusiness(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	business, err := c.businessUC.GetBusinessByID(businessID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, business)
}

// UpdateBusiness updates business settings
func (c *BusinessController) UpdateBusiness(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	var req Domain.UpdateBusinessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	business, err := c.businessUC.UpdateBusiness(businessID, userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, business)
}
