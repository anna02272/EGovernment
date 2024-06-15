package handlers

import (
	"court-service/domain"
	"court-service/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScheduleHandler struct {
	service services.ScheduleService
	DB      *mongo.Collection
}

func NewScheduleHandler(service services.ScheduleService, db *mongo.Collection) ScheduleHandler {
	return ScheduleHandler{
		service: service,
		DB:      db,
	}
}

func (sh *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var schedule domain.Schedule

	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	createdSchedule, err := sh.service.CreateSchedule(&schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	c.JSON(http.StatusCreated, createdSchedule)
}
func (sh *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid ID format"})
		return
	}

	schedule, err := sh.service.GetScheduleByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get schedule"})
		return
	}

	if schedule == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}
