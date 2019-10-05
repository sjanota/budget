package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) CreateExpense(ctx context.Context, budgetID primitive.ObjectID, reportID primitive.ObjectID, in *models.ExpenseInput) (*models.Expense, error) {
	if err := s.validateExpenseInput(ctx, budgetID, reportID, in); err != nil {
		return nil, err
	}

	toInsertCategories := make([]*models.ExpenseCategory, len(in.Categories))
	for i, categoryIn := range in.Categories {
		toInsertCategories[i] = &models.ExpenseCategory{
			Amount:     categoryIn.Amount,
			CategoryID: categoryIn.CategoryID,
			BudgetID:   budgetID,
		}
	}

	toInsert := &models.Expense{
		Title:      in.Title,
		Categories: toInsertCategories,
		AccountID:  in.AccountID,
		BudgetID:   budgetID,
		Date:       in.Date,
	}

	find := doc{
		"_id":      reportID,
		"budgetid": budgetID,
	}
	update := doc{
		"$push": doc{
			"expenses": toInsert,
		},
	}
	res, err := s.monthlyReports.UpdateOne(ctx, find, update)
	if err != nil {
		return nil, err
	} else if res.MatchedCount == 0 {
		return nil, ErrNoReport
	}
	return toInsert, nil
}

func (s *Storage) GetExpenses(ctx context.Context, budgetID primitive.ObjectID, monthID primitive.ObjectID) ([]*models.Expense, error) {
	find := doc{
		"_id":      monthID,
		"budgetid": budgetID,
	}
	project := doc{
		"_id":      0,
		"expenses": 1,
	}
	opts := options.FindOne().SetProjection(project)
	res := s.monthlyReports.FindOne(ctx, find, opts)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, ErrNoReport
	} else if err != nil {
		return nil, err
	}

	result := &models.MonthlyReport{}
	err := res.Decode(result)
	if err != nil {
		return nil, err
	}

	return result.Expenses, nil
}

func (s *Storage) validateExpenseInput(ctx context.Context, budgetID, reportID primitive.ObjectID, in *models.ExpenseInput) error {
	if err := s.validateExpenseInputReferences(ctx, budgetID, in); err != nil {
		return err
	}

	return s.validateExpenseInputMonth(ctx, budgetID, reportID, in)
}

func (s *Storage) validateExpenseInputReferences(ctx context.Context, budgetID primitive.ObjectID, in *models.ExpenseInput) error {
	find := doc{
		"_id": budgetID,
	}
	project := doc{
		"accounts": doc{
			"$elemMatch": doc{
				"_id": in.AccountID,
			},
		},
		"categories": 1,
	}
	opts := options.FindOne().SetProjection(project)
	res := s.budgets.FindOne(ctx, find, opts)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return ErrNoBudget
	} else if err != nil {
		return err
	}

	result := &models.Budget{}
	err := res.Decode(result)
	if err != nil {
		return err
	}

	if len(result.Accounts) == 0 {
		return ErrInvalidReference
	}
	for _, category := range in.Categories {
		if result.Category(category.CategoryID) == nil {
			return ErrInvalidReference
		}
	}
	return nil
}

func (s *Storage) validateExpenseInputMonth(ctx context.Context, budgetID, reportID primitive.ObjectID, in *models.ExpenseInput) error {
	find := doc{
		"_id": reportID,
		"budgetid": budgetID,
	}
	projection := doc {
		"month": 1,
		"year": 1,
	}
	opts := options.FindOne().SetProjection(projection)
	res := s.monthlyReports.FindOne(ctx, find, opts)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return ErrNoReport
	} else if err != nil {
		return err
	}

	result := &models.MonthlyReport{}
	err := res.Decode(result)
	if err != nil {
		return err
	}

	if result.Month != in.Date.Month {
		return ErrWrongDate
	}
	if result.Year != in.Date.Year {
		return ErrWrongDate
	}

	return nil

}
