package services

import (
	"context"
	"errors"
	"statistics-service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportCarAccidentDegreeServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewReportCarAccidentDegreeImpl(collection *mongo.Collection, ctx context.Context) ReportCarAccidentDegreeService {
	return &ReportCarAccidentDegreeServiceImpl{collection, ctx}
}

func (r *ReportCarAccidentDegreeServiceImpl) Create(report *domain.ReportCarAccidentDegree) (error, bool) {
	filter := bson.M{
		"degree": report.Degree,
		"year":   report.Year,
	}

	existing := &domain.ReportCarAccidentDegree{}
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

func (r *ReportCarAccidentDegreeServiceImpl) GetAll() ([]*domain.ReportCarAccidentDegree, error) {
	var reports []*domain.ReportCarAccidentDegree
	filter := bson.D{}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportCarAccidentDegree
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

func (r *ReportCarAccidentDegreeServiceImpl) GetById(id string) (*domain.ReportCarAccidentDegree, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var report domain.ReportCarAccidentDegree
	err = r.collection.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportCarAccidentDegreeServiceImpl) GetAllByCarAccidentDegree(degree domain.DegreeOfAccident) ([]*domain.ReportCarAccidentDegree, error) {
	var reports []*domain.ReportCarAccidentDegree
	filter := bson.M{"degree": degree}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportCarAccidentDegree
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

func (r *ReportCarAccidentDegreeServiceImpl) GetAllByCarAccidentDegreeAndYear(degree domain.DegreeOfAccident, year int) ([]*domain.ReportCarAccidentDegree, error) {
	var reports []*domain.ReportCarAccidentDegree
	filter := bson.M{"degree": degree, "year": year}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.ReportCarAccidentDegree
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
