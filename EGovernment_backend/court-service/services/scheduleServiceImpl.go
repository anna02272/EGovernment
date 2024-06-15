package services

import (
	"context"
	"court-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScheduleServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewScheduleServiceImpl(collection *mongo.Collection, ctx context.Context) ScheduleService {
	return &ScheduleServiceImpl{collection, ctx}
}

func (ss *ScheduleServiceImpl) CreateSchedule(schedule *domain.Schedule) (*domain.Schedule, error) {
	result, err := ss.collection.InsertOne(ss.ctx, schedule)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}

	schedule.ID = insertedID
	return schedule, nil
}

func (ss *ScheduleServiceImpl) GetScheduleByID(id primitive.ObjectID) (*domain.Schedule, error) {
	var schedule domain.Schedule
	err := ss.collection.FindOne(ss.ctx, bson.M{"_id": id}).Decode(&schedule)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}
