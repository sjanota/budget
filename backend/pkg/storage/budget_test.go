package storage_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBudget_CreateBudget(t *testing.T) {
	drop(t)
	ctx := context.Background()
	name := "my-budget"

	budget, err := testStorage.Budgets().CreateBudget(ctx, name)
	require.NoError(t, err)

	assert.Equal(t, name, budget.Name)
	assert.Empty(t, budget.Expenses)
}
