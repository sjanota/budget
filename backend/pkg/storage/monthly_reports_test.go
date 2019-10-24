package storage_test

import (
	"testing"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/sjanota/budget/backend/pkg/models"
)

func TestStorage_CreateMonthlyReport(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	otherBudget := whenSomeBudgetExists(t, ctx)

	t.Run("Success", func(t *testing.T) {
		month := mock_models.Month()
		report, err := testStorage.CreateMonthlyReport(ctx, budget.ID, month, make([]*models.PlanInput, 0))
		require.NoError(t, err)
		assert.NotEqual(t, primitive.ObjectID{}, report.ID)
		assert.Equal(t, &models.MonthlyReport{
			ID:        report.ID,
			Expenses:  []*models.Expense{},
			Transfers: []*models.Transfer{},
			Plans:     []*models.Plan{},
		}, report)
	})

	t.Run("Duplicated date", func(t *testing.T) {
		month := mock_models.Month()
		_, err := testStorage.CreateMonthlyReport(ctx, budget.ID, month, make([]*models.PlanInput, 0))
		require.NoError(t, err)

		_, err = testStorage.CreateMonthlyReport(ctx, budget.ID, month, make([]*models.PlanInput, 0))
		require.EqualError(t, err, storage.ErrAlreadyExists.Error())
	})

	t.Run("Duplicated date on different budget", func(t *testing.T) {
		month := mock_models.Month()
		_, err := testStorage.CreateMonthlyReport(ctx, budget.ID, month, make([]*models.PlanInput, 0))
		require.NoError(t, err)

		_, err = testStorage.CreateMonthlyReport(ctx, otherBudget.ID, month, make([]*models.PlanInput, 0))
		require.NoError(t, err)
	})
}

func TestStorage_GetMonthlyReport(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)

	t.Run("Success", func(t *testing.T) {
		got, err := testStorage.GetMonthlyReport(ctx, report.ID)
		require.NoError(t, err)
		assert.Equal(t, report, got)
	})
}
