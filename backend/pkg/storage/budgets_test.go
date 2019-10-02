package storage_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestStorage_CreateBudget(t *testing.T) {
	ctx := before(t)

	budget, err := testStorage.CreateBudget(ctx)
	require.NoError(t, err)
	assert.NotEqual(t, primitive.ObjectID{}, budget.ID)
	assert.Empty(t, budget.Accounts)
	assert.Empty(t, budget.Envelopes)
	assert.Empty(t, budget.Categories)
}

func TestStorage_GetBudget(t *testing.T) {
	ctx := before(t)

	inserted, err := testStorage.CreateBudget(ctx)
	require.NoError(t, err)

	budget, err := testStorage.GetBudget(ctx, inserted.ID)
	require.NoError(t, err)
	assert.NotEqual(t, primitive.ObjectID{}, budget.ID)
	assert.Empty(t, budget.Accounts)
	assert.Empty(t, budget.Envelopes)
	assert.Empty(t, budget.Categories)
}