package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	OPEN      Status = "OPEN"
	CLOSED    Status = "CLOSED"
	WAITING   Status = "WAITING"
	SCHEDULED Status = "SCHEDULED"
	REJECTED  Status = "REJECTED"
)

type Subject struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Judgment    string             `bson:"judgment" json:"judgment"`
	Status      Status             `bson:"status" json:"status"`
	Compromis   string             `bson:"compromis" json:"compromis"`
	ViolationID string             `bson:"violation_id" json:"violation_id"`
	Accused     Citizen            `bson:"accused" json:"accused"`
}