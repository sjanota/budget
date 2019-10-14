package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/schema"
)

var _ schema.EnvelopeResolver = &envelopeResolver{}

type envelopeResolver struct {
	*Resolver
}

func (r *envelopeResolver) Balance(ctx context.Context, obj *models.Envelope) (*models.Amount, error) {
	budget, err := r.Storage.GetBudget(ctx, budgetFromContext(ctx))
	if err != nil {
		return nil, err
	}

	return r.Storage.GetExpensesTotalForEnvelope(ctx, budget.CurrentMonthID(), obj.ID)
}
