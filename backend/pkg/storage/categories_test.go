package storage_test

import (
	"testing"

	"github.com/sjanota/budget/backend/pkg/mocks"

	"github.com/stretchr/testify/assert"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStorage_CreateCategory(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)

	t.Run("Success", func(t *testing.T) {
		input := &models.CategoryInput{Name: *mocks.Name(), EnvelopeID: envelope.ID}

		created, err := testStorage.CreateCategory(ctx, budget.ID, input)
		require.NoError(t, err)
		assert.NotEqual(t, primitive.ObjectID{}, created.ID)
		assert.Equal(t, &models.Category{
			ID:         created.ID,
			Name:       input.Name,
			EnvelopeID: input.EnvelopeID,
			BudgetID:   budget.ID,
		}, created)
	})

	t.Run("DuplicateName", func(t *testing.T) {
		input := &models.CategoryInput{Name: *mocks.Name(), EnvelopeID: envelope.ID}

		_, err := testStorage.CreateCategory(ctx, budget.ID, input)
		require.NoError(t, err)

		_, err = testStorage.CreateCategory(ctx, budget.ID, input)
		require.EqualError(t, err, storage.ErrAlreadyExists.Error())
	})

	t.Run("EnvelopDoesNotExist", func(t *testing.T) {
		input := &models.CategoryInput{Name: *mocks.Name(), EnvelopeID: primitive.NewObjectID()}

		_, err := testStorage.CreateCategory(ctx, budget.ID, input)
		require.EqualError(t, err, storage.ErrInvalidReference.Error())
	})

	t.Run("NoBudget", func(t *testing.T) {
		input := &models.CategoryInput{Name: *mocks.Name(), EnvelopeID: primitive.NewObjectID()}

		_, err := testStorage.CreateCategory(ctx, primitive.NewObjectID(), input)
		require.EqualError(t, err, storage.ErrNoBudget.Error())
	})

}

func TestStorage_GetCategory(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category := whenSomeCategoryExists(t, ctx, budget.ID, envelope.ID)

	t.Run("Success", func(t *testing.T) {
		got, err := testStorage.GetCategory(ctx, budget.ID, category.ID)
		require.NoError(t, err)
		assert.Equal(t, category, got)
	})

	t.Run("NotFound", func(t *testing.T) {
		category, err := testStorage.GetCategory(ctx, budget.ID, primitive.NewObjectID())
		require.NoError(t, err)
		assert.Nil(t, category)
	})

	t.Run("NoBudget", func(t *testing.T) {
		_, err := testStorage.GetCategory(ctx, primitive.NewObjectID(), primitive.NewObjectID())
		require.EqualError(t, err, storage.ErrNoBudget.Error())
	})

}

func TestStorage_UpdateCategory(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)
	otherEnvelope := whenSomeEnvelopeExists(t, ctx, budget.ID)
	category := whenSomeCategoryExists(t, ctx, budget.ID, envelope.ID)

	t.Run("Success", func(t *testing.T) {
		changes := models.Changes{"name": *mocks.Name(), "envelopeID": otherEnvelope.ID}
		updated, err := testStorage.UpdateCategory(ctx, budget.ID, category.ID, changes)
		require.NoError(t, err)
		assert.Equal(t, changes["name"], updated.Name)
		assert.Equal(t, changes["envelopeID"], updated.EnvelopeID)
		assert.Equal(t, budget.ID, updated.BudgetID)
		assert.Equal(t, category.ID, updated.ID)
	})

	t.Run("EnvelopeDoesNotExist", func(t *testing.T) {
		changes := models.Changes{"name": *mocks.Name(), "envelopeID": primitive.NewObjectID()}
		_, err := testStorage.UpdateCategory(ctx, budget.ID, category.ID, changes)
		require.EqualError(t, err, storage.ErrInvalidReference.Error())
	})

	t.Run("NotFound", func(t *testing.T) {
		changes := models.Changes{"name": *mocks.Name(), "envelopeID": otherEnvelope.ID}
		_, err := testStorage.UpdateCategory(ctx, budget.ID, primitive.NewObjectID(), changes)
		assert.EqualError(t, err, storage.ErrDoesNotExists.Error())
	})

}
