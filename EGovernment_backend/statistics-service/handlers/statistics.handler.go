package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"statistics-service/services"
)

type StatisticsHandler struct {
	service services.StatisticsService
	DB      *mongo.Collection
}

func NewStatisticsHandler(service services.StatisticsService, db *mongo.Collection) StatisticsHandler {
	return StatisticsHandler{
		service: service,
		DB:      db,
	}

}
