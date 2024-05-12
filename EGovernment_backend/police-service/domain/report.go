package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID                    primitive.ObjectID `bson:"_id" json:"id"`
	PolicemanID           primitive.ObjectID `bson:"policeman_id" json:"policeman_id"`
	DelictID              primitive.ObjectID `bson:"delict_id" json:"delict_id"`
	CarAccidentID         primitive.ObjectID `bson:"car_accident_id" json:"car_accident_id"`
	DriverEmail           string             `bson:"driver_email" json:"driver_email"`
	Date                  primitive.DateTime `bson:"date" json:"date"`
	Location              string             `bson:"location" json:"location"`
	Description           string             `bson:"description" json:"description"`
	CarAccidentType       CarAccidentType    `bson:"car_accident_type" json:"car_accident_type"`
	DelictType            DelictType         `bson:"delict_type" json:"delict_type"`
	DegreeOfAccident      DegreeOfAccident   `bson:"degree_of_accident" json:"degree_of_accident"`
	Status                Status             `bson:"status" json:"status"`
	NumberOfPenaltyPoints int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}

type Status string

const (
	InProgress        = "InProgress"
	TrialScheduled    = "TrialScheduled"
	SettlementReached = "SettlementReached"
	AdjournedTrial    = "AdjournedTrial"
	Rejected          = "Rejected"
	Closed            = "Closed"
)

type ReportCreate struct {
	//PolicemanID           primitive.ObjectID `bson:"policeman_id" json:"policeman_id"`
	//DelictID              primitive.ObjectID `bson:"delict_id" json:"delict_id"`
	//CarAccidentID         primitive.ObjectID `bson:"car_accident_id" json:"car_accident_id"`
	DriverEmail           string           `bson:"driver_email" json:"driver_email"`
	Location              string           `bson:"location" json:"location"`
	Description           string           `bson:"description" json:"description"`
	CarAccidentType       CarAccidentType  `bson:"car_accident_type" json:"car_accident_type"`
	DelictType            DelictType       `bson:"delict_type" json:"delict_type"`
	DegreeOfAccident      DegreeOfAccident `bson:"degree_of_accident" json:"degree_of_accident"`
	Status                Status           `bson:"status" json:"status"`
	NumberOfPenaltyPoints int64            `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}
