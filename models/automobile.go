package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InsertAutomobile struct {
	Into           bool   `json:"into" bson:"into"`
	CarType        string `json:"carType" bson:"carType"`
	Identification string `json:"identification" bson:"identification"`
	Autorization   bool   `json:"autorization" bson:"autorization"`
	EventType      string `json:"event-type" bson:"event-type"`
	Date           string `json:"date" bson:"date"`
	Worker         string `json:"worker" bson:"worker"`
	Description    string `json:"description" bson:"description"`
}

type Automobile struct {
	Id             primitive.ObjectID `bson:"_id" json:"_id"`
	Into           bool               `json:"into" bson:"into"`
	CarType        string             `json:"carType" bson:"carType"`
	Identification string             `json:"identification" bson:"identification"`
	Autorization   bool               `json:"autorization" bson:"autorization"`
	EventType      string             `json:"event-type" bson:"event-type"`
	Date           string             `json:"date" bson:"date"`
	Worker         string             `json:"worker" bson:"worker"`
	Description    string             `json:"description" bson:"description"`
}
