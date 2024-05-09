package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"vehicles-service/domain"
)

type VehicleServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewVehicleServiceImpl(collection *mongo.Collection, ctx context.Context) VehicleService {
	return &VehicleServiceImpl{collection, ctx}
}

func (s *VehicleServiceImpl) InsertVehicle(vehicle *domain.VehicleCreate) (*domain.Vehicle, string, error) {
	var vehicleToInsert domain.Vehicle
	vehicleToInsert.ID = primitive.NewObjectID()
	vehicleToInsert.RegistrationPlate = vehicle.RegistrationPlate
	vehicleToInsert.VehicleModel = vehicle.VehicleModel
	vehicleToInsert.VehicleOwner = vehicle.VehicleOwner
	vehicleToInsert.RegistrationDate = vehicle.RegistrationDate
	result, err := s.collection.InsertOne(context.Background(), vehicle)
	if err != nil {
		return nil, "", err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, "", errors.New("failed to get inserted ID")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return &vehicleToInsert, insertedID.Hex(), nil
}
