package services

import (
	"context"
	"vehicles-service/domain"
)

type VehicleService interface {
	InsertVehicle(vehicle *domain.VehicleCreate) (*domain.Vehicle, string, error)
	GetAllVehicles() ([]*domain.Vehicle, error)
	GetVehicleByID(registrationPlate string, ctx context.Context) (*domain.Vehicle, error)
}
