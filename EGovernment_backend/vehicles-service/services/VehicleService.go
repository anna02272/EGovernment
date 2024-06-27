package services

import (
	"context"
	"vehicles-service/domain"
)

type VehicleService interface {
	InsertVehicle(vehicle *domain.VehicleCreate) (*domain.Vehicle, string, error)
	GetAllVehicles() ([]*domain.Vehicle, error)
	GetAllRegisteredVehicles() ([]*domain.Vehicle, error)
	GetNumberOfRegisteredVehiclesByCategory(category domain.Category) (int64, error)
	GetVehicleByID(registrationPlate string, ctx context.Context) (*domain.Vehicle, error)
	GetAllVehiclesByCategoryAndYear(category domain.Category, year int) ([]*domain.Vehicle, error)
	GetAllRegisteredVehiclesByCategory(category domain.Category) ([]*domain.Vehicle, error)
}
