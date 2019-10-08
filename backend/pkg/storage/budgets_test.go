package storage_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	"github.com/sjanota/budget/backend/pkg/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_CreateBudget(t *testing.T) {
	ctx := before()
	currentMonth := mock_models.Month()
	name := *mock_models.Name()

	budget, err := testStorage.CreateBudget(ctx, name, currentMonth)
	require.NoError(t, err)
	assert.NotEqual(t, primitive.ObjectID{}, budget.ID)
	assert.Equal(t, &models.Budget{
		ID:           budget.ID,
		Name:         name,
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
