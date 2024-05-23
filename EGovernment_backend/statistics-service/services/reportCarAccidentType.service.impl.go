package services

import (
	"context"
	"errors"
	"statistics-service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportCarAccidentTypeServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewReportCarAccidentTypeImpl(collection *mongo.Collection, ctx context.Context) ReportCarAccidentTypeService {
	return &ReportCarAccidentTypeServiceImpl{collection, ctx}
}

func (r *ReportCarAccidentTypeServiceImpl) Create(report *domain.ReportCarAccidentType) (error, bool) {
	filter := bson.M{
		"type": report.Type,
		"year": report.Year,
	}

	existing := &domain.ReportCarAccidentType{}
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

func (r *ReportCarAccidentTypeServiceImpl) GetAll() ([]*domain.ReportCarAccidentType, error) {
	var reports []*domain.ReportCarAccidentType
	filter := bson.D{}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportCarAccidentType
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

func (r *ReportCarAccidentTypeServiceImpl) GetById(id string) (*domain.ReportCarAccidentType, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var report domain.ReportCarAccidentType
	err = r.collection.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportCarAccidentTypeServiceImpl) GetAllByCarAccidentType(carAccidentType domain.CarAccidentType) ([]*domain.ReportCarAccidentType, error) {
	var reports []*domain.ReportCarAccidentType
	filter := bson.M{"type": carAccidentType}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportCarAccidentType
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

func (r *ReportCarAccidentTypeServiceImpl) GetAllByCarAccidentTypeAndYear(carAccidentType domain.CarAccidentType, year int) ([]*domain.ReportCarAccidentType, error) {
	var reports []*domain.ReportCarAccidentType
	filter := bson.M{"type": carAccidentType, "year": year}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportCarAccidentType
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
