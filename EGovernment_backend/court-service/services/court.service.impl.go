package services

import (
	"context"
	"court-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourtServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCourtServiceImpl(collection *mongo.Collection, ctx context.Context) CourtService {
	return &CourtServiceImpl{collection, ctx}
}

func (cs *CourtServiceImpl) CreateCourt(court *domain.Court) (*domain.Court, error) {
	result, err := cs.collection.InsertOne(cs.ctx, court)
	if err != nil {
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to get inserted ID")
	}

	court.ID = insertedID
	return court, nil
}

func (cs *CourtServiceImpl) GetCourtByID(id primitive.ObjectID) (*domain.Court, error) {
	var court domain.Court
	err := cs.collection.FindOne(cs.ctx, bson.M{"_id": id}).Decode(&court)
	if err != nil {
		return nil, err
	}
	return &court, nil
}
