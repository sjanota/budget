package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/playground"
)

type envelopeResolver struct {
	*Resolver
	storage playground.Storage
}

func (r *envelopeResolver) Balance(ctx context.Context, envelope *models.Envelope) (*models.Amount, error) {
	var change models.Amount
	budget, err := r.storage.GetBudget(ctx, envelope.BudgetID)
	if err != nil {
		return nil, err
	}

	monthlyBudget, err := r.storage.GetCurrentMonthlyBudget(ctx, envelope.BudgetID)
	if err != nil {
		return nil, err
	}

	for _, expense := range monthlyBudget.Expenses {
		for _, expenseCategory := range expense.Categories {
			if budget.Category(expenseCategory.CategoryID).EnvelopeID == envelope.ID {
				change = change.Sub(expenseCategory.Balance)
			}
		}
	}

	for _, plan := range monthlyBudget.Plans {
		if plan.FromEnvelopeID == envelope.ID {
			change = change.Sub(plan.Balance)
		}
		if plan.ToEnvelopeID == envelope.ID {
			change = change.Add(plan.Balance)
		}
	}

	result := envelope.Balance.Add(change)
	return &result, nil
}
