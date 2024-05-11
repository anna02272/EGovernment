package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Vehicle struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	RegistrationPlate string             `bson:"_registration_plate" json:"registration_plate"`
	VehicleModel      VehicleModel       `bson:"vehicle_model" json:"vehicle_model"`
	VehicleOwner      string             `bson:"vehicle_owner" json:"vehicle_owner"`
	RegistrationDate  time.Time          `bson:"registration_date" json:"registration_date"`
}

type VehicleCreate struct {
	RegistrationPlate string       `bson:"_registration_plate" json:"registration_plate"`
	VehicleModel      VehicleModel `bson:"vehicle_model" json:"vehicle_model"`
	VehicleOwner      string       `bson:"vehicle_owner" json:"vehicle_owner"`
	RegistrationDate  time.Time    `bson:"registration_date" json:"registration_date"`
}

type VehicleModel string

const (
	CarBMW    = "CarBMW"
	MotorOPEL = "EngineOPEL"
	Truck     = "Truck"
)
