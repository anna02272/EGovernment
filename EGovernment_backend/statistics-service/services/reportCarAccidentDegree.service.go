package services

import (
	"statistics-service/domain"
)

type ReportCarAccidentDegreeService interface {
	Create(report *domain.ReportCarAccidentDegree) (error, bool)
	GetAll() ([]*domain.ReportCarAccidentDegree, error)
	GetById(id string) (*domain.ReportCarAccidentDegree, error)
	GetAllByCarAccidentDegree(degree domain.DegreeOfAccident) ([]*domain.ReportCarAccidentDegree, error)
	GetAllByCarAccidentDegreeAndYear(degree domain.DegreeOfAccident, year int) ([]*domain.ReportCarAccidentDegree, error)
}
