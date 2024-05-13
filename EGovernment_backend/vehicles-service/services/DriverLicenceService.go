package services

import (
	"context"
	"vehicles-service/domain"
)

type DriverLicenceService interface {
	InsertDriverLicence(licence *domain.DriverLicenceCreate) (*domain.DriverLicence, string, error)
	GetDriverLicenceById(driverLicenceNumber string, ctx context.Context) (*domain.DriverLicence, error)
}
