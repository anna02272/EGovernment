package services

import (
	"context"
	"vehicles-service/domain"
)

type VehicleDriverService interface {
	InsertVehicleDriver(driver *domain.VehicleDriverCreate) (*domain.VehicleDriver, string, error)
	GetAllVehicleDrivers() ([]*domain.VehicleDriver, error)
	GetVehicleDriverByID(IdentificationNumber string, ctx context.Context) (*domain.VehicleDriver, error)
}
