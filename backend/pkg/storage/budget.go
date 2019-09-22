package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type budgetsRepository struct {
	storage *Storage
}

func newBudgetsRepository(storage *Storage) *budgetsRepository {
	return &budgetsRepository{
		storage:  storage,
	}
}

func (r *budgetsRepository) session() *Budgets {
	return &Budgets{r}
}

type Budgets struct {
	*budgetsRepository
}

func (r *budgetsRepository) Create(ctx context.Context, name string) (budget *models.Budget, err error) {
	budget = &models.Budget{
		Name: name,
		Expenses: make([]*models.Expense, 0),
	}
	budget.ID, err = r.storage.insertOne(ctx, budget)
	return
}

func (r *budgetsRepository) FindByID(ctx context.Context, id primitive.ObjectID) (result *models.Budget, err error) {
	result = &models.Budget{}
	err = r.storage.findByID(ctx, id, result)
	return
}

func (r *budgetsRepository) Delete(ctx context.Context, id primitive.ObjectID) (result *models.Budget, err error) {
	result = &models.Budget{}
	err = r.storage.deleteByID(ctx, id, result)
	return
}
