package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Image     string             `bson:"image" json:"image"`
	DesertRef string             `bson:"desertref" json:"desertref"`
}

type Profile struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Image     string             `bson:"image" json:"image"`
	DesertRef string             `bson:"desertref" json:"desertref"`
}

type InsertUser struct {
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

type UpdateUser struct {
	Id        string `bson:"_id" json:"_id"`
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Image     string `bson:"image" json:"image"`
	DesertRef string `bson:"desertref" json:"desertref"`
}

