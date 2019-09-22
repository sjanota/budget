package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
)

type categoryResolver struct{ *Resolver }

func (r *categoryResolver) Envelope(ctx context.Context, obj *models.Category) (*models.Envelope, error) {
	return &models.Envelope{
		ID:   obj.EnvelopeID,
		Name: "123",
		Balance: &models.MoneyAmount{
			Integer: 12,
			Decimal: 54,
		},
	}, nil
}

func (r *categoryResolver) Expenses(ctx context.Context, obj *models.Category, since *string, until *string) ([]*models.Expense, error) {
	panic("implement me")
}
