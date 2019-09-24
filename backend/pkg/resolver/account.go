package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type accountResolver struct {
	*Resolver
}

func (r *accountResolver) Balance(ctx context.Context, obj *models.Account) (*models.MoneyAmount, error) {
	return r.Storage.Expenses(obj.BudgetID).TotalBalanceForAccount(ctx, obj.ID)
}
