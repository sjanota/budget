package resolver

import (
	"context"
	"testing"
	"time"

	storage "github.com/sjanota/budget/backend/pkg/storage"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	. "github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMutationResolver_CreateBudget(t *testing.T) {
	now := time.Now()
	ctx := context.TODO()
	testBudget := mock_models.Budget()
	testMonth := models.Month{Month: now.Month(), Year: now.Year()}
	testErr := errors.New("test error")

	t.Run("Success", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.CreateBudget(Eq(ctx), Eq(testMonth)).Return(testBudget, nil)

		budget, err := resolver.Mutation().CreateBudget(ctx)
		require.NoError(t, err)
		assert.Equal(t, budget, budget)
	})

	t.Run("Error", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.CreateBudget(Eq(ctx), Eq(testMonth)).Return(nil, testErr)

		_, err := resolver.Mutation().CreateBudget(ctx)
		require.Equal(t, testErr, err)
	})
}

func TestMutationResolver_CreateExpense(t *testing.T) {
	ctx := context.TODO()
	testBudget := mock_models.Budget()
	testInput := mock_models.ExpenseInput()
	testExpense := mock_models.Expense()
	testReport := mock_models.MonthlyReport()
	testErr := errors.New("test error")

	t.Run("GetBudget error", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(nil, testErr)

		_, err := resolver.Mutation().CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})

	t.Run("Report exists", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(testExpense, nil)

		created, err := resolver.Mutation().CreateExpense(ctx, testBudget.ID, *testInput)
		require.NoError(t, err)
		require.Equal(t, testExpense, created)
	})

	t.Run("Report does not exist and is created", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(testExpense, nil)
		storageExpect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(testReport, nil)

		created, err := resolver.Mutation().CreateExpense(ctx, testBudget.ID, *testInput)
		require.NoError(t, err)
		require.Equal(t, testExpense, created)
	})

	t.Run("Other create expense error", func(t *testing.T) {
		resolver, storageExpect, after := before(t)

		defer after()

		storageExpect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, testErr)

		_, err := resolver.Mutation().CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})

	t.Run("Other create expense error on second attempt", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, testErr)
		storageExpect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(testReport, nil)

		_, err := resolver.Mutation().CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})

	t.Run("Report is created testInput the meanwhile", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(testExpense, nil)
		storageExpect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(nil, storage.ErrAlreadyExists)

		created, err := resolver.Mutation().CreateExpense(ctx, testBudget.ID, *testInput)
		require.NoError(t, err)
		require.Equal(t, testExpense, created)
	})

	t.Run("CreateMonthlyReport error", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		storageExpect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		storageExpect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(nil, testErr)

		_, err := resolver.Mutation().CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})
}
