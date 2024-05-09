package services

import (
	"vehicles-service/domain"
)

type VehicleService interface {
	InsertVehicle(vehicle *domain.VehicleCreate) (*domain.Vehicle, string, error)
}
