package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
)

type Budget struct {
	*Storage
}

func (b *Budget) CreateBudget(ctx context.Context, name string) (budget *models.Budget, err error) {
	budget = &models.Budget{Name: name}
	budget.ID, err = b.insertOne(ctx, budget)
	return
}
