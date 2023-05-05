package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

func ListEvents(ctx context.Context) ([]models.Event, error) {
	return implementation.ListEvents(ctx)
}
