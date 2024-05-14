package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Text       string             `bson:"text" json:"text"`
	Attachment string             `bson:"attachment" json:"attachment"`
	Accepted   bool               `bson:"accepted" json:"accepted" `
	SendTo     User               `bson:"send_to" json:"send_to"`
	Date       primitive.DateTime `bson:"date" json:"date"`
}
