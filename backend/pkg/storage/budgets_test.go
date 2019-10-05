package storage_test

import (
	"testing"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStorage_CreateBudget(t *testing.T) {
	ctx := before()
	id := primitive.NewObjectID()
	currentMonth := mock.Month()

	budget, err := testStorage.CreateBudget(ctx, id, currentMonth)
	require.NoError(t, err)
	assert.Equal(t, &models.Budget{
		ID:           id,
		Accounts:     []*models.Account{},
		Envelopes:    []*models.Envelope{},
		Categories:   []*models.Category{},
		CurrentMonth: currentMonth,
	}, budget)
}

func TestStorage_GetBudget(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)

	got, err := testStorage.GetBudget(ctx, budget.ID)
	require.NoError(t, err)
	assert.Equal(t, budget, got)
}
