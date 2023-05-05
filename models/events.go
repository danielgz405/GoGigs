package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	Id                 primitive.ObjectID `bson:"_id" json:"_id"`
	Into               bool               `json:"into" bson:"into"`
	CarType            string             `json:"carType" bson:"carType"`
	TypeIdentification string             `json:"type-identification" bson:"type-identification"`
	Identification     string             `json:"identification" bson:"identification"`
	Tower              string             `json:"tower" bson:"tower"`
	Floor              string             `json:"floor" bson:"floor"`
	Apto               string             `json:"apto" bson:"apto"`
	Automobile         bool               `json:"automobile" bson:"automobile"`
	EventType          string             `json:"event-type" bson:"event-type"`
	Date               string             `json:"date" bson:"date"`
	Autorization       bool               `json:"autorization" bson:"autorization"`
	Worker             string             `json:"worker" bson:"worker"`
	Description        string             `json:"description" bson:"description"`
}
