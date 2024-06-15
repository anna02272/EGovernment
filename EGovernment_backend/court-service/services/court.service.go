package services

import (
	"court-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourtService interface {
	CreateCourt(court *domain.Court) (*domain.Court, error)
	GetCourtByID(id primitive.ObjectID) (*domain.Court, error)
}
