package storage_test

import (
	"testing"

	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/sjanota/budget/backend/pkg/storage/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sjanota/budget/backend/pkg/models"
)

func TestStorage_CreateMonthlyReport(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)
	otherBudget := whenSomeBudgetExists(t, ctx)

	t.Run("Success", func(t *testing.T) {
		input := &models.MonthlyReportInput{Month: mock.Month(), Year: mock.Year()}
		report, err := testStorage.CreateMonthlyReport(ctx, budget.ID, input)
		require.NoError(t, err)
		assert.NotEqual(t, primitive.ObjectID{}, report.ID)
		assert.Equal(t, &models.MonthlyReport{
			ID:        report.ID,
			Month:     input.Month,
			Year:      input.Year,
			Expenses:  []*models.Expense{},
			Transfers: []*models.Transfer{},
			Plans:     []*models.Plan{},
			BudgetID:  budget.ID,
		}, report)
	})

	t.Run("Duplicated date", func(t *testing.T) {
		input := &models.MonthlyReportInput{Month: mock.Month(), Year: mock.Year()}
		_, err := testStorage.CreateMonthlyReport(ctx, budget.ID, input)
		require.NoError(t, err)

		_, err = testStorage.CreateMonthlyReport(ctx, budget.ID, input)
		require.EqualError(t, err, storage.ErrAlreadyExists.Error())
	})

	t.Run("Duplicated date on different budget", func(t *testing.T) {
		input := &models.MonthlyReportInput{Month: mock.Month(), Year: mock.Year()}
		_, err := testStorage.CreateMonthlyReport(ctx, budget.ID, input)
		require.NoError(t, err)

		_, err = testStorage.CreateMonthlyReport(ctx, otherBudget.ID, input)
		require.NoError(t, err)
	})
}
