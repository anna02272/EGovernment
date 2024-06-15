package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Court struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	Address string             `bson:"address" json:"address"`
}

type Hearing struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SubjectID primitive.ObjectID `json:"subject_id"`
	Date      string             `bson:"date" json:"date"`
	JudgeID   primitive.ObjectID `bson:"judge_id" json:"judge_id"`
}

type Schedule struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HearingID primitive.ObjectID `bson:"hearing_id" json:"hearing_id"`
	CourtID   primitive.ObjectID `bson:"court_id" json:"court_id"`
	StartTime string             `bson:"start_time" json:"start_time"`
	EndTime   string             `bson:"end_time" json:"end_time"`
}
