package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
)

type budgetResolver struct {
	*Resolver
}

func (r *budgetResolver) CurrentMonth(ctx context.Context, obj *models.Budget) (*models.MonthlyReport, error) {
	report, err := r.Storage.GetMonthlyReport(ctx, models.MonthlyReportID{BudgetID: obj.ID, Month: obj.CurrentMonth})
	if err == storage.ErrNoReport {
		err = nil
		report = &models.MonthlyReport{
			ID:        models.MonthlyReportID{
				Month:    obj.CurrentMonth,
				BudgetID: obj.ID,
			},
			Expenses:  []*models.Expense{},
			Transfers: []*models.Transfer{},
			Plans:     []*models.Plan{},
		}
	}
	return report, err
}
