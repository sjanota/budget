package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type categoryResolver struct {
	*Resolver
}

func (r *categoryResolver) Envelope(ctx context.Context, obj *models.Category) (*models.Envelope, error) {
	return r.Storage.GetEnvelope(ctx, budgetFromContext(ctx), obj.EnvelopeID)
}
