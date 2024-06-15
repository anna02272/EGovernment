package services

import (
	"court-service/domain"
)

type CitizenService interface {
	InsertCitizen(driver *domain.Citizen) (*domain.Citizen, string, error)
	GetAllCitizens() ([]*domain.Citizen, error)
	GetCitizenByID(jmbg string) (*domain.Citizen, error)
}
