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

func TestEnvelopes_Insert(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := envelope1
	inserted, err := testStorage.Envelopes(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.Equal(t, expected, inserted)
}

func TestEnvelopes_Insert_BudgetNotExist(t *testing.T) {
	ctx := context.Background()

	in := envelope1
	_, err := testStorage.Envelopes(primitive.NewObjectID()).Insert(ctx, *in)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestEnvelopes_FindByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := envelope1
	inserted, err := testStorage.Envelopes(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Envelopes(budget.ID).FindByID(ctx, inserted.ID)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, found)
	assert.Equal(t, expected, found)
}

func TestEnvelopes_FindByID_NotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := envelope1
	_, err := testStorage.Envelopes(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Envelopes(budget.ID).FindByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestEnvelopes_ReplaceByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := envelope1
	inserted, err := testStorage.Envelopes(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	in = envelope2
	replaced, err := testStorage.Envelopes(budget.ID).ReplaceByID(ctx, inserted.ID, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, replaced)
	assert.Equal(t, expected, replaced)
}

func TestEnvelopes_ReplaceByID_NotExist(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := envelope2
	replaced, err := testStorage.Envelopes(budget.ID).ReplaceByID(ctx, primitive.NewObjectID(), *in)
	require.NoError(t, err)
	require.Nil(t, replaced)
}

func TestEnvelopes_FindAll(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	inserted1, err := testStorage.Envelopes(budget.ID).Insert(ctx, *envelope1)
	require.NoError(t, err)
	inserted2, err := testStorage.Envelopes(budget.ID).Insert(ctx, *envelope2)
	require.NoError(t, err)

	found, err := testStorage.Envelopes(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 2)
	assert.Contains(t, found, inserted1)
	assert.Contains(t, found, inserted2)
}

func TestEnvelopes_FindAll_None(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	found, err := testStorage.Envelopes(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 0)
}

var envelope1 = &models.EnvelopeInput{
	Name: "envelope1",
}

var envelope2 = &models.EnvelopeInput{
	Name: "envelope2",
}
