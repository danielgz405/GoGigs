package database

import (
	"context"

	"github.com/dg/acordia/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *MongoRepo) InsertPackage(ctx context.Context, packageEvent *models.InsertPackage) (*models.Package, error) {
	collection := repo.client.Database("gigs").Collection("events")
	result, err := collection.InsertOne(ctx, packageEvent)
	if err != nil {
		return nil, err
	}
	createdPackage, err := repo.GetPackageById(ctx, result.InsertedID.(primitive.ObjectID).Hex())
	if err != nil {
		return nil, err
	}
	return createdPackage, nil
}

func (repo *MongoRepo) GetPackageById(ctx context.Context, id string) (*models.Package, error) {
	collection := repo.client.Database("gigs").Collection("events")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var packageEvent models.Package
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&packageEvent)
	if err != nil {
		return nil, err
	}
	return &packageEvent, nil
}
