package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CarAccident struct {
	ID                          primitive.ObjectID `bson:"_id" json:"id"`
	PolicemanID                 string             `bson:"policeman_id" json:"policeman_id"`
	DriverIdentificationNumber1 string             `bson:"driver_identification_number_first" json:"driver_identification_number_first"`
	DriverIdentificationNumber2 string             `bson:"driver_identification_number_second" json:"driver_identification_number_second"`
	VehicleLicenceNumber1       string             `bson:"vehicle_licence_number_first" json:"vehicle_licence_number_first"`
	VehicleLicenceNumber2       string             `bson:"vehicle_licence_number_second" json:"vehicle_licence_number_second"`
	DriverEmail                 string             `bson:"driver_email" json:"driver_email"`
	Date                        primitive.DateTime `bson:"date" json:"date"`
	Location                    string             `bson:"location" json:"location"`
	Description                 string             `bson:"description" json:"description"`
	CarAccidentType             CarAccidentType    `bson:"car_accident_type" json:"car_accident_type"`
	DegreeOfAccident            DegreeOfAccident   `bson:"degree_of_accident" json:"degree_of_accident"`
	NumberOfPenaltyPoints       int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
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

type CarAccidentCreate struct {
	PolicemanID                 string             `bson:"policeman_id" json:"policeman_id"`
	DriverIdentificationNumber1 string             `bson:"driver_identification_number_first" json:"driver_identification_number_first"`
	DriverIdentificationNumber2 string             `bson:"driver_identification_number_second" json:"driver_identification_number_second"`
	VehicleLicenceNumber1       string             `bson:"vehicle_licence_number_first" json:"vehicle_licence_number_first"`
	VehicleLicenceNumber2       string             `bson:"vehicle_licence_number_second" json:"vehicle_licence_number_second"`
	DriverEmail                 string             `bson:"driver_email" json:"driver_email"`
	Date                        primitive.DateTime `bson:"date" json:"date"`
	Location                    string             `bson:"location" json:"location"`
	Description                 string             `bson:"description" json:"description"`
	CarAccidentType             CarAccidentType    `bson:"car_accident_type" json:"car_accident_type"`
	DegreeOfAccident            DegreeOfAccident   `bson:"degree_of_accident" json:"degree_of_accident"`
	NumberOfPenaltyPoints       int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}
