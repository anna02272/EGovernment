package services

import (
	"context"
	"vehicles-service/domain"
)

type DriverLicenceService interface {
	InsertDriverLicence(licence *domain.DriverLicenceCreate) (*domain.DriverLicence, string, error)
	GetAllDriverLicences() ([]*domain.DriverLicence, error)
	GetDriverLicenceById(driverLicenceNumber string, ctx context.Context) (*domain.DriverLicence, error)
	GetDriverLicenceByDriver(driver string, ctx context.Context) (*domain.DriverLicence, error)
}
