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

func (d CarAccidentServiceImpl) InsertCarAccident(carAccident *domain.CarAccidentCreate, policemanID string) (*domain.CarAccident, string, error) {
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

	result, err := d.collection.InsertOne(context.Background(), trafficCarAccident)
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

func (d CarAccidentServiceImpl) GetAllCarAccident() ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.D{}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
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

func (d CarAccidentServiceImpl) GetCarAccidentById(carAccidentId string, ctx context.Context) (*domain.CarAccident, error) {
	objID, err := primitive.ObjectIDFromHex(carAccidentId)
	if err != nil {
		return nil, err
	}

	var carAccident domain.CarAccident
	err = d.collection.FindOne(d.ctx, bson.M{"_id": objID}).Decode(&carAccident)
	if err != nil {
		return nil, err
	}
	return &carAccident, nil
}

func (d CarAccidentServiceImpl) GetAllCarAccidentsByType(carAccidentType domain.CarAccidentType) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"car_accident_type": carAccidentType}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
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

func (d *CarAccidentServiceImpl) GetAllCarAccidentsByTypeAndYear(carAccidentType domain.CarAccidentType, year int) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := startOfYear.AddDate(1, 0, 0)

	filter := bson.M{
		"car_accident_type": carAccidentType,
		"date": bson.M{
			"$gte": startOfYear,
			"$lt":  endOfYear,
		},
	}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
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

func (d CarAccidentServiceImpl) GetAllCarAccidentsByDegree(degreeOfAccident domain.DegreeOfAccident) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"degree_of_accident": degreeOfAccident}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
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

func (d *CarAccidentServiceImpl) GetAllCarAccidentsByDegreeAndYear(degreeOfAccident domain.DegreeOfAccident, year int) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := startOfYear.AddDate(1, 0, 0)

	filter := bson.M{
		"degree_of_accident": degreeOfAccident,
		"date": bson.M{
			"$gte": startOfYear,
			"$lt":  endOfYear,
		},
	}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
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

func (d CarAccidentServiceImpl) GetAllCarAccidentsByPolicemanID(policemanID string) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"policeman_id": policemanID}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
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

func (d CarAccidentServiceImpl) GetAllCarAccidentsByDriver(driverEmail string) ([]*domain.CarAccident, error) {
	var carAccidents []*domain.CarAccident
	filter := bson.M{"driver_email": driverEmail}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
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
