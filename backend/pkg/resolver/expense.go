package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
)

type expenseResolver struct{ *Resolver }

func (r *expenseResolver) Account(ctx context.Context, obj *models.Expense) (*models.Account, error) {
	return &models.Account{
		ID:        *obj.AccountID,
		Name:      "Konto Szymon",
		Available: 0.1,
		Expenses:  nil,
		Transfers: nil,
	}, nil
}

