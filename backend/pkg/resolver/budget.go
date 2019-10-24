package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type budgetResolver struct {
	*Resolver
}

func (r *budgetResolver) CurrentMonth(ctx context.Context, obj *models.Budget) (*models.MonthlyReport, error) {
	report, err := r.Storage.GetMonthlyReport(ctx, models.MonthlyReportID{BudgetID: obj.ID, Month: obj.CurrentMonth})
	if err != nil {
		return nil, err
	}
	if report != nil {
		return report, nil
	}

	previousReportID := models.MonthlyReportID{
		Month:    obj.CurrentMonth.Previous(),
		BudgetID: obj.ID,
	}
	plans := make([]*models.Plan, 0)
	previousReport, err := r.Storage.GetMonthlyReport(ctx, previousReportID)
	if err != nil {
		return nil, err
	}

	if previousReport != nil {
		for _, p := range previousReport.Plans {
			if p.RecurringAmount != nil {
				p.CurrentAmount = *p.RecurringAmount
				plans = append(plans, p)
			}
		}
	}
	report = &models.MonthlyReport{
		ID: models.MonthlyReportID{
			Month:    obj.CurrentMonth,
			BudgetID: obj.ID,
		},
		Expenses:  []*models.Expense{},
		Transfers: []*models.Transfer{},
		Plans:     plans,
	}
	return report, nil
}
