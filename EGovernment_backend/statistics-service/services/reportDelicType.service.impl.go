package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"statistics-service/domain"
)

type ReportDelicTypeServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewReportDelicTypeImpl(collection *mongo.Collection, ctx context.Context) ReportDelicTypeService {
	return &ReportDelicTypeServiceImpl{collection, ctx}
}

func (r *ReportDelicTypeServiceImpl) Create(report *domain.ReportDelict) (error, bool) {
	filter := bson.M{
		"type": report.Type,
	}

	existing := &domain.ReportDelict{}
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

func (r *ReportDelicTypeServiceImpl) GetAll() ([]*domain.ReportDelict, error) {
	var reports []*domain.ReportDelict
	filter := bson.D{}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportDelict
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

func (r *ReportDelicTypeServiceImpl) GetById(id string) (*domain.ReportDelict, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var report domain.ReportDelict
	err = r.collection.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportDelicTypeServiceImpl) GetAllByDelictType(delictType domain.DelictType) ([]*domain.ReportDelict, error) {
	var reports []*domain.ReportDelict
	filter := bson.M{"type": delictType}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportDelict
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
