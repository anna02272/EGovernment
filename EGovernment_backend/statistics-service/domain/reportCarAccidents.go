package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportCarAccidentType struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Date        primitive.DateTime `bson:"date" json:"date"`
	TotalNumber int                `bson:"total_number" json:"total_number"`
	Type        CarAccidentType    `bson:"type" json:"type"`
	Year        int                `bson:"year" json:"year"`
}

type ReportCarAccidentDegree struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Date        primitive.DateTime `bson:"date" json:"date"`
	TotalNumber int                `bson:"total_number" json:"total_number"`
	Degree      DegreeOfAccident   `bson:"degree" json:"degree"`
	Year        int                `bson:"year" json:"year"`
}

type CarAccident struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Date             primitive.DateTime `bson:"date" json:"date"`
	CarAccidentType  CarAccidentType    `bson:"car_accident_type" json:"car_accident_type"`
	DegreeOfAccident DegreeOfAccident   `bson:"degree_of_accident" json:"degree_of_accident"`
}

type CarAccidentType string

const (
	KnockingDownPedestrians        = "Obaranje pesaka"
	VehicleLandingFromRoad         = "Sletanje vozila sa puta"
	CollisionFromOppositeDirection = "Sudar iz suprotnog smera"
	CollisionInSameDirection       = "Sudar u istom smeru"
	SideCollision                  = "Bocni sudar"
	Rollover                       = "Prevrtanje vozila"
)

type DegreeOfAccident string

const (
	NoHarm             = "Bez stete"
	WithMaterialDamage = "Sa materijalnom stetom"
	WithInjuredPersons = "Sa povredjenim osobama"
	WithDeadBodies     = "Sa poginulim osobama"
)
