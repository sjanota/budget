package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

//go:generate mockgen -destination=../mocks/budget_resolver_storage.go -package=mocks github.com/sjanota/budget/backend/pkg/resolver BudgetResolverStorage
type BudgetResolverStorage interface {
	GetMonthlyReport(ctx context.Context, id models.MonthlyReportID) (*models.MonthlyReport, error)
}

type budgetResolver struct {
	Storage BudgetResolverStorage
}

func (r *budgetResolver) CurrentMonth(ctx context.Context, obj *models.Budget) (*models.MonthlyReport, error) {
	return r.Storage.GetMonthlyReport(ctx, models.MonthlyReportID{BudgetID: obj.ID, Month: obj.CurrentMonth})
}
