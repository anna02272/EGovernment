package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatisticsServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewStatisticsServiceImpl(collection *mongo.Collection, ctx context.Context) StatisticsService {
	return &StatisticsServiceImpl{collection, ctx}
}
