package storage

import (
	"context"
	"time"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const monthly = "playground-monthly"
const budgets = "playground-budgets"

func (s *Storage) GetCurrentMonthlyBudget(ctx context.Context, budgetID primitive.ObjectID) (*models.MonthlyBudget, error) {
	panic("implement me")
}

func (s *Storage) GetCurrentExpensesForAccount(ctx context.Context, budgetID primitive.ObjectID, monthlyBudgetID primitive.ObjectID, accountID primitive.ObjectID) ([]*models.Expense, error) {
	panic("implement me")
}

func (s *Storage) EnsureMonthlyBudget(ctx context.Context, budgetID primitive.ObjectID, month time.Month, year uint) (*models.MonthlyBudget, error) {
	panic("implement me")
}

func (s *Storage) UpdateMonthlyBudget(ctx context.Context, budgetID primitive.ObjectID, monthlyBudget *models.MonthlyBudget) (*models.MonthlyBudget, error) {
	panic("implement me")
}

func (s *Storage) UpdateBudget(ctx context.Context, budget *models.Budget) (*models.Budget, error) {
	panic("implement me")
}
