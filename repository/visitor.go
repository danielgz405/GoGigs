package repository

import (
	"context"

	"github.com/dg/acordia/models"
)

func InsertVisitor(ctx context.Context, visitor *models.InsertVisitor) (*models.Visitor, error) {
	return implementation.InsertVisitor(ctx, visitor)
}
