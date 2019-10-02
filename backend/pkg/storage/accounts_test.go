package storage_test

import (
	"testing"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStorage_CreateAccount(t *testing.T) {
	ctx := before(t)
	input := &models.AccountInput{Name: "test-account"}
	budget := whenSomeBudgetExists(t, ctx)

	account, err := testStorage.CreateAccount(ctx, budget.ID, input)
	require.NoError(t, err)
	assert.Equal(t, input.Name, account.Name)
	assert.Equal(t, models.Amount{0, 0}, account.Balance)
	assert.Equal(t, budget.ID, account.BudgetID)
	assert.NotEqual(t, primitive.ObjectID{}, account.ID)
}

func TestStorage_CreateAccount_DuplicateName(t *testing.T) {
	ctx := before(t)
	input := &models.AccountInput{Name: "test-account"}
	budget := whenSomeBudgetExists(t, ctx)

	_, err := testStorage.CreateAccount(ctx, budget.ID, input)
	require.NoError(t, err)

	_, err = testStorage.CreateAccount(ctx, budget.ID, input)
	require.EqualError(t, err, storage.ErrAccountAlreadyExists.Error())
}

func TestStorage_CreateAccount_NoBudget(t *testing.T) {
	ctx := before(t)
	input := &models.AccountInput{Name: "test-account"}

	_, err := testStorage.CreateAccount(ctx, primitive.NewObjectID(), input)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestStorage_GetAccount(t *testing.T) {
	ctx := before(t)
	input := &models.AccountInput{Name: "test-account"}
	budget := whenSomeBudgetExists(t, ctx)

	created, err := testStorage.CreateAccount(ctx, budget.ID, input)
	require.NoError(t, err)

	account, err := testStorage.GetAccount(ctx, budget.ID, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.Name, account.Name)
	assert.Equal(t, created.Balance, account.Balance)
	assert.Equal(t, budget.ID, account.BudgetID)
	assert.Equal(t, created.ID, account.ID)
}

func TestStorage_GetAccount_NotFound(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)

	account, err := testStorage.GetAccount(ctx, budget.ID, primitive.NewObjectID())
	require.NoError(t, err)
	assert.Nil(t, account)
}

func TestStorage_GetAccount_NoBudget(t *testing.T) {
	ctx := before(t)

	_, err := testStorage.GetAccount(ctx, primitive.NewObjectID(), primitive.NewObjectID())
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestStorage_UpdateAccount(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)
	input := &models.AccountInput{Name: "test-account"}

	created, err := testStorage.CreateAccount(ctx, budget.ID, input)
	require.NoError(t, err)

	changes := models.Changes{"name": "new-name"}
	updated, err := testStorage.UpdateAccount(ctx, budget.ID, created.ID, changes)
	require.NoError(t, err)
	assert.Equal(t, changes["name"], updated.Name)
	assert.Equal(t, created.Balance, updated.Balance)
	assert.Equal(t, budget.ID, updated.BudgetID)
	assert.Equal(t, created.ID, updated.ID)
}

func TestStorage_UpdateAccount_NotFound(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)

	changes := models.Changes{"name": "new-name"}
	_, err := testStorage.UpdateAccount(ctx, budget.ID, primitive.NewObjectID(), changes)
	assert.EqualError(t, err, storage.ErrAccountDoesNotExists.Error())
}
