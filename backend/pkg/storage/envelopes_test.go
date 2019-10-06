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

func TestStorage_CreateEnvelope(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)

	t.Run("Success", func(t *testing.T) {
		input := &models.EnvelopeInput{Name: *mock_models.Name(), Limit: mock_models.Amount()}

		created, err := testStorage.CreateEnvelope(ctx, budget.ID, input)
		require.NoError(t, err)
		assert.NotEqual(t, primitive.ObjectID{}, created.ID)
		assert.Equal(t, &models.Envelope{
			ID:       created.ID,
			Name:     input.Name,
			Limit:    input.Limit,
			Balance:  models.Amount{},
			BudgetID: budget.ID,
		}, created)
	})

	t.Run("DuplicatedName", func(t *testing.T) {
		input := &models.EnvelopeInput{Name: *mock_models.Name(), Limit: mock_models.Amount()}
		budget := whenSomeBudgetExists(t, ctx)

		_, err := testStorage.CreateEnvelope(ctx, budget.ID, input)
		require.NoError(t, err)

		_, err = testStorage.CreateEnvelope(ctx, budget.ID, input)
		require.EqualError(t, err, storage.ErrAlreadyExists.Error())
	})

	t.Run("NoBudget", func(t *testing.T) {
		input := &models.EnvelopeInput{Name: *mock_models.Name(), Limit: mock_models.Amount()}

		_, err := testStorage.CreateEnvelope(ctx, primitive.NewObjectID(), input)
		require.EqualError(t, err, storage.ErrNoBudget.Error())
	})

}

func TestStorage_GetEnvelope(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)

	t.Run("Success", func(t *testing.T) {
		got, err := testStorage.GetEnvelope(ctx, budget.ID, envelope.ID)
		require.NoError(t, err)
		assert.Equal(t, envelope, got)
	})

	t.Run("NotFound", func(t *testing.T) {
		got, err := testStorage.GetEnvelope(ctx, budget.ID, primitive.NewObjectID())
		require.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("NoBudget", func(t *testing.T) {
		_, err := testStorage.GetEnvelope(ctx, primitive.NewObjectID(), primitive.NewObjectID())
		require.EqualError(t, err, storage.ErrNoBudget.Error())
	})
}

func TestStorage_UpdateEnvelope(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)

	t.Run("Success", func(t *testing.T) {

		changes := models.Changes{"name": *mock_models.Name(), "limit": mock_models.Amount()}
		updated, err := testStorage.UpdateEnvelope(ctx, budget.ID, envelope.ID, changes)
		require.NoError(t, err)
		assert.Equal(t, &models.Envelope{
			ID:       envelope.ID,
			Name:     changes["name"].(string),
			Limit:    changes["limit"].(*models.Amount),
			Balance:  envelope.Balance,
			BudgetID: budget.ID,
		}, updated)
	})

	t.Run("NotFound", func(t *testing.T) {
		changes := models.Changes{"name": *mock_models.Name()}
		_, err := testStorage.UpdateAccount(ctx, budget.ID, primitive.NewObjectID(), changes)
		assert.EqualError(t, err, storage.ErrDoesNotExists.Error())
	})
}
