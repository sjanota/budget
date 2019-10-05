package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStorage_CreateBudget(t *testing.T) {
	ctx := before()
	id := primitive.NewObjectID()
	reportID := primitive.NewObjectID()

	budget, err := testStorage.CreateBudget(ctx, id, reportID)
	require.NoError(t, err)
	assert.Equal(t, id, budget.ID)
	assert.Equal(t, reportID, budget.CurrentMonthID)
	assert.Empty(t, budget.Accounts)
	assert.Empty(t, budget.Envelopes)
	assert.Empty(t, budget.Categories)
}

func TestStorage_GetBudget(t *testing.T) {
	ctx := before()
	created := whenSomeBudgetExists(t, ctx)

	budget, err := testStorage.GetBudget(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, budget.ID)
	assert.Equal(t, created.CurrentMonthID, budget.CurrentMonthID)
	assert.Empty(t, created.Accounts, budget.Accounts)
	assert.Empty(t, created.Envelopes, budget.Envelopes)
	assert.Empty(t, created.Categories, budget.Categories)
}
