package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type DriverLicence struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	VehicleDriver    string             `bson:"vehicle_driver" json:"vehicle_driver"`
	LicenceNumber    string             `bson:"licence_number" json:"licence_number"`
	LocationLicenced Location           `bson:"location_licenced" json:"location_licenced"`
	Categories       []Category         `bson:"categories" json:"categories"`
}

type DriverLicenceCreate struct {
	VehicleDriver    string     `bson:"vehicle_driver" json:"vehicle_driver"`
	LicenceNumber    string     `bson:"licence_number" json:"licence_number"`
	LocationLicenced Location   `bson:"location_licenced" json:"location_licenced"`
	Categories       []Category `bson:"categories" json:"categories"`
}

type Location string

const (
	NoviSad      = "Novi Sad"
	Smederevo    = "Smederevo"
	Beograd      = "Beograd"
	BackaPalanka = "BackaPalanka"
)

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
