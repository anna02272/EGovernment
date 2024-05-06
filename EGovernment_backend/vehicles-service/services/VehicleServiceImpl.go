package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type VehicleServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewVehicleServiceImpl(collection *mongo.Collection, ctx context.Context) VehicleService {
	return &VehicleServiceImpl{collection, ctx}
}
