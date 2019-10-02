package storage_test

import (
	"testing"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStorage_CreateEnvelope(t *testing.T) {
	ctx := before(t)
	input := &models.EnvelopeInput{Name: "test-envelope", Limit: models.Amount{12, 36}}
	budget := whenSomeBudgetExists(t, ctx)

	created, err := testStorage.CreateEnvelope(ctx, budget.ID, input)
	require.NoError(t, err)
	assert.Equal(t, input.Name, created.Name)
	assert.Equal(t, models.Amount{0, 0}, created.Balance)
	assert.Equal(t, input.Limit, created.Limit)
	assert.Equal(t, budget.ID, created.BudgetID)
	assert.NotEqual(t, primitive.ObjectID{}, created.ID)
}

func TestStorage_CreateEnvelope_DuplicateName(t *testing.T) {
	ctx := before(t)
	input := &models.EnvelopeInput{Name: "test-envelope", Limit: models.Amount{12, 36}}
	budget := whenSomeBudgetExists(t, ctx)

	_, err := testStorage.CreateEnvelope(ctx, budget.ID, input)
	require.NoError(t, err)

	_, err = testStorage.CreateEnvelope(ctx, budget.ID, input)
	require.EqualError(t, err, storage.ErrAlreadyExists.Error())
}

func TestStorage_CreateEnvelope_NoBudget(t *testing.T) {
	ctx := before(t)
	input := &models.EnvelopeInput{Name: "test-envelope", Limit: models.Amount{12, 36}}

	_, err := testStorage.CreateEnvelope(ctx, primitive.NewObjectID(), input)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestStorage_GetEnvelope(t *testing.T) {
	ctx := before(t)
	input := &models.EnvelopeInput{Name: "test-envelope", Limit: models.Amount{12, 36}}
	budget := whenSomeBudgetExists(t, ctx)

	created, err := testStorage.CreateEnvelope(ctx, budget.ID, input)
	require.NoError(t, err)

	envelope, err := testStorage.GetEnvelope(ctx, budget.ID, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.Name, envelope.Name)
	assert.Equal(t, created.Balance, envelope.Balance)
	assert.Equal(t, created.Limit, envelope.Limit)
	assert.Equal(t, created.BudgetID, envelope.BudgetID)
	assert.Equal(t, created.ID, envelope.ID)
}

func TestStorage_GetEnvelope_NotFound(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)

	envelope, err := testStorage.GetEnvelope(ctx, budget.ID, primitive.NewObjectID())
	require.NoError(t, err)
	assert.Nil(t, envelope)
}

func TestStorage_GetEnvelope_NoBudget(t *testing.T) {
	ctx := before(t)

	_, err := testStorage.GetEnvelope(ctx, primitive.NewObjectID(), primitive.NewObjectID())
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}
