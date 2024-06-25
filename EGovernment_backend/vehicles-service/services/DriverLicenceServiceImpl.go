package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"vehicles-service/domain"
)

type DriverLicenceServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewDriverLicenceServiceImpl(collection *mongo.Collection, ctx context.Context) DriverLicenceService {
	return &DriverLicenceServiceImpl{collection, ctx}
}

func (s *DriverLicenceServiceImpl) InsertDriverLicence(licence *domain.DriverLicenceCreate) (*domain.DriverLicence, string, error) {
	var licenceToInsert domain.DriverLicence
	licenceToInsert.ID = primitive.NewObjectID()
	licenceToInsert.VehicleDriver = licence.VehicleDriver
	licenceToInsert.LicenceNumber = uuid.NewString()
	licenceToInsert.LocationLicenced = licence.LocationLicenced
	licenceToInsert.Categories = licence.Categories
	result, err := s.collection.InsertOne(context.Background(), licenceToInsert)
	if err != nil {
		return nil, "", err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, "", errors.New("failed to get inserted ID")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return &licenceToInsert, insertedID.Hex(), nil
}

func (s *DriverLicenceServiceImpl) GetDriverLicenceById(driverLicenceNumber string, ctx context.Context) (*domain.DriverLicence, error) {
	var driverLicence domain.DriverLicence
	filter := bson.M{"licence_number": driverLicenceNumber}

	err := s.collection.FindOne(ctx, filter).Decode(&driverLicence)
	if err != nil {
		return nil, err
	}

	return &driverLicence, nil
}

func (s *DriverLicenceServiceImpl) GetDriverLicenceByDriver(driver string, ctx context.Context) (*domain.DriverLicence, error) {
	var driverLicence domain.DriverLicence
	filter := bson.M{"vehicle_driver": driver}

	err := s.collection.FindOne(ctx, filter).Decode(&driverLicence)
	if err != nil {
		return nil, err
	}

	return &driverLicence, nil
}

func (s *DriverLicenceServiceImpl) GetAllDriverLicences() ([]*domain.DriverLicence, error) {
	var driverLicences []*domain.DriverLicence

	filter := bson.D{}

	cursor, err := s.collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var driverLicence domain.DriverLicence
		if err := cursor.Decode(&driverLicence); err != nil {
			return nil, err
		}
		driverLicences = append(driverLicences, &driverLicence)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return driverLicences, nil
}
