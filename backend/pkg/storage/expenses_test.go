package storage_test

import (
	"context"
	"testing"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExpenses_Insert(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.Equal(t, expected, inserted)
	assertEvent(t, watch, models.EventTypeCreated, expected)

	all, err := testStorage.Expenses(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	assert.Len(t, all, 1)
}

func TestExpenses_Insert_BudgetNotExist(t *testing.T) {
	ctx := context.Background()

	in := expenseInput1
	_, err := testStorage.Expenses(primitive.NewObjectID()).Insert(ctx, *in)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestExpenses_Watch_BudgetNotExist(t *testing.T) {
	ctx := context.Background()

	_, err := testStorage.Expenses(primitive.NewObjectID()).Watch(ctx)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestExpenses_DeleteByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	deleted, err := testStorage.Expenses(budget.ID).DeleteByID(ctx, inserted.ID)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.Equal(t, expected, deleted)
	assertEvent(t, watch, models.EventTypeDeleted, expected)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, deleted.ID)
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestExpenses_DeleteByID_NotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	deleted, err := testStorage.Expenses(budget.ID).DeleteByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	assert.Nil(t, deleted)
	assertNoEvent(t, watch)
}

func TestExpenses_FindByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, inserted.ID)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, found)
	assert.Equal(t, expected, inserted)
}

func TestExpenses_FindByID_NotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	_, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestExpenses_ReplaceByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	in = expenseInput2
	replaced, err := testStorage.Expenses(budget.ID).ReplaceByID(ctx, inserted.ID, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, replaced)
	assert.Equal(t, expected, replaced)
	assertEvent(t, watch, models.EventTypeUpdated, replaced)
}

func TestExpenses_ReplaceByID_NotExist(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput2
	replaced, err := testStorage.Expenses(budget.ID).ReplaceByID(ctx, primitive.NewObjectID(), *in)
	require.NoError(t, err)
	require.Nil(t, replaced)
}

var expenseInput1 = &models.ExpenseInput{
	Title:    "title",
	Location: strPtr("location"),
	Entries: []*models.ExpenseEntryInput{
		{
			Title:      "food",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 28,
			},
		},
		{
			Title:      "sweets",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 28,
			},
		},
	},
	TotalBalance: &models.MoneyAmountInput{
		Integer: 3,
		Decimal: 28,
	},
	Date:      strPtr("12032019"),
	AccountID: idPtr(primitive.NewObjectID()),
}

var expenseInput2 = &models.ExpenseInput{
	Title:    "title",
	Location: strPtr("location"),
	Entries: []*models.ExpenseEntryInput{
		{
			Title:      "food",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 26,
			},
		},
		{
			Title:      "cat food",
			CategoryID: primitive.ObjectID{},
			Balance: &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 26,
			},
		},
		{
			Title:      "medicine",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 28,
			},
		},
	},
	TotalBalance: &models.MoneyAmountInput{
		Integer: 3,
		Decimal: 28,
	},
	Date:      strPtr("12032019"),
	AccountID: idPtr(primitive.NewObjectID()),
}

func assertEvent(t *testing.T, ch <-chan *models.ExpenseEvent, eventType models.EventType, expense *models.Expense) {
	select {
	case event := <-ch:
		assert.Equal(t, &models.ExpenseEvent{Type: eventType, Expense: expense}, event)
	default:
		t.Error("no event received")
	}
}

func assertNoEvent(t *testing.T, ch <-chan *models.ExpenseEvent) {
	select {
	case event := <-ch:
		t.Errorf("unexpected event: %#v", event)
	default:
	}
}
