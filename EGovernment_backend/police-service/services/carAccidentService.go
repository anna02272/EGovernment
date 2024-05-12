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
	GetAllCarAccidentsByDegree(degreeOfAccident domain.DegreeOfAccident) ([]*domain.CarAccident, error)
}
