package storage_test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjanota/budget/backend/pkg/models"
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
		input := mock_models.ExpenseInput().
			WithDate(mock_models.DateInReport(report)).
			WithAccount(account.ID).
			WithCategories(
				mock_models.ExpenseCategoryInput().WithCategory(category1.ID),
				mock_models.ExpenseCategoryInput().WithCategory(category2.ID),
			)

		created, err := testStorage.CreateExpense(ctx, report.ID, input)
		require.NoError(t, err)
		require.NotEqual(t, primitive.ObjectID{}, created.ID)
		assert.Equal(t, &models.Expense{
			ID:    created.ID,
			Title: input.Title,
			Categories: []*models.ExpenseCategory{
				{input.Categories[0].Amount, input.Categories[0].CategoryID},
				{input.Categories[1].Amount, input.Categories[1].CategoryID},
			},
			AccountID: input.AccountID,
			Date:      input.Date,
		}, created)
	})

	t.Run("Account does not exist", func(t *testing.T) {
		input := mock_models.ExpenseInput().
			WithDate(mock_models.DateInReport(report)).
			WithCategories(
				mock_models.ExpenseCategoryInput().WithCategory(category1.ID),
				mock_models.ExpenseCategoryInput().WithCategory(category2.ID),
			)

		_, err := testStorage.CreateExpense(ctx, report.ID, input)
		require.EqualError(t, err, storage.ErrInvalidReference.Error())
	})

	t.Run("One of categories does not exist", func(t *testing.T) {
		input := mock_models.ExpenseInput().
			WithDate(mock_models.DateInReport(report)).
			WithAccount(account.ID).
			WithCategories(
				mock_models.ExpenseCategoryInput().WithCategory(category1.ID),
				mock_models.ExpenseCategoryInput(),
			)
		_, err := testStorage.CreateExpense(ctx, report.ID, input)
		require.EqualError(t, err, storage.ErrInvalidReference.Error())
	})

	t.Run("Report does not exist", func(t *testing.T) {
		input := mock_models.ExpenseInput().
			WithAccount(account.ID).
			WithCategories(
				mock_models.ExpenseCategoryInput().WithCategory(category1.ID),
				mock_models.ExpenseCategoryInput().WithCategory(category2.ID),
			)
		reportID := mock_models.MonthlyReportID().WithBudget(*budget).WithMonth(input.Date.ToMonth())

		_, err := testStorage.CreateExpense(ctx, *reportID, input)
		require.EqualError(t, err, storage.ErrNoReport.Error())
	})

	t.Run("Date does not match report", func(t *testing.T) {
		input := mock_models.ExpenseInput().
			WithAccount(account.ID).
			WithCategories(
				mock_models.ExpenseCategoryInput().WithCategory(category1.ID),
				mock_models.ExpenseCategoryInput().WithCategory(category2.ID),
			)

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
		reportID := mock_models.MonthlyReportID().WithBudget(*budget)
		_, err := testStorage.GetExpenses(ctx, *reportID)
		require.EqualError(t, err, storage.ErrNoReport.Error())
	})
}

func TestStorage_GetExpensesTotalForAccount(t *testing.T) {
	ctx := context.Background()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)
	account1 := whenSomeAccountExists(t, ctx, budget.ID)
	account2 := whenSomeAccountExists(t, ctx, budget.ID)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category := whenSomeCategoryExists(t, ctx, budget.ID, envelope.ID)
	expense1 := whenSomeExpenseExists(t, ctx, account1.ID, category.ID, category.ID, report)
	expense2 := whenSomeExpenseExists(t, ctx, account1.ID, category.ID, category.ID, report)
	_ = whenSomeExpenseExists(t, ctx, account2.ID, category.ID, category.ID, report)

	t.Run("Report exists", func(t *testing.T) {
		expectedTotal := expense1.TotalAmount().Add(expense2.TotalAmount())
		total, err := testStorage.GetExpensesTotalForAccount(ctx, report.ID, account1.ID)
		require.NoError(t, err)
		assert.Equal(t, expectedTotal, total)
	})

	t.Run("Report does not exist", func(t *testing.T) {
		total, err := testStorage.GetExpensesTotalForAccount(ctx, mock_models.MonthlyReportID(), account1.ID)
		require.NoError(t, err)
		assert.Equal(t, models.NewAmount(), total)
	})
}

func TestStorage_GetExpensesTotalForEnvelope(t *testing.T) {
	ctx := context.Background()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)
	account := whenSomeAccountExists(t, ctx, budget.ID)
	envelope1 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category1 := whenSomeCategoryExists(t, ctx, budget.ID, envelope1.ID)
	category2 := whenSomeCategoryExists(t, ctx, budget.ID, envelope1.ID)
	envelope2 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category3 := whenSomeCategoryExists(t, ctx, budget.ID, envelope2.ID)
	expense1 := whenSomeExpenseExists(t, ctx, account.ID, category1.ID, category2.ID, report)
	expense2 := whenSomeExpenseExists(t, ctx, account.ID, category1.ID, category3.ID, report)

	t.Run("Report exists", func(t *testing.T) {
		expectedTotal := expense1.TotalAmount().Add(expense2.Categories[0].Amount)
		total, err := testStorage.GetExpensesTotalForEnvelope(ctx, report.ID, envelope1.ID)
		require.NoError(t, err)
		assert.Equal(t, expectedTotal, total)
	})

	t.Run("Report does not exist", func(t *testing.T) {
		total, err := testStorage.GetExpensesTotalForAccount(ctx, mock_models.MonthlyReportID(), envelope1.ID)
		require.NoError(t, err)
		assert.Equal(t, models.NewAmount(), total)
	})
}

func TestStorage_DeleteExpense(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category1 := whenSomeCategoryExists(t, ctx, budget.ID, envelope.ID)
	category2 := whenSomeCategoryExists(t, ctx, budget.ID, envelope.ID)
	account := whenSomeAccountExists(t, ctx, budget.ID)
	expense1 := whenSomeExpenseExists(t, ctx, account.ID, category1.ID, category2.ID, report)
	expense2 := whenSomeExpenseExists(t, ctx, account.ID, category1.ID, category2.ID, report)

	t.Run("Success", func(t *testing.T) {
		deleted, err := testStorage.DeleteExpense(ctx, report.ID, expense1.ID)
		require.NoError(t, err)
		assert.Equal(t, expense1, deleted)

		expenses, err := testStorage.GetExpenses(ctx, report.ID)
		require.NoError(t, err)
		assert.ElementsMatch(t, []*models.Expense{expense2}, expenses)
	})

}
