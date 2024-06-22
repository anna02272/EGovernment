package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Delict struct {
	ID                         primitive.ObjectID `bson:"_id" json:"id"`
	PolicemanID                string             `bson:"policeman_id" json:"policeman_id"`
	DriverIdentificationNumber string             `bson:"driver_identification_number" json:"driver_identification_number"`
	VehicleLicenceNumber       string             `bson:"vehicle_licence_number" json:"vehicle_licence_number"`
	DriverEmail                string             `bson:"driver_email" json:"driver_email"`
	DriverJmbg                 string             `bson:"driver_jmbg" json:"driver_jmbg"`
	Date                       primitive.DateTime `bson:"date" json:"date"`
	Location                   string             `bson:"location" json:"location"`
	Description                string             `bson:"description" json:"description"`
	DelictType                 DelictType         `bson:"delict_type" json:"delict_type"`
	DelictStatus               DelictStatus       `bson:"delict_status" json:"delict_status"`
	PriceOfFine                float64            `bson:"price_of_fine" json:"price_of_fine"`
	NumberOfPenaltyPoints      int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}
type DelictType string

const (
	Speeding                                  = "Speeding"
	DrivingUnderTheInfluence                  = "DrivingUnderTheInfluence"
	DrivingUnderTheInfluenceOfAlcohol         = "DrivingUnderTheInfluenceOfAlcohol"
	ImproperOvertaking                        = "ImproperOvertaking"
	ImproperParking                           = "ImproperParking"
	FailureTooComplyWithTrafficLightsAndSigns = "FailureTooComplyWithTrafficLightsAndSigns"
	ImproperUseOfSeatBeltsAndChildSeats       = "ImproperUseOfSeatBeltsAndChildSeats"
	UsingMobilePhoneWhileDriving              = "UsingMobilePhoneWhileDriving"
	ImproperUseOfMotorVehicle                 = "ImproperUseOfMotorVehicle"
)

type DelictStatus string

const (
	FineAwarded = "FineAwarded"
	FinePaid    = "FinePaid"
	SentToCourt = "SentToCourt"
)

type DelictCreate struct {
	PolicemanID                string             `bson:"policeman_id" json:"policeman_id"`
	DriverIdentificationNumber string             `bson:"driver_identification_number" json:"driver_identification_number"`
	VehicleLicenceNumber       string             `bson:"vehicle_licence_number" json:"vehicle_licence_number"`
	DriverEmail                string             `bson:"driver_email" json:"driver_email"`
	DriverJmbg                 string             `bson:"driver_jmbg" json:"driver_jmbg"`
	Date                       primitive.DateTime `bson:"date" json:"date"`
	Location                   string             `bson:"location" json:"location"`
	Description                string             `bson:"description" json:"description"`
	DelictType                 DelictType         `bson:"delict_type" json:"delict_type"`
	DelictStatus               DelictStatus       `bson:"delict_status" json:"delict_status"`
	PriceOfFine                float64            `bson:"price_of_fine" json:"price_of_fine"`
	NumberOfPenaltyPoints      int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}

type Citizen struct {
	JMBG        string    `bson:"jmbg" json:"jmbg"`
	Name        string    `bson:"name" json:"name"`
	Lastname    string    `bson:"lastname" json:"lastname"`
	DateOfBirth time.Time `bson:"date_of_birth" json:"date_of_birth"`
	//Address     string    `bson:"address" json:"address"`
}
