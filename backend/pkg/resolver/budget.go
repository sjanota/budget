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

func (r *BudgetResolver) UpdateExpense(ctx context.Context, id primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.
		Expenses(r.budgetID).
		ReplaceByID(ctx, id, input)
}

func (r *BudgetResolver) DeleteExpense(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	return r.Storage.
		Expenses(r.budgetID).
		DeleteByID(ctx, id)
}

func (r *BudgetResolver) CreateExpense(ctx context.Context, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.
		Expenses(r.budgetID).
		Insert(ctx, input)
}
