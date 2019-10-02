package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/playground"

	"github.com/sjanota/budget/backend/pkg/models"
)

type accountResolver struct {
	*Resolver
	storage playground.Storage
}

func (r *accountResolver) Balance(ctx context.Context, account *models.Account) (*models.Amount, error) {
	var change models.Amount

	monthlyBudget, err := r.storage.GetCurrentMonthlyBudget(ctx, account.BudgetID)
	if err != nil {
		return nil, err
	}

	for _, expense := range monthlyBudget.Expenses {
		if expense.AccountID != nil && *expense.AccountID == account.ID {
			account.Balance = account.Balance.Sub(expense.Balance())
		}
	}

	for _, transfer := range monthlyBudget.Transfers {
		if transfer.FromAccountID == account.ID {
			account.Balance = account.Balance.Sub(transfer.Balance)
		}
		if transfer.ToAccountID == account.ID {
			account.Balance = account.Balance.Add(transfer.Balance)
		}
	}

	result := account.Balance.Add(change)
	return &result, nil
}
