package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"police-service/domain"
	"time"
)

type DelictServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (d *DelictServiceImpl) GetAllDelictsForDriverByDelictType(driverID string) ([]*domain.Delict, error) {
	var delicts []*domain.Delict
	filter := bson.M{
		"driver_identification_number": driverID,
		"delict_type":                  domain.DrivingUnderTheInfluenceOfAlcohol,
	}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
		var delict domain.Delict
		if err := cursor.Decode(&delict); err != nil {
			return nil, err
		}
		delicts = append(delicts, &delict)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return delicts, nil
}

func NewDelictServiceImpl(collection *mongo.Collection, ctx context.Context) DelictService {
	return &DelictServiceImpl{collection, ctx}
}

func (d *DelictServiceImpl) InsertDelict(delict *domain.DelictCreate, policemanID string) (*domain.Delict, string, error) {
	delict.PolicemanID = policemanID
	currentTime := time.Now()
	dateTime := primitive.NewDateTimeFromTime(currentTime)
	delict.Date = dateTime
	var trafficDelict domain.Delict
	trafficDelict.ID = primitive.NewObjectID()
	trafficDelict.PolicemanID = policemanID
	trafficDelict.DriverIdentificationNumber = delict.DriverIdentificationNumber
	trafficDelict.VehicleLicenceNumber = delict.VehicleLicenceNumber
	trafficDelict.DriverEmail = delict.DriverEmail
	trafficDelict.Date = dateTime
	trafficDelict.Location = delict.Location
	trafficDelict.Description = delict.Description
	trafficDelict.DelictType = delict.DelictType
	trafficDelict.DelictStatus = delict.DelictStatus
	trafficDelict.PriceOfFine = delict.PriceOfFine
	trafficDelict.NumberOfPenaltyPoints = delict.NumberOfPenaltyPoints

	result, err := d.collection.InsertOne(context.Background(), trafficDelict)
	if err != nil {
		return nil, "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, "", errors.New("failed to get inserted ID")
	}

	insertedID = result.InsertedID.(primitive.ObjectID)

	log.Printf("Inserted Delict DB : %+v\n", &trafficDelict)

	log.Printf("Inserted insertedID DB : %+v\n", insertedID)

	return &trafficDelict, insertedID.Hex(), nil
}

func (d *DelictServiceImpl) GetAllDelicts() ([]*domain.Delict, error) {
	var delicts []*domain.Delict
	filter := bson.D{}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
		var delict domain.Delict
		if err := cursor.Decode(&delict); err != nil {
			return nil, err
		}
		delicts = append(delicts, &delict)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return delicts, nil
}

func (d DelictServiceImpl) GetDelictById(delictId string, ctx context.Context) (*domain.Delict, error) {
	objID, err := primitive.ObjectIDFromHex(delictId)
	if err != nil {
		return nil, err
	}

	var delict domain.Delict
	err = d.collection.FindOne(d.ctx, bson.M{"_id": objID}).Decode(&delict)
	if err != nil {
		return nil, err
	}
	return &delict, nil
}

func (d *DelictServiceImpl) GetAllDelictsByPolicemanID(policemanID string) ([]*domain.Delict, error) {
	var delicts []*domain.Delict
	filter := bson.M{"policeman_id": policemanID}
	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)
	for cursor.Next(d.ctx) {
		var delict domain.Delict
		if err := cursor.Decode(&delict); err != nil {
			return nil, err
		}
		delicts = append(delicts, &delict)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return delicts, nil
}

func (d *DelictServiceImpl) GetAllDelictsByDriver(driverEmail string) ([]*domain.Delict, error) {
	var delicts []*domain.Delict
	filter := bson.M{"driver_email": driverEmail}
	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)
	for cursor.Next(d.ctx) {
		var delict domain.Delict
		if err := cursor.Decode(&delict); err != nil {
			return nil, err
		}
		delicts = append(delicts, &delict)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return delicts, nil
}

func (d *DelictServiceImpl) UpdateDelictStatus(delict *domain.Delict) error {
	filter := bson.M{"_id": delict.ID}
	update := bson.M{"$set": bson.M{"delict_status": delict.DelictStatus}}
	_, err := d.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (d *DelictServiceImpl) GetAllDelictsByDelictType(delictType domain.DelictType) ([]*domain.Delict, error) {
	var delicts []*domain.Delict
	filter := bson.M{"delict_type": delictType}

	cursor, err := d.collection.Find(d.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(d.ctx)

	for cursor.Next(d.ctx) {
		var delict domain.Delict
		if err := cursor.Decode(&delict); err != nil {
			return nil, err
		}
		delicts = append(delicts, &delict)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return delicts, nil
}

func (d *DelictServiceImpl) GetAllDelictsByDelictTypeAndYear(delictType domain.DelictType, year int) ([]*domain.Delict, error) {
	var delicts []*domain.Delict
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := startOfYear.AddDate(1, 0, 0)

	filter := bson.M{
		"delict_type": delictType,
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
		var delict domain.Delict
		if err := cursor.Decode(&delict); err != nil {
			return nil, err
		}
		delicts = append(delicts, &delict)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return delicts, nil
}
