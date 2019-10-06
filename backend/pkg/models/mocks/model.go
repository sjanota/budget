// Package mock_models contains utility functions to create mocked objects from models package
package mock_models

import (
	"math/rand"
	"time"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Name() *string {
	name := primitive.NewObjectID().Hex()
	return &name
}

func Amount() *models.Amount {
	return &models.Amount{
		Integer: rand.Int() % 10000000,
		Decimal: rand.Int() % 100,
	}
}

func ExpenseInput() *models.ExpenseInput {
	return &models.ExpenseInput{
		Title: Name(),
		Categories: []*models.ExpenseCategoryInput{
			ExpenseCategoryInput(),
			ExpenseCategoryInput(),
		},
		AccountID:   primitive.NewObjectID(),
		TotalAmount: *Amount(),
		Date:        Date(),
	}
}

func ExpenseCategoryInput() *models.ExpenseCategoryInput {
	return &models.ExpenseCategoryInput{
		CategoryID: primitive.NewObjectID(),
		Amount:     *Amount(),
	}
}

func ExpenseCategory() *models.ExpenseCategory {
	return &models.ExpenseCategory{
		Amount:     *Amount(),
		CategoryID: primitive.NewObjectID(),
		BudgetID:   primitive.NewObjectID(),
	}
}

func MonthlyReportID() models.MonthlyReportID {
	return models.MonthlyReportID{
		Month: Month(),
		BudgetID: primitive.NewObjectID(),
	}
}

func day() int {
	return rand.Int()%29 + 1
}

func year() int {
	return rand.Int()%100 + 1990
}

func month() time.Month {
	return time.Month(rand.Int()%12 + 1)
}

func DateInReport(report *models.MonthlyReport) models.Date {
	return models.Date{
		Year:  report.Month().Year,
		Month: report.Month().Month,
		Day:   day(),
	}
}

func Date() models.Date {
	return models.Date{
		Year:  year(),
		Month: month(),
		Day:   day(),
	}
}

func Month() models.Month {
	return models.Month{
		Year:  year(),
		Month: month(),
	}
}

func Budget() *models.Budget {
	return &models.Budget{
		ID:           primitive.NewObjectID(),
		Accounts:     []*models.Account{Account()},
		Envelopes:    []*models.Envelope{Envelope()},
		Categories:   []*models.Category{Category()},
		CurrentMonth: Month(),
	}
}

func Category() *models.Category {
	return &models.Category{
		ID:         primitive.NewObjectID(),
		Name:       *Name(),
		EnvelopeID: primitive.NewObjectID(),
		BudgetID:   primitive.NewObjectID(),
	}
}

func Account() *models.Account {
	return &models.Account{
		ID:       primitive.NewObjectID(),
		Name:     *Name(),
		Balance:  *Amount(),
		BudgetID: primitive.NewObjectID(),
	}
}

func Envelope() *models.Envelope {
	return &models.Envelope{
		ID:       primitive.NewObjectID(),
		Name:     *Name(),
		Limit:    Amount(),
		Balance:  *Amount(),
		BudgetID: primitive.NewObjectID(),
	}
}

func MonthlyReport() *models.MonthlyReport {
	return &models.MonthlyReport{
		ID:        MonthlyReportID(),
		Expenses:  []*models.Expense{Expense()},
		Transfers: []*models.Transfer{},
		Plans:     []*models.Plan{},
	}
}

func Expense() *models.Expense {
	return &models.Expense{
		Title:      Name(),
		Categories: []*models.ExpenseCategory{},
		Date:       Date(),
		AccountID:  primitive.NewObjectID(),
		BudgetID:   primitive.NewObjectID(),
	}
}
