package services

import (
	"vehicles-service/domain"
)

type VehicleService interface {
	InsertVehicleDriver(driver *domain.VehicleDriverCreate) (*domain.VehicleDriver, string, error)
}
