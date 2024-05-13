package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Delict struct {
	ID                         primitive.ObjectID `bson:"_id" json:"id"`
	PolicemanID                string             `bson:"policeman_id" json:"policeman_id"`
	DriverIdentificationNumber string             `bson:"driver_identification_number" json:"driver_identification_number"`
	VehicleLicenceNumber       string             `bson:"vehicle_licence_number" json:"vehicle_licence_number"`
	DriverEmail                string             `bson:"driver_email" json:"driver_email"`
	Date                       primitive.DateTime `bson:"date" json:"date"`
	Location                   string             `bson:"location" json:"location"`
	Description                string             `bson:"description" json:"description"`
	DelictType                 DelictType         `bson:"delict_type" json:"delict_type"`
	NumberOfPenaltyPoints      int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}

type DelictType string

const (
	DrivingUnderAlchocolism                   = "DrivingUnderAlchocolism"
	Speeding                                  = "Speeding"
	DrivingUnderTheInfluenceOfAlcohol         = "DrivingUnderTheInfluenceOfAlcohol"
	ImproperOvertaking                        = "ImproperOvertaking"
	ImproperParking                           = "ImproperParking"
	FailureTooComplyWithTrafficLightsAndSigns = "FailureTooComplyWithTrafficLightsAndSigns"
	ImproperUseOfSeatBeltsAndChildSeats       = "ImproperUseOfSeatBeltsAndChildSeats"
	UsingMobilePhoneWhileDriving              = "UsingMobilePhoneWhileDriving"
	ImproperUseOfMotorVehicle                 = "ImproperUseOfMotorVehicle"
)

type DelictCreate struct {
	PolicemanID                string             `bson:"policeman_id" json:"policeman_id"`
	DriverIdentificationNumber string             `bson:"driver_identification_number" json:"driver_identification_number"`
	VehicleLicenceNumber       string             `bson:"vehicle_licence_number" json:"vehicle_licence_number"`
	DriverEmail                string             `bson:"driver_email" json:"driver_email"`
	Date                       primitive.DateTime `bson:"date" json:"date"`
	Location                   string             `bson:"location" json:"location"`
	Description                string             `bson:"description" json:"description"`
	DelictType                 DelictType         `bson:"delict_type" json:"delict_type"`
	NumberOfPenaltyPoints      int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}
