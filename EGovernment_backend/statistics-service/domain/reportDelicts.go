package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportDelict struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Date        primitive.DateTime `bson:"date" json:"date"`
	TotalNumber int                `bson:"total_number" json:"total_number"`
	Type        DelictType         `bson:"type" json:"type"`
	Year        int                `bson:"year" json:"year"`
}

type Delict struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Date       primitive.DateTime `bson:"date" json:"date"`
	Location   string             `bson:"location" json:"location"`
	DelictType DelictType         `bson:"delict_type" json:"delict_type"`
}

type DelictType string

const (
	Speeding                                  = "Prekoracenje brzine"
	DrivingUnderTheInfluence                  = "Voznja pod uticajem narkotika"
	DrivingUnderTheInfluenceOfAlcohol         = "Voznja pod uticajem alkohola"
	ImproperOvertaking                        = "Nepravilno preticanje"
	ImproperParking                           = "Nepravilno parkiranje"
	FailureTooComplyWithTrafficLightsAndSigns = "Nepostovanje saobracajnih znakova"
	ImproperUseOfSeatBeltsAndChildSeats       = "Nepravilna upotreba sigurnosnog pojasa i sedista za decu"
	UsingMobilePhoneWhileDriving              = "Koriscenje mobilnog telefona tokom voznje"
	ImproperUseOfMotorVehicle                 = "Nepravilna upotreba motornih vozila"
	Other                                     = "Drugo"
)
