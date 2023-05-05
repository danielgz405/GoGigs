package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InsertPackage struct {
	CompanyName        string `json:"company-name" bson:"company-name"`
	MessengerBoy       string `json:"messenger-boy" bson:"messenger-boy"`
	TypeIdentification string `json:"type-identification" bson:"type-identification"`
	Identification     string `json:"identification" bson:"identification"`
	Receiver           string `json:"receiver" bson:"receiver"`
	Tower              string `json:"tower" bson:"tower"`
	Office             string `json:"office" bson:"office"`
	Apto               string `json:"apto" bson:"apto"`
	PlaceSave          string `json:"place-save" bson:"place-save"`
	EventType          string `json:"event-type" bson:"event-type"`
	Date               string `json:"date" bson:"date"`
	Autorization       bool   `json:"autorization" bson:"autorization"`
	Worker             string `json:"worker" bson:"worker"`
	Description        string `json:"description" bson:"description"`
}

type Package struct {
	Id                 primitive.ObjectID `bson:"_id" json:"_id"`
	CompanyName        string             `json:"company-name" bson:"company-name"`
	MessengerBoy       string             `json:"messenger-boy" bson:"messenger-boy"`
	TypeIdentification string             `json:"type-identification" bson:"type-identification"`
	Identification     string             `json:"identification" bson:"identification"`
	Receiver           string             `json:"receiver" bson:"receiver"`
	Tower              string             `json:"tower" bson:"tower"`
	Office             string             `json:"office" bson:"office"`
	Apto               string             `json:"apto" bson:"apto"`
	PlaceSave          string             `json:"place-save" bson:"place-save"`
	EventType          string             `json:"event-type" bson:"event-type"`
	Date               string             `json:"date" bson:"date"`
	Autorization       bool               `json:"autorization" bson:"autorization"`
	Worker             string             `json:"worker" bson:"worker"`
	Description        string             `json:"description" bson:"description"`
}
