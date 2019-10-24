package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) CreateMonthlyReport(ctx context.Context, budgetID primitive.ObjectID, month models.Month, plans []*models.Plan) (*models.MonthlyReport, error) {
	toInsert := &models.MonthlyReport{
		ID: models.MonthlyReportID{
			Month:    month,
			BudgetID: budgetID,
		},
		Expenses:  make([]*models.Expense, 0),
		Transfers: make([]*models.Transfer, 0),
		Plans:     plans,
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
