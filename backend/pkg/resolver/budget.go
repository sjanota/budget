package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BudgetResolver struct {
	*Resolver
	budgetID primitive.ObjectID
}

func (r *BudgetResolver) Expenses(ctx context.Context, obj *models.Budget) ([]*models.Expense, error) {
	return r.Storage.
		Expenses(obj.ID).
		FindAll(ctx)
}
