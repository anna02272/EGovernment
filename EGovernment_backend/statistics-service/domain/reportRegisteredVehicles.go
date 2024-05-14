package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportRegisteredVehicle struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Date        primitive.DateTime `bson:"date" json:"date"`
	TotalNumber int                `bson:"total_number" json:"total_number"`
	Category    Category           `bson:"category" json:"category"`
}

type Category string

const (
	A  = "A"
	B  = "B"
	B1 = "B1"
	A1 = "A1"
	C  = "C"
	AM = "AM"
	A2 = "A2"
)
