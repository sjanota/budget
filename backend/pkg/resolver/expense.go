package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type expenseResolver struct{ *Resolver }

func (r *expenseResolver) Account(ctx context.Context, obj *models.Expense) (*models.Account, error) {
	if obj.AccountID == nil {
		return nil, nil
	}
	return &models.Account{
		ID:   *obj.AccountID,
		Name: "Konto Szymon",
		Balance: &models.MoneyAmount{
			Integer: 10,
			Decimal: 99,
		},
	}, nil
}

type expenseEntryResolver struct{ *Resolver }

func (r *expenseEntryResolver) Category(ctx context.Context, obj *models.ExpenseEntry) (*models.Category, error) {
	return &models.Category{
		ID:          obj.CategoryID,
		Name:        "",
		Description: nil,
		EnvelopeID:  primitive.NewObjectID(),
	}, nil
}
