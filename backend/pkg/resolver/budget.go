package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BudgetResolver struct {
	*Resolver
	Storage *storage.Budget
}

func (r *BudgetResolver) UpdateExpense(ctx context.Context, id primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses().ReplaceByID(ctx, id, input)
}

func (r *BudgetResolver) DeleteExpense(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	return r.Storage.Expenses().DeleteByID(ctx, id)
}

func (r *BudgetResolver) CreateExpense(ctx context.Context, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses().InsertOne(ctx, input)
}

func (r *BudgetResolver) ExpenseEvent(ctx context.Context) (<-chan *models.ExpenseEvent, error) {
	return r.Storage.Expenses().Watch(ctx)
}