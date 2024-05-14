package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"police-service/domain"
	"time"
)

type CarAccidentServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCarAccidentServiceImpl(collection *mongo.Collection, ctx context.Context) CarAccidentService {
	return &CarAccidentServiceImpl{collection, ctx}
}

func (c CarAccidentServiceImpl) InsertCarAccident(carAccident *domain.CarAccidentCreate, policemanID string) (*domain.CarAccident, string, error) {
	carAccident.PolicemanID = policemanID
	currentTime := time.Now()
	dateTime := primitive.NewDateTimeFromTime(currentTime)
	carAccident.Date = dateTime
	var trafficCarAccident domain.CarAccident
	trafficCarAccident.ID = primitive.NewObjectID()
	trafficCarAccident.PolicemanID = policemanID
	trafficCarAccident.DriverIdentificationNumber1 = carAccident.DriverIdentificationNumber1
	trafficCarAccident.DriverIdentificationNumber2 = carAccident.DriverIdentificationNumber2
	trafficCarAccident.VehicleLicenceNumber1 = carAccident.VehicleLicenceNumber1
	trafficCarAccident.VehicleLicenceNumber2 = carAccident.VehicleLicenceNumber2
	trafficCarAccident.DriverEmail = carAccident.DriverEmail
	trafficCarAccident.Date = dateTime
	trafficCarAccident.Location = carAccident.Location
	trafficCarAccident.Description = carAccident.Description
	trafficCarAccident.CarAccidentType = carAccident.CarAccidentType
	trafficCarAccident.DegreeOfAccident = carAccident.DegreeOfAccident
	trafficCarAccident.NumberOfPenaltyPoints = carAccident.NumberOfPenaltyPoints

	result, err := c.collection.InsertOne(context.Background(), carAccident)
	if err != nil {
		return nil, "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, "", errors.New("failed to get inserted ID")
	}

	insertedID = result.InsertedID.(primitive.ObjectID)

	return &trafficCarAccident, insertedID.Hex(), nil
}

func (c CarAccidentServiceImpl) GetAllCarAccident() ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.D{}

	cursor, err := c.collection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.ctx)

	for cursor.Next(c.ctx) {
		var carAccident domain.CarAccident
		if err := cursor.Decode(&carAccident); err != nil {
			return nil, err
		}
		carAccidents = append(carAccidents, &carAccident)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return carAccidents, nil
}

func (c CarAccidentServiceImpl) GetCarAccidentById(carAccidentId string, ctx context.Context) (*domain.CarAccident, error) {
	objID, err := primitive.ObjectIDFromHex(carAccidentId)
	if err != nil {
		return nil, err
	}

	var carAccident domain.CarAccident
	err = c.collection.FindOne(c.ctx, bson.M{"_id": objID}).Decode(&carAccident)
	if err != nil {
		return nil, err
	}
	return &carAccident, nil
}

func (c CarAccidentServiceImpl) GetAllCarAccidentsByType(carAccidentType domain.CarAccidentType) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"car_accident_type": carAccidentType}

	cursor, err := c.collection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.ctx)

	for cursor.Next(c.ctx) {
		var carAccident domain.CarAccident
		if err := cursor.Decode(&carAccident); err != nil {
			return nil, err
		}
		carAccidents = append(carAccidents, &carAccident)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return carAccidents, nil
}

func (c CarAccidentServiceImpl) GetAllCarAccidentsByDegree(degreeOfAccident domain.DegreeOfAccident) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"degree_of_accident": degreeOfAccident}

	cursor, err := c.collection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.ctx)

	for cursor.Next(c.ctx) {
		var carAccident domain.CarAccident
		if err := cursor.Decode(&carAccident); err != nil {
			return nil, err
		}
		carAccidents = append(carAccidents, &carAccident)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return carAccidents, nil
}

func (c CarAccidentServiceImpl) GetAllCarAccidentsByPolicemanID(policemanID string) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"policeman_id": policemanID}

	cursor, err := c.collection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.ctx)

	for cursor.Next(c.ctx) {
		var carAccident domain.CarAccident
		if err := cursor.Decode(&carAccident); err != nil {
			return nil, err
		}
		carAccidents = append(carAccidents, &carAccident)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return carAccidents, nil
}

func (c CarAccidentServiceImpl) GetAllCarAccidentsByDriver(driverEmail string) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"driver_email": driverEmail}

	cursor, err := c.collection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c.ctx)

	for cursor.Next(c.ctx) {
		var carAccident domain.CarAccident
		if err := cursor.Decode(&carAccident); err != nil {
			return nil, err
		}
		carAccidents = append(carAccidents, &carAccident)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return carAccidents, nil
}
