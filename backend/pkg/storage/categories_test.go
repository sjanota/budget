package storage_test

import (
	"testing"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gotest.tools/assert"
)

func TestStorage_CreateCategory(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)
	input := &models.CategoryInput{Name: "test-category", EnvelopeID: envelope.ID}

	created, err := testStorage.CreateCategory(ctx, budget.ID, input)
	require.NoError(t, err)
	assert.Equal(t, input.Name, created.Name)
	assert.Equal(t, input.EnvelopeID, created.EnvelopeID)
	assert.Equal(t, budget.ID, created.BudgetID)
}

func TestStorage_CreateCategory_DuplicateName(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)
	envelope := whenSomeEnvelopeExists(t, ctx, budget.ID)
	input := &models.CategoryInput{Name: "test-category", EnvelopeID: envelope.ID}

	_, err := testStorage.CreateCategory(ctx, budget.ID, input)
	require.NoError(t, err)

	_, err = testStorage.CreateCategory(ctx, budget.ID, input)
	require.EqualError(t, err, storage.ErrAlreadyExists.Error())
}

func TestStorage_CreateCategory_EnvelopDoesNotExist(t *testing.T) {
	ctx := before(t)
	budget := whenSomeBudgetExists(t, ctx)
	input := &models.CategoryInput{Name: "test-category", EnvelopeID: primitive.NewObjectID()}

	_, err := testStorage.CreateCategory(ctx, budget.ID, input)
	require.EqualError(t, err, storage.ErrInvalidReference.Error())
}

func TestStorage_CreateCategory_NoBudget(t *testing.T) {
	ctx := before(t)
	input := &models.CategoryInput{Name: "test-category", EnvelopeID: primitive.NewObjectID()}

	_, err := testStorage.CreateCategory(ctx, primitive.NewObjectID(), input)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}
