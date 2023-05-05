package database

import (
	"context"

	"github.com/dg/acordia/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertVisitor(ctx context.Context, visitor *models.InsertVisitor) (*models.Visitor, error) {
	collection := repo.client.Database("gigs").Collection("events")
	result, err := collection.InsertOne(ctx, visitor)
	if err != nil {
		return nil, err
	}
	createdVisitor, err := repo.GetVisitorById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdVisitor, nil
}

func (repo *MongoRepo) GetVisitorById(ctx context.Context, id string) (*models.Visitor, error) {
	collection := repo.client.Database("gigs").Collection("events")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var visitor models.Visitor
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&visitor)
	if err != nil {
		return nil, err
	}
	return &visitor, nil
}
