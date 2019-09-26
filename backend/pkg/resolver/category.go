package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type categoryResolver struct{ *Resolver }

func (r *categoryResolver) Envelope(ctx context.Context, obj *models.Category) (*models.Envelope, error) {
	return r.Storage.Envelopes(obj.BudgetID).FindByID(ctx, obj.EnvelopeID)
}
