package database

import (
	"context"

	"github.com/dg/acordia/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *MongoRepo) ListEvents(ctx context.Context) ([]models.Event, error) {
	collection := repo.client.Database("gigs").Collection("events")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var events []models.Event
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}
