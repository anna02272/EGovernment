package handlers

import (
	"court-service/services"
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
