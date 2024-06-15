package services

import (
	"court-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HearingService interface {
	CreateHearing(hearing *domain.Hearing) (*domain.Hearing, error)
	GetHearingByID(id primitive.ObjectID) (*domain.Hearing, error)
}
