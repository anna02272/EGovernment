package handlers

import (
	"court-service/domain"
	"court-service/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type HearingHandler struct {
	service services.HearingService
	DB      *mongo.Collection
}

func NewHearingHandler(service services.HearingService, db *mongo.Collection) HearingHandler {
	return HearingHandler{
		service: service,
		DB:      db,
	}
}

func (hh *HearingHandler) CreateHearing(c *gin.Context) {
	var hearing domain.Hearing

	if err := c.ShouldBindJSON(&hearing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	createdHearing, err := hh.service.CreateHearing(&hearing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hearing"})
		return
	}

	c.JSON(http.StatusCreated, createdHearing)
}
func (hh *HearingHandler) GetHearingByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid ID format"})
		return
	}

	hearing, err := hh.service.GetHearingByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get hearing"})
		return
	}

	if hearing == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Hearing not found"})
		return
	}

	c.JSON(http.StatusOK, hearing)
}
