package storage_test

import (
	"testing"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStorage_CreateAccount(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)

	t.Run("Success", func(t *testing.T) {
		input := &models.AccountInput{Name: *mock_models.Name()}

		account, err := testStorage.CreateAccount(ctx, budget.ID, input)
		require.NoError(t, err)
		assert.NotEqual(t, primitive.ObjectID{}, account.ID)
		assert.Equal(t, &models.Account{
			ID:       account.ID,
			Name:     input.Name,
			Balance:  models.Amount{},
			BudgetID: budget.ID,
		}, account)
	})

	t.Run("DuplicateName", func(t *testing.T) {
		input := &models.AccountInput{Name: *mock_models.Name()}

		_, err := testStorage.CreateAccount(ctx, budget.ID, input)
		require.NoError(t, err)

		_, err = testStorage.CreateAccount(ctx, budget.ID, input)
		require.EqualError(t, err, storage.ErrAlreadyExists.Error())
	})

	t.Run("NoBudget", func(t *testing.T) {
		input := &models.AccountInput{Name: *mock_models.Name()}

		_, err := testStorage.CreateAccount(ctx, primitive.NewObjectID(), input)
		require.EqualError(t, err, storage.ErrNoBudget.Error())
	})
}

func TestStorage_GetAccount(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	created := whenSomeAccountExists(t, ctx, budget.ID)

	t.Run("Success", func(t *testing.T) {
		account, err := testStorage.GetAccount(ctx, budget.ID, created.ID)
		require.NoError(t, err)
		assert.Equal(t, created, account)
	})

	t.Run("NotFound", func(t *testing.T) {
		account, err := testStorage.GetAccount(ctx, budget.ID, primitive.NewObjectID())
		require.NoError(t, err)
		assert.Nil(t, account)
	})

	t.Run("NoBudget", func(t *testing.T) {
		_, err := testStorage.GetAccount(ctx, primitive.NewObjectID(), primitive.NewObjectID())
		require.EqualError(t, err, storage.ErrNoBudget.Error())
	})
}

func TestStorage_UpdateAccount(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	account := whenSomeAccountExists(t, ctx, budget.ID)

	t.Run("Success", func(t *testing.T) {
		changes := models.Changes{"name": *mock_models.Name()}
		updated, err := testStorage.UpdateAccount(ctx, budget.ID, account.ID, changes)
		require.NoError(t, err)
		assert.Equal(t, &models.Account{
			ID:       account.ID,
			Name:     changes["name"].(string),
			Balance:  account.Balance,
			BudgetID: budget.ID,
		}, updated)
	})

	t.Run("NotFound", func(t *testing.T) {
		changes := models.Changes{"name": *mock_models.Name()}
		_, err := testStorage.UpdateAccount(ctx, budget.ID, primitive.NewObjectID(), changes)
		assert.EqualError(t, err, storage.ErrDoesNotExists.Error())
	})
}
