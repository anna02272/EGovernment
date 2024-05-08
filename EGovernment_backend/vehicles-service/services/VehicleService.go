package services

import (
	"context"
	"vehicles-service/domain"
)

type VehicleService interface {
	InsertVehicleDriver(driver *domain.VehicleDriverCreate, ctx context.Context) (*domain.VehicleDriver, string, error)
}
