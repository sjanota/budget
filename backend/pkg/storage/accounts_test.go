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

func TestAccounts_Insert(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := account1
	inserted, err := testStorage.Accounts(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.Equal(t, expected, inserted)
}

func TestAccounts_Insert_BudgetNotExist(t *testing.T) {
	ctx := context.Background()

	in := account1
	_, err := testStorage.Accounts(primitive.NewObjectID()).Insert(ctx, *in)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestAccounts_FindByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := account1
	inserted, err := testStorage.Accounts(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Accounts(budget.ID).FindByID(ctx, inserted.ID)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, found)
	assert.Equal(t, expected, found)
}

func TestAccounts_FindByID_NotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := account1
	_, err := testStorage.Accounts(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Accounts(budget.ID).FindByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestAccounts_ReplaceByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := account1
	inserted, err := testStorage.Accounts(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	in = account2
	replaced, err := testStorage.Accounts(budget.ID).ReplaceByID(ctx, inserted.ID, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, replaced)
	assert.Equal(t, expected, replaced)
}

func TestAccounts_ReplaceByID_NotExist(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := account2
	replaced, err := testStorage.Accounts(budget.ID).ReplaceByID(ctx, primitive.NewObjectID(), *in)
	require.NoError(t, err)
	require.Nil(t, replaced)
}

func TestAccounts_FindAll(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	inserted1, err := testStorage.Accounts(budget.ID).Insert(ctx, *account1)
	require.NoError(t, err)
	inserted2, err := testStorage.Accounts(budget.ID).Insert(ctx, *account2)
	require.NoError(t, err)

	found, err := testStorage.Accounts(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 2)
	assert.Contains(t, found, inserted1)
	assert.Contains(t, found, inserted2)
}

func TestAccounts_FindAll_None(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	found, err := testStorage.Accounts(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 0)
}

var account1 = &models.AccountInput{
	Name: "account1",
}

var account2 = &models.AccountInput{
	Name: "account2",
}
