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

func (r *envelopeResolver) Balance(ctx context.Context, obj *models.Envelope) (models.Amount, error) {
	budget, err := r.Storage.GetBudget(ctx, budgetFromContext(ctx))
	if err != nil {
		return models.NewAmount(), err
	}

	forExpenses, err := r.Storage.GetExpensesTotalForEnvelope(ctx, budget.CurrentMonthID(), obj.ID)
	if err != nil {
		return models.NewAmount(), err
	}

	forTransfers, err := r.Storage.GetPlansTotalForEnvelope(ctx, budget.CurrentMonthID(), obj.ID)
	if err != nil {
		return models.NewAmount(), err
	}
	sub := forTransfers.Sub(forExpenses)
	return sub, nil
}
