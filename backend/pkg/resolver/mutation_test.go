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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMutationResolver_CreateBudget(t *testing.T) {
	now := time.Date(2020, time.March, 30, 0, 0, 0, 0, time.UTC)
	ctx := context.TODO()
	testBudget := mock_models.Budget()
	testMonth := models.Month{Month: now.Month(), Year: now.Year()}
	testErr := errors.New("test error")

	newResolver := func(baseResolver *Resolver) *mutationResolver {
		return &mutationResolver{
			Resolver:    baseResolver,
			Now:         func() time.Time { return now },
			NewObjectID: func() primitive.ObjectID { return testBudget.ID },
		}
	}

	t.Run("Success", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := newResolver(baseResolver)
		defer after()

		expect.CreateBudget(Eq(ctx), Eq(testMonth)).Return(testBudget, nil)

		budget, err := resolver.CreateBudget(ctx)
		require.NoError(t, err)
		assert.Equal(t, budget, budget)
	})

	t.Run("Error", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := newResolver(baseResolver)
		defer after()

		expect.CreateBudget(Eq(ctx), Eq(testMonth)).Return(nil, testErr)

		_, err := resolver.CreateBudget(ctx)
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
		baseResolver, expect, after := before(t)
		resolver := &mutationResolver{Resolver: baseResolver}
		defer after()

		expect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(nil, testErr)

		_, err := resolver.CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})

	t.Run("Report exists", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := &mutationResolver{Resolver: baseResolver}
		defer after()

		expect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(testExpense, nil)

		created, err := resolver.CreateExpense(ctx, testBudget.ID, *testInput)
		require.NoError(t, err)
		require.Equal(t, testExpense, created)
	})

	t.Run("Report does not exist and is created", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := &mutationResolver{Resolver: baseResolver}
		defer after()

		expect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(testExpense, nil)
		expect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(testReport, nil)

		created, err := resolver.CreateExpense(ctx, testBudget.ID, *testInput)
		require.NoError(t, err)
		require.Equal(t, testExpense, created)
	})

	t.Run("Other create expense error", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := &mutationResolver{Resolver: baseResolver}
		defer after()

		expect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, testErr)

		_, err := resolver.CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})

	t.Run("Other create expense error on second attempt", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := &mutationResolver{Resolver: baseResolver}
		defer after()

		expect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, testErr)
		expect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(testReport, nil)

		_, err := resolver.CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})

	t.Run("Report is created testInput the meanwhile", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := &mutationResolver{Resolver: baseResolver}
		defer after()

		expect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(testExpense, nil)
		expect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(nil, storage.ErrAlreadyExists)

		created, err := resolver.CreateExpense(ctx, testBudget.ID, *testInput)
		require.NoError(t, err)
		require.Equal(t, testExpense, created)
	})

	t.Run("CreateMonthlyReport error", func(t *testing.T) {
		baseResolver, expect, after := before(t)
		resolver := &mutationResolver{Resolver: baseResolver}
		defer after()

		expect.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expect.CreateExpense(Eq(ctx), Eq(testBudget.CurrentMonthID()), Eq(testInput)).Return(nil, storage.ErrNoReport)
		expect.CreateMonthlyReport(Eq(ctx), Eq(testBudget.ID), Eq(testBudget.CurrentMonth)).Return(nil, testErr)

		_, err := resolver.CreateExpense(ctx, testBudget.ID, *testInput)
		require.Equal(t, testErr, err)
	})
}
