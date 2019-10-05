package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) CreateMonthlyReport(ctx context.Context, budgetID primitive.ObjectID, input *models.MonthlyReportInput) (*models.MonthlyReport, error) {
	toInsert := &models.MonthlyReport{
		Month:     input.Month,
		Year:      input.Year,
		Expenses:  make([]*models.Expense, 0),
		Transfers: make([]*models.Transfer, 0),
		Plans:     make([]*models.Plan, 0),
		BudgetID:  budgetID,
	}

	res, err := s.monthlyReports.InsertOne(ctx, toInsert)
	if err != nil {
		if isDuplicateKeyError(err) {
			return nil, ErrAlreadyExists
		}
		return nil, err
	}

	toInsert.ID = res.InsertedID.(primitive.ObjectID)
	return toInsert, nil
}

func (s *Storage) GetMonthlyReport(ctx context.Context, budgetID, id primitive.ObjectID) (*models.MonthlyReport, error) {
	find := doc{
		"_id":      id,
		"budgetid": budgetID,
	}

	res := s.monthlyReports.FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := &models.MonthlyReport{}
	err := res.Decode(result)
	return result, err
}

func (s *Storage) createMonthlyReportIndexes(ctx context.Context) error {
	index := mongo.IndexModel{
		Keys: doc{
			"budgetid": 1,
			"month":    1,
			"year":     1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := s.monthlyReports.Indexes().CreateOne(ctx, index)
	return err
}
