package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"vehicles-service/services"
)

type VehicleHandler struct {
	service services.VehicleService
	DB      *mongo.Collection
}

func NewVehicleHandler(service services.VehicleService, db *mongo.Collection) VehicleHandler {
	return VehicleHandler{
		service: service,
		DB:      db,
	}

}
