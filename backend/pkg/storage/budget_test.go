package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBudgets_Insert(t *testing.T) {
	ctx := before(t)

	name := "create-budget"

	budget, err := testStorage.Budgets().Insert(ctx, name)
	require.NoError(t, err)

	assert.Equal(t, name, budget.Name)
	assert.Empty(t, budget.Expenses)
}

func TestBudgets_Insert_Unique(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	budget, err := testStorage.Budgets().Insert(ctx, budget.Name)
	require.Error(t, err)
}

func TestBudgets_DeleteByID(t *testing.T) {
	ctx, budget, _ := beforeWithBudget(t)

	deleted, err := testStorage.Budgets().DeleteByID(ctx, budget.ID)
	require.NoError(t, err)

	assert.Equal(t, budget, deleted)
}

func TestBudgets_DeleteByID_NotExist(t *testing.T) {
	ctx := before(t)

	deleted, err := testStorage.Budgets().DeleteByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	assert.Nil(t, deleted)
}

func TestBudgets_FindByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	found, err := testStorage.Budgets().FindByID(ctx, budget.ID)
	require.NoError(t, err)

	assert.Equal(t, budget, found)
}

func TestBudgets_FindByID_NotExist(t *testing.T) {
	ctx := before(t)

	deleted, err := testStorage.Budgets().FindByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	assert.Nil(t, deleted)
}

func TestBudgets_FindAll(t *testing.T) {
	ctx := before(t)
	name1 := "budget 1"
	name2 := "budget 2"

	budget1, err := testStorage.Budgets().Insert(ctx, name1)
	require.NoError(t, err)
	budget2, err := testStorage.Budgets().Insert(ctx, name2)
	require.NoError(t, err)

	budgets, err := testStorage.Budgets().FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, budgets)
	assert.Len(t, budgets, 2)
	assert.Contains(t, budgets, budget1)
	assert.Contains(t, budgets, budget2)
}

func TestBudgets_FindAll_None(t *testing.T) {
	ctx := before(t)

	budgets, err := testStorage.Budgets().FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, budgets)
	assert.Len(t, budgets, 0)
}
