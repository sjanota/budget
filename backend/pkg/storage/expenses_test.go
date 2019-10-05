package storage_test

import (
	"testing"

	"github.com/sjanota/budget/backend/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage/mock"
)

func TestStorage_CreateExpense(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)
	envelope1 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category1 := whenSomeCategoryExists(t, ctx, budget.ID, envelope1.ID)
	envelope2 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category2 := whenSomeCategoryExists(t, ctx, budget.ID, envelope2.ID)
	account := whenSomeAccountExists(t, ctx, budget.ID)

	t.Run("Success", func(t *testing.T) {
		input := mock.ExpenseInput(mock.DateInReport(report), account.ID, category1.ID, category2.ID)

		created, err := testStorage.CreateExpense(ctx, report.ID, input)
		require.NoError(t, err)
		assert.Equal(t, &models.Expense{
			Title: input.Title,
			Categories: []*models.ExpenseCategory{
				{input.Categories[0].Amount, input.Categories[0].CategoryID, budget.ID},
				{input.Categories[1].Amount, input.Categories[1].CategoryID, budget.ID},
			},
			AccountID: input.AccountID,
			BudgetID:  budget.ID,
			Date:      input.Date,
		}, created)
	})

	t.Run("Account does not exist", func(t *testing.T) {
		input := mock.ExpenseInput(mock.DateInReport(report), primitive.NewObjectID(), category1.ID, category2.ID)

		_, err := testStorage.CreateExpense(ctx, report.ID, input)
		require.EqualError(t, err, storage.ErrInvalidReference.Error())
	})

	t.Run("One of categories does not exist", func(t *testing.T) {
		input := mock.ExpenseInput(mock.DateInReport(report), account.ID, category1.ID, primitive.NewObjectID())

		_, err := testStorage.CreateExpense(ctx, report.ID, input)
		require.EqualError(t, err, storage.ErrInvalidReference.Error())
	})

	t.Run("Report does not exist", func(t *testing.T) {
		input := mock.ExpenseInput(mock.Date(), account.ID, category1.ID, category2.ID)

		_, err := testStorage.CreateExpense(ctx, mock.MonthlyReportID(budget.ID, input.Date), input)
		require.EqualError(t, err, storage.ErrNoReport.Error())
	})

	t.Run("Date does not match report", func(t *testing.T) {
		input := mock.ExpenseInput(mock.Date(), account.ID, category1.ID, category2.ID)

		_, err := testStorage.CreateExpense(ctx, report.ID, input)
		require.EqualError(t, err, storage.ErrWrongDate.Error())
	})
}

func TestStorage_GetExpenses(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)
	envelope1 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category1 := whenSomeCategoryExists(t, ctx, budget.ID, envelope1.ID)
	envelope2 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category2 := whenSomeCategoryExists(t, ctx, budget.ID, envelope2.ID)
	account := whenSomeAccountExists(t, ctx, budget.ID)
	expense1 := whenSomeExpenseExists(t, ctx, account.ID, category1.ID, category2.ID, report)
	expense2 := whenSomeExpenseExists(t, ctx, account.ID, category1.ID, category2.ID, report)

	t.Run("Success", func(t *testing.T) {
		got, err := testStorage.GetExpenses(ctx, report.ID)
		require.NoError(t, err)
		assert.ElementsMatch(t, []*models.Expense{expense1, expense2}, got)
	})

	t.Run("Report does not exist", func(t *testing.T) {
		_, err := testStorage.GetExpenses(ctx, mock.MonthlyReportID(budget.ID))
		require.EqualError(t, err, storage.ErrNoReport.Error())
	})
}
