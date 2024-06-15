package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"statistics-service/domain"
)

type ReportRegisteredVehiclesServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewReportRegisteredVehiclesServiceImpl(collection *mongo.Collection, ctx context.Context) ReportRegisteredVehiclesService {
	return &ReportRegisteredVehiclesServiceImpl{collection, ctx}
}

func (r *ReportRegisteredVehiclesServiceImpl) Create(report *domain.ReportRegisteredVehicle) (error, bool) {
	filter := bson.M{
		"category": report.Category,
		"year":     report.Year,
	}

	existing := &domain.ReportRegisteredVehicle{}
	err := r.collection.FindOne(context.Background(), filter).Decode(existing)

	if err == nil {
		update := bson.M{
			"$set": bson.M{
				"title":        report.Title,
				"total_number": report.TotalNumber,
				"date":         report.Date,
			},
		}

		_, err := r.collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return err, false
		}
		return nil, false
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		return err, false
	}

	_, err = r.collection.InsertOne(context.Background(), report)
	if err != nil {
		return err, false
	}

	return nil, true
}

func (r *ReportRegisteredVehiclesServiceImpl) GetAll() ([]*domain.ReportRegisteredVehicle, error) {
	var reports []*domain.ReportRegisteredVehicle
	filter := bson.D{}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportRegisteredVehicle
		if err := cursor.Decode(&report); err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *ReportRegisteredVehiclesServiceImpl) GetById(id string) (*domain.ReportRegisteredVehicle, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var report domain.ReportRegisteredVehicle
	err = r.collection.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportRegisteredVehiclesServiceImpl) GetAllByCategory(category domain.Category) ([]*domain.ReportRegisteredVehicle, error) {
	var reports []*domain.ReportRegisteredVehicle
	filter := bson.M{"category": category}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportRegisteredVehicle
		if err := cursor.Decode(&report); err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *ReportRegisteredVehiclesServiceImpl) GetAllByCategoryAndYear(category domain.Category, year int) ([]*domain.ReportRegisteredVehicle, error) {
	var reports []*domain.ReportRegisteredVehicle
	filter := bson.M{"category": category, "year": year}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportRegisteredVehicle
		if err := cursor.Decode(&report); err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
