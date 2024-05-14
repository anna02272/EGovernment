package handlers

import (
	"court-service/domain"
	"court-service/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type SubjectHandler struct {
	service services.SubjectService
	DB      *mongo.Collection
}

func NewSubjectHandler(service services.SubjectService, db *mongo.Collection) SubjectHandler {
	return SubjectHandler{
		service: service,
		DB:      db,
	}
}

func (sh *SubjectHandler) CreateSubject(c *gin.Context) {
	var subject *domain.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	createdSubject, err := sh.service.CreateSubject(subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject"})
		return
	}

	c.JSON(http.StatusCreated, createdSubject)
}
