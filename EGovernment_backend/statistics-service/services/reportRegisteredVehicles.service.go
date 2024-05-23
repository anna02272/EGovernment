package services

import (
	"statistics-service/domain"
)

type ReportRegisteredVehiclesService interface {
	Create(report *domain.ReportRegisteredVehicle) (error, bool)
	GetAll() ([]*domain.ReportRegisteredVehicle, error)
	GetById(id string) (*domain.ReportRegisteredVehicle, error)
	GetAllByCategory(category domain.Category) ([]*domain.ReportRegisteredVehicle, error)
	GetAllByCategoryAndYear(category domain.Category, year int) ([]*domain.ReportRegisteredVehicle, error)
}
