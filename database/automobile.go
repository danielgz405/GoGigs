package database

import (
	"context"

	"github.com/dg/acordia/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertAutomobile(ctx context.Context, automobile *models.InsertAutomobile) (*models.Automobile, error) {
	collection := repo.client.Database("gigs").Collection("events")
	result, err := collection.InsertOne(ctx, automobile)
	if err != nil {
		return nil, err
	}
	createdAutomobile, err := repo.GetAutomobileById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdAutomobile, nil
}

func (repo *MongoRepo) GetAutomobileById(ctx context.Context, id string) (*models.Automobile, error) {
	collection := repo.client.Database("gigs").Collection("events")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var automobile models.Automobile
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&automobile)
	if err != nil {
		return nil, err
	}
	return &automobile, nil
}
