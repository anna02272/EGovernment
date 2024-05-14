package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"statistics-service/domain"
	"time"
)

type ReportDelicTypeServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewReportDelicTypeImpl(collection *mongo.Collection, ctx context.Context) ReportDelicTypeService {
	return &ReportDelicTypeServiceImpl{collection, ctx}
}

func (r *ReportDelicTypeServiceImpl) Create(report *domain.ReportDelict) (error, bool) {
	timeObj := report.Date.Time()
	year := timeObj.Year()

	filter := bson.M{
		"date": bson.M{"$gte": time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC), "$lt": time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)},
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
