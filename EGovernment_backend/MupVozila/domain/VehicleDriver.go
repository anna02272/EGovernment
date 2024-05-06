package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type VehicleDriver struct {
	ID                    primitive.ObjectID `bson:"_id" json:"id"`
	IdentificationNumber  string             `bson:"identification_number" json:"identification_number"`
	Name                  string             `bson:"name" json:"name"`
	LastName              string             `bson:"last_name" json:"last_name"`
	DateOfBirth           time.Time          `bson:"date_of_birth" json:"date_of_birth"`
	HasDelict             bool               `bson:"has_delict" json:"has_delict"`
	Gender                Gender             `bson:"gender" json:"gender"`
	NumberOfPenaltyPoints int64              `bson:"number_of_penalty_points" json:"number_of_penalty_points"`
}

type Gender string

const (
	Male   = "Male"
	Female = "Female"
)
