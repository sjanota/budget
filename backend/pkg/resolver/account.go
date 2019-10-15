package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/schema"
)

var _ schema.AccountResolver = &accountResolver{}

type accountResolver struct {
	*Resolver
}

func (r *accountResolver) Balance(ctx context.Context, obj *models.Account) (models.Amount, error) {
	budget, err := r.Storage.GetBudget(ctx, budgetFromContext(ctx))
	if err != nil {
		return models.NewAmount(), err
	}

	forExpenses, err := r.Storage.GetExpensesTotalForAccount(ctx, budget.CurrentMonthID(), obj.ID)
	if err != nil {
		return models.NewAmount(), err
	}

	forTransfers, err := r.Storage.GetTransfersTotalForAccount(ctx, budget.CurrentMonthID(), obj.ID)
	if err != nil {
		return models.NewAmount(), err
	}
	sub := forTransfers.Sub(forExpenses)
	return sub, nil
}
