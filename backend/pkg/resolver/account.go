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

func (r *accountResolver) Balance(ctx context.Context, obj *models.Account) (*models.Amount, error) {
	budget, err := r.Storage.GetBudget(ctx, budgetFromContext(ctx))
	if err != nil {
		return nil, err
	}

	return r.Storage.GetExpensesTotalForAccount(ctx, budget.CurrentMonthID(), obj.ID)
}
