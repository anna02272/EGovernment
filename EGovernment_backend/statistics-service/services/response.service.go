package services

import (
	"statistics-service/domain"
)

type ResponseService interface {
	Create(request *domain.Response) (error, bool)
	GetAll() ([]*domain.Response, error)
	GetById(id string) (*domain.Response, error)
}
