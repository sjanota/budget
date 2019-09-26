package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type envelopeResolver struct {
	*Resolver
}

func (r *envelopeResolver) Balance(ctx context.Context, obj *models.Envelope) (*models.MoneyAmount, error) {
	return r.Storage.Expenses(obj.BudgetID).TotalBalanceForEnvelope(ctx, obj.ID)
}
