package storage_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBudgets_Create(t *testing.T) {
	ctx := before(t)

	name := "create-budget"

	budget, err := testStorage.Budgets().Create(ctx, name)
	require.NoError(t, err)

	assert.Equal(t, name, budget.Name)
	assert.Empty(t, budget.Expenses)
}

func TestBudgets_Delete(t *testing.T) {
	ctx, budget, _ := beforeWithBudget(t)

	deleted, err := testStorage.Budgets().Delete(ctx, budget.ID)
	require.NoError(t, err)

	assert.Equal(t, budget, deleted)
}

func TestBudgets_FindByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	found, err := testStorage.Budgets().FindByID(ctx, budget.ID)
	require.NoError(t, err)

	assert.Equal(t, budget, found)
}

func TestBudgets_RecreateFails(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	budget, err := testStorage.Budgets().Create(ctx, budget.Name)
	require.Error(t, err)
}
