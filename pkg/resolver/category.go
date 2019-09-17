package resolver

import (
	"context"
	"github.com/sjanota/budget/pkg/models"
)

type categoryResolver struct{ *Resolver }

func (r *categoryResolver) Envelope(ctx context.Context, obj *models.Category) (*models.Envelope, error) {
	return &models.Envelope{
		ID:          "123",
		Name:        "123",
		Available:   1.23,
		Expenses:    nil,
		BudgetPlans: nil,
	}, nil
}

func (c *categoryResolver) Expenses(ctx context.Context, obj *models.Category, since *string, until *string) ([]*models.Expense, error) {
	panic("implement me")
}

