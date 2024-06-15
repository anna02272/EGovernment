package services

import (
	"court-service/domain"
)

type SubjectService interface {
	CreateSubject(subject *domain.Subject) (*domain.Subject, error)
}
