package resolver

import (
	"testing"

	mock_resolver "github.com/sjanota/budget/backend/pkg/resolver/mocks"

	. "github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sjanota/budget/backend/pkg/models"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountResolver_Balance(t *testing.T) {
	testBudget := mock_models.Budget()
	ctx := mock_resolver.MockContext(testBudget.ID)
	testAccount := mock_models.Account()
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
		expectStorage.GetExpensesTotalForAccount(Eq(ctx), Eq(expectedReportID), Eq(testAccount.ID)).Return(testAmount, nil)

		balance, err := resolver.Account().Balance(ctx, testAccount)
		require.NoError(t, err)
		assert.Equal(t, testAmount, balance)
	})

	t.Run("GetBudget error", func(t *testing.T) {
		resolver, expectStorage, after := before(t)
		defer after()

		expectStorage.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(nil, testErr)

		_, err := resolver.Account().Balance(ctx, testAccount)
		require.Equal(t, testErr, err)
	})

	t.Run("GetExpensesTotalForAccount error", func(t *testing.T) {
		resolver, expectStorage, after := before(t)
		defer after()

		expectStorage.GetBudget(Eq(ctx), Eq(testBudget.ID)).Return(testBudget, nil)
		expectStorage.GetExpensesTotalForAccount(Eq(ctx), Eq(expectedReportID), Eq(testAccount.ID)).Return(nil, testErr)

		_, err := resolver.Account().Balance(ctx, testAccount)
		require.Equal(t, testErr, err)
	})
}
