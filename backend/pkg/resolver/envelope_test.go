package resolver

import (
	"context"
	"testing"

	. "github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sjanota/budget/backend/pkg/models"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvelopeResolver_Balance(t *testing.T) {
	ctx := context.TODO()
	testBudget := mock_models.Budget()
	testEnvelope := mock_models.Envelope().WithBudget(testBudget.ID)
	testAmount := mock_models.Amount()
	testErr := errors.New("test error")
	expectedReportID := models.MonthlyReportID{
		Month:    testBudget.CurrentMonth,
		BudgetID: testBudget.ID,
	}

	t.Run("Success", func(t *testing.T) {
		resolver, expectStorage, after := before(t)
		defer after()

		expectStorage.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expectStorage.GetExpensesTotalForEnvelope(Eq(ctx), Eq(expectedReportID), Eq(testEnvelope.ID)).Return(testAmount, nil)

		balance, err := resolver.Envelope().Balance(ctx, testEnvelope)
		require.NoError(t, err)
		assert.Equal(t, testAmount, balance)
	})

	t.Run("GetBudget error", func(t *testing.T) {
		resolver, expectStorage, after := before(t)
		defer after()

		expectStorage.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(nil, testErr)

		_, err := resolver.Envelope().Balance(ctx, testEnvelope)
		require.Equal(t, testErr, err)
	})

	t.Run("GetExpensesTotalForEnvelope error", func(t *testing.T) {
		resolver, expectStorage, after := before(t)
		defer after()

		expectStorage.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expectStorage.GetExpensesTotalForEnvelope(Eq(ctx), Eq(expectedReportID), Eq(testEnvelope.ID)).Return(nil, testErr)

		_, err := resolver.Envelope().Balance(ctx, testEnvelope)
		require.Equal(t, testErr, err)
	})
}