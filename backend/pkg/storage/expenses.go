package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) CreateExpense(ctx context.Context, reportID models.MonthlyReportID, in *models.ExpenseInput) (*models.Expense, error) {
	if err := s.validateExpenseInput(ctx, reportID, in); err != nil {
		return nil, err
	}

	toInsertCategories := make([]*models.ExpenseCategory, len(in.Categories))
	for i, categoryIn := range in.Categories {
		toInsertCategories[i] = &models.ExpenseCategory{
			Amount:     categoryIn.Amount,
			CategoryID: categoryIn.CategoryID,
			BudgetID:   reportID.BudgetID,
		}
	}

	toInsert := &models.Expense{
		Title:      in.Title,
		Categories: toInsertCategories,
		AccountID:  in.AccountID,
		BudgetID:   reportID.BudgetID,
		Date:       in.Date,
	}

	find := doc{
		"_id": reportID,
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

func (s *Storage) GetExpenses(ctx context.Context, reportID models.MonthlyReportID) ([]*models.Expense, error) {
	project := doc{
		"expenses": 1,
	}
	opts := options.FindOne().SetProjection(project)
	result, err := s.monthlyReports.FindOneByID(ctx, reportID, opts)
	if err != nil {
		return nil, err
	}
	return result.Expenses, nil
}

func (s *Storage) GetExpensesTotalForAccount(ctx context.Context, reportID models.MonthlyReportID, accountID primitive.ObjectID) (*models.Amount, error) {
	result, err := s.monthlyReports.Aggregate(ctx, list{
		doc{"$match": doc{"_id": reportID}},
		doc{"$project": doc{"expenses": 1, "_id": 0}},
		doc{"$unwind": "$expenses"},
		doc{"$unwind": "$expenses.categories"},
		doc{"$match": doc{"expenses.accountid": accountID}},
		doc{"$group": doc{
			"_id": nil,
			"integer": doc{
				"$sum": "$expenses.categories.amount.integer",
			},
			"decimal": doc{
				"$sum": "$expenses.categories.amount.decimal",
			},
		}},
		doc{"$project": doc{
			"_id": 0,
			"integer": doc{
				"$sum": list{"$integer", doc{
					"$floor": doc{
						"$divide": list{"$decimal", 100},
					},
				},
				},
			},
			"decimal": doc{
				"$mod": list{"$decimal", 100},
			},
		}},
	})
	if err != nil {
		return nil, err
	}
	if !result.Next(ctx) {
		return &models.Amount{}, nil
	}

	amount := &models.Amount{}
	err = result.Decode(&amount)
	return amount, err
}

func (s *Storage) validateExpenseInput(ctx context.Context, reportID models.MonthlyReportID, in *models.ExpenseInput) error {
	if err := s.validateExpenseInputReferences(ctx, reportID.BudgetID, in); err != nil {
		return err
	}

	return s.validateExpenseInputMonth(reportID, in)
}

func (s *Storage) validateExpenseInputReferences(ctx context.Context, budgetID primitive.ObjectID, in *models.ExpenseInput) error {
	project := doc{
		"accounts": doc{
			"$elemMatch": doc{
				"_id": in.AccountID,
			},
		},
		"categories": 1,
	}
	opts := options.FindOne().SetProjection(project)
	budget, err := s.budgets.FindOneByID(ctx, budgetID, opts)
	if err != nil {
		return err
	}

	if len(budget.Accounts) == 0 {
		return ErrInvalidReference
	}
	for _, category := range in.Categories {
		if budget.Category(category.CategoryID) == nil {
			return ErrInvalidReference
		}
	}
	return nil
}

func (s *Storage) validateExpenseInputMonth(reportID models.MonthlyReportID, in *models.ExpenseInput) error {
	if !reportID.Month.Contains(in.Date) {
		return ErrWrongDate
	}

	return nil
}
