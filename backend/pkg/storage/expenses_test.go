package storage_test

import (
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestExpenses_Insert(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)
	assert.Equal(t, in.ToModel(budget.ID).WithID(inserted.ID), inserted)

	all, err := testStorage.Expenses(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	assert.Len(t, all, 1)
}

func TestExpenses_Delete(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	deleted, err := testStorage.Expenses(budget.ID).DeleteByID(ctx, inserted.ID)
	require.NoError(t, err)
	assert.Equal(t, in.ToModel(budget.ID).WithID(inserted.ID), deleted)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, deleted.ID)
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestExpenses_FindOne(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, inserted.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, in.ToModel(budget.ID).WithID(inserted.ID), inserted)
}

func TestExpenses_FindOneNotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	_, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	require.Nil(t, found)
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