package resolver

import (
	"context"
	"time"

	"github.com/sjanota/budget/backend/pkg/models"
)

type monthlyReportResolver struct {
	*Resolver
}

func (r *monthlyReportResolver) Problems(ctx context.Context, obj *models.MonthlyReport) ([]models.Problem, error) {
	problems := make([]models.Problem, 0)

	now := time.Now()
	if now.Year() == obj.Month().Year && now.Month() == obj.Month().Month {
		problems = append(problems, models.MonthStillInProgress{
			Severity: models.SeverityInfo,
		})
	}

	income := obj.TotalIncomeAmount()
	planned := obj.TotalPlannedAmount()
	if income != planned {
		problems = append(problems, models.Misplanned{
			Severity:    models.SeverityError,
			Overplanned: planned > income,
		})
	}

	budget, err := r.Storage.GetBudget(ctx, obj.ID.BudgetID)
	if err != nil {
		return nil, err
	}

	obj.ApplyTo(budget)

	for _, envelope := range budget.Envelopes {
		if envelope.Balance.IsNegative() {
			problems = append(problems, models.NegativeBalanceOnEnvelope{
				Severity: models.SeverityError,
				ID:       envelope.ID,
			})
		}

		if envelope.Limit != nil && envelope.Balance.IsBiggerThan(*envelope.Limit) {
			problems = append(problems, models.EnvelopeOverLimit{
				Severity: models.SeverityError,
				ID:       envelope.ID,
			})
		}
	}

	for _, account := range budget.Accounts {
		if account.Balance.IsNegative() {
			problems = append(problems, models.NegativeBalanceOnAccount{
				Severity: models.SeverityWarning,
				ID:       account.ID,
			})
		}
	}

	return problems, nil
}
