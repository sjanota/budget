package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) CreateMonthlyReport(ctx context.Context, budgetID primitive.ObjectID, input *models.MonthlyReportInput) (*models.MonthlyReport, error) {
	toInsert := &models.MonthlyReport{
		ID: models.MonthlyReportID{
			Month:    input.Month,
			BudgetID: budgetID,
		},
		Expenses:  make([]*models.Expense, 0),
		Transfers: make([]*models.Transfer, 0),
		Plans:     make([]*models.Plan, 0),
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
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}
