package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email" validate:"required,email"`
	Name     string             `bson:"name" json:"name"`
	Lastname string             `bson:"lastname" json:"lastname"`
	UserRole UserRole           `bson:"userRole" json:"userRole"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"email" `
	Password string `json:"password" bson:"password"`
}

type UserResponse struct {
	Username string   `bson:"username" json:"username"`
	Email    string   `bson:"email" json:"email" validate:"required,email"`
	UserRole UserRole `bson:"userRole" json:"userRole"`
}

type UserRole string

const (
	Citizen   = "Citizen"
	Employe   = "Employe"
	Policeman = "Policeman"
)
