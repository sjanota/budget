package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
)

type budgetsRepository struct {
	storage *Storage
}

func newBudgetsRepository(storage *Storage) *budgetsRepository {
	return &budgetsRepository{
		storage:  storage,
	}
}

func (b *budgetsRepository) CreateBudget(ctx context.Context, name string) (budget *models.Budget, err error) {
	budget = &models.Budget{
		Name: name,
		Expenses: make([]*models.Expense, 0),
	}
	budget.ID, err = b.storage.insertOne(ctx, budget)
	return
}
