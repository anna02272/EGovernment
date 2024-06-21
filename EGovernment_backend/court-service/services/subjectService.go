package services

import (
	"court-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubjectService interface {
	CreateSubject(subject *domain.Subject) (*domain.Subject, error)
	GetSubjectByID(id primitive.ObjectID) (*domain.Subject, error)
	GetAllSubjects() ([]domain.Subject, error)
	UpdateSubjectStatus(id primitive.ObjectID, status domain.Status) error
	UpdateSubjectJudgment(id primitive.ObjectID, judgment string) error
	UpdateSubjectCompromis(id primitive.ObjectID, compromis string) error
}
