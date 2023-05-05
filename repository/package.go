package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

func InsertPackage(ctx context.Context, packageEvent *models.InsertPackage) (*models.Package, error) {
	return implementation.InsertPackage(ctx, packageEvent)
}
