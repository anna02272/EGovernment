package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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
	vehicleToInsert.Category = vehicle.Category
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

func (s *VehicleServiceImpl) GetAllVehicles() ([]*domain.Vehicle, error) {
	var vehicles []*domain.Vehicle

	filter := bson.D{}

	cursor, err := s.collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var vehicle domain.Vehicle
		if err := cursor.Decode(&vehicle); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &vehicle)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (s *VehicleServiceImpl) GetAllRegisteredVehicles() ([]*domain.Vehicle, error) {
	var vehicles []*domain.Vehicle

	cutoffDate := time.Now().AddDate(-1, 0, 0)

	filter := bson.M{
		"registration_date": bson.M{
			"$gte": cutoffDate,
		},
	}

	cursor, err := s.collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var vehicle domain.Vehicle
		if err := cursor.Decode(&vehicle); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &vehicle)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (s *VehicleServiceImpl) GetAllVehiclesByCategoryAndYear(category domain.Category, year int) ([]*domain.Vehicle, error) {
	var vehicles []*domain.Vehicle
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := startOfYear.AddDate(1, 0, 0)

	filter := bson.M{
		"category": category,
		"registration_date": bson.M{
			"$gte": startOfYear,
			"$lt":  endOfYear,
		},
	}

	cursor, err := s.collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var vehicle domain.Vehicle
		if err := cursor.Decode(&vehicle); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, &vehicle)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (s *VehicleServiceImpl) GetVehicleByID(registrationPlate string, ctx context.Context) (*domain.Vehicle, error) {
	var vehicle domain.Vehicle
	filter := bson.M{"_registration_plate": registrationPlate}

	err := s.collection.FindOne(ctx, filter).Decode(&vehicle)
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}
