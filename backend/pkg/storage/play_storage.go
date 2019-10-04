package storage

import (
	"context"
	"time"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) GetCurrentMonthlyBudget(ctx context.Context, budgetID primitive.ObjectID) (*models.MonthlyReport, error) {
	panic("implement me")
}

func (s *Storage) GetCurrentExpensesForAccount(ctx context.Context, budgetID primitive.ObjectID, monthlyBudgetID primitive.ObjectID, accountID primitive.ObjectID) ([]*models.Expense, error) {
	panic("implement me")
}

func (s *Storage) EnsureMonthlyBudget(ctx context.Context, budgetID primitive.ObjectID, month time.Month, year uint) (*models.MonthlyReport, error) {
	panic("implement me")
}

func (s *Storage) UpdateMonthlyBudget(ctx context.Context, budgetID primitive.ObjectID, monthlyBudget *models.MonthlyReport) (*models.MonthlyReport, error) {
	panic("implement me")
}

func (s *Storage) UpdateBudget(ctx context.Context, budget *models.Budget) (*models.Budget, error) {
	panic("implement me")
}
