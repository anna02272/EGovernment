package services

import (
	"statistics-service/domain"
)

type RequestService interface {
	Create(request *domain.Request) (error, bool)
	GetAll() ([]*domain.Request, error)
	GetById(id string) (*domain.Request, error)
}
