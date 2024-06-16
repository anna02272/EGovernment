package handlers

import (
	"court-service/domain"
	"court-service/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourtHandler struct {
	service services.CourtService
	DB      *mongo.Collection
}

func NewCourtHandler(service services.CourtService, db *mongo.Collection) CourtHandler {
	return CourtHandler{
		service: service,
		DB:      db,
	}
}

func (ch *CourtHandler) CreateCourt(c *gin.Context) {
	var court domain.Court

	if err := c.ShouldBindJSON(&court); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	createdCourt, err := ch.service.CreateCourt(&court)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create court"})
		return
	}

	c.JSON(http.StatusCreated, createdCourt)
}
func (ch *CourtHandler) GetCourtByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid ID format"})
		return
	}

	court, err := ch.service.GetCourtByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get court"})
		return
	}

	if court == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Court not found"})
		return
	}

	c.JSON(http.StatusOK, court)
}
func (ch *CourtHandler) GetAllCourts(c *gin.Context) {
	courts, err := ch.service.GetAllCourts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get courts"})
		return
	}

	c.JSON(http.StatusOK, courts)
}
