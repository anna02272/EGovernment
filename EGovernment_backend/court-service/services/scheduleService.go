package services

import (
	"court-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScheduleService interface {
	CreateSchedule(schedule *domain.Schedule) (*domain.Schedule, error)
	GetScheduleByID(id primitive.ObjectID) (*domain.Schedule, error)
	GetScheduleByHearingID(hearingID primitive.ObjectID) (*domain.Schedule, error) // New method

}
