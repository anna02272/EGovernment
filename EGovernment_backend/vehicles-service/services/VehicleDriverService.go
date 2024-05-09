package services

import "vehicles-service/domain"

type VehicleDriverService interface {
	InsertVehicleDriver(driver *domain.VehicleDriverCreate) (*domain.VehicleDriver, string, error)
}
