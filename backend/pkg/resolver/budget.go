package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/storage"

	"github.com/sjanota/budget/backend/pkg/models"
)

type budgetResolver struct {
	*Resolver
}

func (r *budgetResolver) CurrentMonth(ctx context.Context, obj *models.Budget) (*models.MonthlyReport, error) {
	var report *models.MonthlyReport
	var err error
	err = r.withMonthlyReport(ctx, obj.ID, func(reportID models.MonthlyReportID) error {
		report, err = r.Storage.GetMonthlyReport(ctx, reportID)
		if report == nil {
			return storage.ErrNoReport
		}
		return err
	})
	return report, err
}
