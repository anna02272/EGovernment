package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"vehicles-service/domain"
)

type VehicleDriverServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewVehicleDriverServiceImpl(collection *mongo.Collection, ctx context.Context) VehicleDriverService {
	return &VehicleDriverServiceImpl{collection, ctx}
}

func (s *VehicleDriverServiceImpl) InsertVehicleDriver(driver *domain.VehicleDriverCreate) (*domain.VehicleDriver, string, error) {

	var vehicleDriver domain.VehicleDriver
	vehicleDriver.ID = primitive.NewObjectID()
	vehicleDriver.IdentificationNumber = driver.IdentificationNumber
	vehicleDriver.Name = driver.Name
	vehicleDriver.LastName = driver.LastName
	vehicleDriver.DateOfBirth = driver.DateOfBirth
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

func (s *VehicleDriverServiceImpl) GetAllVehicleDrivers() ([]*domain.VehicleDriver, error) {
	var vehicleDrivers []*domain.VehicleDriver
	filter := bson.D{}

	cursor, err := s.collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var vehicleDriver domain.VehicleDriver
		if err := cursor.Decode(&vehicleDriver); err != nil {
			return nil, err
		}
		vehicleDrivers = append(vehicleDrivers, &vehicleDriver)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return vehicleDrivers, nil
}

func (s *VehicleDriverServiceImpl) GetVehicleDriverByID(identificationNumber string, ctx context.Context) (*domain.VehicleDriver, error) {
	var vehicleDriver domain.VehicleDriver
	filter := bson.M{"identification_number": identificationNumber}

	err := s.collection.FindOne(ctx, filter).Decode(&vehicleDriver)
	if err != nil {
		return nil, err
	}

	return &vehicleDriver, nil
}

func (s *VehicleDriverServiceImpl) UpdatePenaltyPoints(identificationNumber string, points int64, ctx context.Context) error {
	filter := bson.M{"identification_number": identificationNumber}
	update := bson.M{"$inc": bson.M{"number_of_penalty_points": points}}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}
