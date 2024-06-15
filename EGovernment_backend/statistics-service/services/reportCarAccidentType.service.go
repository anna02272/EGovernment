package services

import (
	"statistics-service/domain"
)

type ReportCarAccidentTypeService interface {
	Create(report *domain.ReportCarAccidentType) (error, bool)
	GetAll() ([]*domain.ReportCarAccidentType, error)
	GetById(id string) (*domain.ReportCarAccidentType, error)
	GetAllByCarAccidentType(carAccidentType domain.CarAccidentType) ([]*domain.ReportCarAccidentType, error)
	GetAllByCarAccidentTypeAndYear(carAccidentType domain.CarAccidentType, year int) ([]*domain.ReportCarAccidentType, error)
}
