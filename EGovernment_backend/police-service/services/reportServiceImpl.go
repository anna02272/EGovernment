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

type ReportServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewReportServiceImpl(collection *mongo.Collection, ctx context.Context) ReportService {
	return &ReportServiceImpl{collection, ctx}
}

func (r *ReportServiceImpl) InsertReport(report *domain.ReportCreate, policemanID string, delictID string, carAccidentID string) (*domain.Report, string, error) {
	policeman, err := primitive.ObjectIDFromHex(policemanID)
	if err != nil {
		return nil, "Error converting ID", nil
	}
	delict, err := primitive.ObjectIDFromHex(delictID)
	if err != nil {
		return nil, "Error converting ID", nil
	}
	carAccident, err := primitive.ObjectIDFromHex(carAccidentID)
	if err != nil {
		return nil, "Error converting ID", nil
	}
	//report.PolicemanID = objID
	var trafficReport domain.Report
	trafficReport.ID = primitive.NewObjectID()
	trafficReport.PolicemanID = policeman
	trafficReport.DelictID = delict
	trafficReport.CarAccidentID = carAccident
	trafficReport.DriverEmail = report.DriverEmail
	trafficReport.Date = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))
	trafficReport.Location = report.Location
	trafficReport.Description = report.Description
	trafficReport.CarAccidentType = report.CarAccidentType
	trafficReport.DelictType = report.DelictType
	trafficReport.DegreeOfAccident = report.DegreeOfAccident
	trafficReport.Status = report.Status
	trafficReport.NumberOfPenaltyPoints = report.NumberOfPenaltyPoints

	// Log the content of the trafficReport object
	log.Printf("Traffic Report: %+v\n", trafficReport)

	result, err := r.collection.InsertOne(context.Background(), report)
	if err != nil {
		log.Println("Error inserting report to database:", err)
		return nil, "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, "", errors.New("failed to get inserted ID")
	}

	insertedID = result.InsertedID.(primitive.ObjectID)

	// Log the inserted report
	log.Printf("Inserted traffic report: %+v\n", trafficReport)

	// Log the inserted report
	log.Printf("Inserted report: %+v\n", report)

	return &trafficReport, insertedID.Hex(), nil
}

func (r *ReportServiceImpl) GetAllReport() ([]*domain.Report, error) {
	var reports []*domain.Report
	filter := bson.D{}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.Report
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

func (r ReportServiceImpl) GetReportById(reportId string, ctx context.Context) (*domain.Report, error) {
	objID, err := primitive.ObjectIDFromHex(reportId)
	if err != nil {
		return nil, err
	}

	var report domain.Report
	err = r.collection.FindOne(r.ctx, bson.M{"_id": objID}).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportServiceImpl) GetAllReportsByDelictType(delictType domain.DelictType) ([]*domain.Report, error) {
	var reports []*domain.Report
	filter := bson.M{"delict_type": delictType}

	cursor, err := r.collection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	for cursor.Next(r.ctx) {
		var report domain.Report
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
