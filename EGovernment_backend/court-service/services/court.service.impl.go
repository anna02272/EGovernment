package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourtServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCourtServiceImpl(collection *mongo.Collection, ctx context.Context) CourtService {
	return &CourtServiceImpl{collection, ctx}
}
