package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type budgetResolver struct {
	*Resolver
}

func (r *budgetResolver) CurrentMonth(ctx context.Context, obj *models.Budget) (*models.MonthlyReport, error) {
	return r.Storage.GetMonthlyReport(ctx, models.MonthlyReportID{BudgetID: obj.ID, Month: obj.CurrentMonth})
}
