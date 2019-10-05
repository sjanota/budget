package mock

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
		Integer: rand.Int(),
		Decimal: rand.Int() % 100,
	}
}

func Month() time.Month {
	return time.Month(rand.Int()%12 + 1)
}

func Year() int {
	return rand.Int()%50 + 1990
}

func ExpenseInput(date models.Date, accountID, categoryID1, categoryID2 primitive.ObjectID) *models.ExpenseInput {
	return &models.ExpenseInput{
		Title: Name(),
		Categories: []*models.ExpenseCategoryInput{
			{categoryID1, *Amount()},
			{categoryID2, *Amount()},
		},
		AccountID:   accountID,
		TotalAmount: models.Amount{},
		Date:        date,
	}
}

func MonthlyReportID(budgetID primitive.ObjectID, date ...models.Date) models.MonthlyReportID {
	d := Date()
	if len(date) > 0 {
		d = date[0]
	}

	return models.MonthlyReportID{
		Month:    d.Month,
		Year:     d.Year,
		BudgetID: budgetID,
	}
}

func DateInReport(report *models.MonthlyReport) models.Date {
	return models.Date{
		Year:  report.Year(),
		Month: report.Month(),
		Day:   rand.Int()%29 + 1,
	}
}

func Date() models.Date {
	return models.Date{
		Year:  rand.Int()%100 + 1990,
		Month: time.Month(rand.Int()%12 + 1),
		Day:   rand.Int()%29 + 1,
	}
}
