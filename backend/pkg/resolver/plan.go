package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
)

type planResolver struct {
	*Resolver
}

func (r *planResolver) FromEnvelope(ctx context.Context, obj *models.Plan) (*models.Envelope, error) {
	if obj.FromEnvelopeID == nil {
		return nil, nil
	}
	return r.Storage.GetEnvelope(ctx, budgetFromContext(ctx), *obj.FromEnvelopeID)
}

func (r *planResolver) ToEnvelope(ctx context.Context, obj *models.Plan) (*models.Envelope, error) {
	return r.Storage.GetEnvelope(ctx, budgetFromContext(ctx), obj.ToEnvelopeID)
}


