package storage_test

import (
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestStorage_CreateAccount(t *testing.T) {
	ctx := before(t)
	input := &models.AccountInput{Name: "test-account"}
	budget := whenSomeBudgetExists(t, ctx)

	account, err := testStorage.CreateAccount(ctx, budget.ID, input)
	require.NoError(t, err)
	assert.Equal(t, input.Name, account.Name)
	assert.Equal(t, models.Amount{0,0}, account.Balance)
	assert.Equal(t, budget.ID, account.BudgetID)
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

	account, err := testStorage.GetAccount(ctx, budget.ID, created.Name)
	require.NoError(t, err)
	assert.Equal(t, created.Name, account.Name)
	assert.Equal(t, created.Balance, account.Balance)
	assert.Equal(t, budget.ID, account.BudgetID)
}

func TestStorage_GetAccount_NotFound(t *testing.T) {
	ctx := before(t)
	input := &models.AccountInput{Name: "test-account"}
	budget := whenSomeBudgetExists(t, ctx)

	account, err := testStorage.GetAccount(ctx, budget.ID, input.Name)
	require.NoError(t, err)
	assert.Nil(t, account)
}