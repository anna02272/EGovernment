package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Lastname    string             `bson:"lastname" json:"lastname"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	PhoneNumber int                `bson:"phone_number" json:"phone_number"`
	Category    CategoryPerson     `bson:"category" json:"category"`
	Question    string             `bson:"question" json:"question"`
	Date        primitive.DateTime `bson:"date" json:"date"`
}

type CategoryPerson string

const (
	Student     = "Student"
	Scientist   = "Scientist"
	Researcher  = "Researcher"
	Analyst     = "Analyst"
	PrivateUser = "PrivateUser"
	Other       = "Other"
)
