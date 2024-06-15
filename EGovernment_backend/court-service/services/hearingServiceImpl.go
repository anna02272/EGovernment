package services

import (
	"context"
	"court-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HearingServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewHearingServiceImpl(collection *mongo.Collection, ctx context.Context) HearingService {
	return &HearingServiceImpl{collection, ctx}
}

func (hs *HearingServiceImpl) CreateHearing(hearing *domain.Hearing) (*domain.Hearing, error) {
	result, err := hs.collection.InsertOne(hs.ctx, hearing)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}

	hearing.ID = insertedID
	return hearing, nil
}

func (hs *HearingServiceImpl) GetHearingByID(id primitive.ObjectID) (*domain.Hearing, error) {
	var hearing domain.Hearing
	err := hs.collection.FindOne(hs.ctx, bson.M{"_id": id}).Decode(&hearing)
	if err != nil {
		return nil, err
	}
	return &hearing, nil
}
func (hs *HearingServiceImpl) GetSubjectById(id primitive.ObjectID) (*domain.Subject, error) {
	var subject domain.Subject
	err := hs.collection.FindOne(hs.ctx, bson.M{"_id": id}).Decode(&subject)
	if err != nil {
		return nil, err
	}
	return &subject, nil
}
