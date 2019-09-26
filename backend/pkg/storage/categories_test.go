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

func TestCategories_Insert(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := category1
	inserted, err := testStorage.Categories(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.Equal(t, expected, inserted)
}

func TestCategories_Insert_BudgetNotExist(t *testing.T) {
	ctx := context.Background()

	in := category1
	_, err := testStorage.Categories(primitive.NewObjectID()).Insert(ctx, *in)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestCategories_FindByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := category1
	inserted, err := testStorage.Categories(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Categories(budget.ID).FindByID(ctx, inserted.ID)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, found)
	assert.Equal(t, expected, found)
}

func TestCategories_FindByID_NotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := category1
	_, err := testStorage.Categories(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Categories(budget.ID).FindByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestCategories_ReplaceByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := category1
	inserted, err := testStorage.Categories(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	in = category2
	replaced, err := testStorage.Categories(budget.ID).ReplaceByID(ctx, inserted.ID, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, replaced)
	assert.Equal(t, expected, replaced)
}

func TestCategories_ReplaceByID_NotExist(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := category2
	replaced, err := testStorage.Categories(budget.ID).ReplaceByID(ctx, primitive.NewObjectID(), *in)
	require.NoError(t, err)
	require.Nil(t, replaced)
}

func TestCategories_FindAll(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	inserted1, err := testStorage.Categories(budget.ID).Insert(ctx, *category1)
	require.NoError(t, err)
	inserted2, err := testStorage.Categories(budget.ID).Insert(ctx, *category2)
	require.NoError(t, err)

	found, err := testStorage.Categories(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 2)
	assert.Contains(t, found, inserted1)
	assert.Contains(t, found, inserted2)
}

func TestCategories_FindAll_None(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	found, err := testStorage.Categories(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 0)
}

var category1 = &models.CategoryInput{
	Name: "category1",
	EnvelopeID: primitive.NewObjectID(),
}

var category2 = &models.CategoryInput{
	Name: "category2",
	EnvelopeID: primitive.NewObjectID(),
}
