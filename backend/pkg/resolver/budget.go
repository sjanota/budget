package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type budgetResolver struct {
	*Resolver
}

func (r *budgetResolver) Expenses(ctx context.Context, obj *models.Budget) ([]*models.Expense, error) {
	return r.Storage.Expenses(obj.ID).FindAll(ctx)
}
