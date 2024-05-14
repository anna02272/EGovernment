package services

import (
	"statistics-service/domain"
)

type ReportDelicTypeService interface {
	Create(report *domain.ReportDelict) (error, bool)
	GetAll() ([]*domain.ReportDelict, error)
	GetById(id string) (*domain.ReportDelict, error)
	GetAllByDelictType(delictType domain.DelictType) ([]*domain.ReportDelict, error)
}
