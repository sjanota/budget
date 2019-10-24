package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) CreateMonthlyReport(ctx context.Context, budgetID primitive.ObjectID, month models.Month, plans []*models.PlanInput) (*models.MonthlyReport, error) {
	plansToInsert := make([]*models.Plan, 0, len(plans))
	for _, p := range plans {
		plansToInsert = append(plansToInsert, &models.Plan{
			ID:              primitive.NewObjectID(),
			Title:           p.Title,
			CurrentAmount:   p.CurrentAmount,
			RecurringAmount: p.RecurringAmount,
			FromEnvelopeID:  p.FromEnvelopeID,
			ToEnvelopeID:    p.ToEnvelopeID,
		})
	}
	toInsert := &models.MonthlyReport{
		ID: models.MonthlyReportID{
			Month:    month,
			BudgetID: budgetID,
		},
		Expenses:  make([]*models.Expense, 0),
		Transfers: make([]*models.Transfer, 0),
		Plans:     plansToInsert,
	}

	_, err := s.monthlyReports.InsertOne(ctx, toInsert)
	if err != nil {
		if isDuplicateKeyError(err) {
			return nil, ErrAlreadyExists
		}
		return nil, err
	}
	return toInsert, nil
}

func (s *Storage) GetMonthlyReport(ctx context.Context, id models.MonthlyReportID) (*models.MonthlyReport, error) {
	result, err := s.monthlyReports.FindOneByID(ctx, id)
	if err == ErrNoReport {
		return nil, nil
	}
	return result, err
}
