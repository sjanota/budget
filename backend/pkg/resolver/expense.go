package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type expenseResolver struct {
	*Resolver
}

func (r *expenseResolver) Account(ctx context.Context, obj *models.Expense) (*models.Account, error) {
	return r.Storage.GetAccount(ctx, obj.BudgetID, obj.AccountID)
}
