package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

func InsertAutomobile(ctx context.Context, automobile *models.InsertAutomobile) (*models.Automobile, error) {
	return implementation.InsertAutomobile(ctx, automobile)
}
