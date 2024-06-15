package domain

import "time"

type Citizen struct {
	JMBG        string    `bson:"jmbg" json:"jmbg"`
	Name        string    `bson:"name" json:"name"`
	Lastname    string    `bson:"lastname" json:"lastname"`
	DateOfBirth time.Time `bson:"date_of_birth" json:"date_of_birth"`
	//Address     string    `bson:"address" json:"address"`
}
