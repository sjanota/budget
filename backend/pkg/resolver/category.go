package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/category_resolver_storage.go -package=mocks github.com/sjanota/budget/backend/pkg/resolver CategoryResolverStorage
type CategoryResolverStorage interface {
	GetEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Envelope, error)
}

type categoryResolver struct {
	Storage CategoryResolverStorage
}

func (r *categoryResolver) Envelope(ctx context.Context, obj *models.Category) (*models.Envelope, error) {
	return r.Storage.GetEnvelope(ctx, obj.BudgetID, obj.EnvelopeID)
}
