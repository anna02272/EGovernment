package services

import (
	"context"
	"police-service/domain"
)

type CarAccidentService interface {
	InsertCarAccident(carAccident *domain.CarAccidentCreate, policemanID string) (*domain.CarAccident, string, error)
	GetAllCarAccident() ([]*domain.CarAccident, error)
	GetCarAccidentById(carAccidentId string, ctx context.Context) (*domain.CarAccident, error)
	GetAllCarAccidentsByType(carAccidentType domain.CarAccidentType) ([]*domain.CarAccident, error)
	GetAllCarAccidentsByTypeAndYear(carAccidentType domain.CarAccidentType, year int) ([]*domain.CarAccident, error)
	GetAllCarAccidentsByDegree(degreeOfAccident domain.DegreeOfAccident) ([]*domain.CarAccident, error)
	GetAllCarAccidentsByDegreeAndYear(degreeOfAccident domain.DegreeOfAccident, year int) ([]*domain.CarAccident, error)
	GetAllCarAccidentsByPolicemanID(policemanID string) ([]*domain.CarAccident, error)
	GetAllCarAccidentsByDriver(driverEmail string) ([]*domain.CarAccident, error)
}
