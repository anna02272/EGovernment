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

func (s *VehicleServiceImpl) InsertVehicleDriver(driver *domain.VehicleDriverCreate) (*domain.VehicleDriver, string, error) {

	var vehicleDriver domain.VehicleDriver
	vehicleDriver.ID = primitive.NewObjectID()
	vehicleDriver.IdentificationNumber = driver.IdentificationNumber
	vehicleDriver.Name = driver.Name
	vehicleDriver.LastName = driver.LastName
	vehicleDriver.DateOfBirth = driver.DateOfBirth.Date()
	vehicleDriver.HasDelict = false
	vehicleDriver.Gender = driver.Gender
	vehicleDriver.NumberOfPenaltyPoints = 0

	result, err := s.collection.InsertOne(context.Background(), driver)
	if err != nil {
		return nil, "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, "", errors.New("failed to get inserted ID")
	}

	insertedID = result.InsertedID.(primitive.ObjectID)

	return &vehicleDriver, insertedID.Hex(), nil
}
