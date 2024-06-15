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
	KnockingDownPedestrians        = "KnockingDownPedestrians"
	VehicleLandingFromRoad         = "VehicleLandingFromRoad"
	CollisionFromOppositeDirection = "CollisionFromOppositeDirection"
	CollisionInSameDirection       = "CollisionInSameDirection"
	SideCollision                  = "SideCollision"
	Rollover                       = "Rollover"
)

type DegreeOfAccident string

const (
	NoHarm             = "NoHarm"
	WithMaterialDamage = "WithMaterialDamage"
	WithInjuredPersons = "WithInjuredPersons"
	WithDeadBodies     = "WithDeadBodies"
)
