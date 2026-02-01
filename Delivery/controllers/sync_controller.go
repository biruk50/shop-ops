package controllers

import (
	"net/http"

	Domain "ShopOps/Domain"
	Usecases "ShopOps/Usecases"
	"github.com/gin-gonic/gin"
)

type SyncController struct {
	syncUC Usecases.SyncUseCase
}

func NewSyncController(syncUC Usecases.SyncUseCase) *SyncController {
	return &SyncController{syncUC: syncUC}
}

// ProcessBatch handles batch sync requests
func (c *SyncController) ProcessBatch(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	_, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var batch Domain.SyncBatch
	if err := ctx.ShouldBindJSON(&batch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set business ID from URL parameter
	batch.BusinessID = businessID

	response, err := c.syncUC.ProcessBatch(batch)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetSyncStatus returns sync status for a business
func (c *SyncController) GetSyncStatus(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	_, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	status, err := c.syncUC.GetSyncStatus(businessID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, status)
}

// GetLastSync returns last sync time for a device
func (c *SyncController) GetLastSync(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	if businessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Business ID is required"})
		return
	}

	deviceID := ctx.Query("device_id")
	if deviceID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Device ID is required"})
		return
	}

	lastSync, err := c.syncUC.GetLastSync(businessID, deviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if lastSync == nil {
		ctx.JSON(http.StatusOK, gin.H{"last_sync": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"last_sync": lastSync})
}
