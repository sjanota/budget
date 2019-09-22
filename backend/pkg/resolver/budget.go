package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BudgetOps struct{
	*Resolver
	id primitive.ObjectID
}

func (r *BudgetOps) UpdateExpense(ctx context.Context, id primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses().ReplaceByID(ctx, id, input)
}

func (r *BudgetOps) DeleteExpense(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	return r.Storage.Expenses().DeleteByID(ctx, id)
}

func (r *BudgetOps) CreateExpense(ctx context.Context, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses().InsertOne(ctx, input)
}