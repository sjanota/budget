package resolver

import (
	"context"
	"errors"
	"testing"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	"github.com/stretchr/testify/assert"

	. "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBudgetResolver_CurrentMonth(t *testing.T) {
	ctx := context.TODO()
	testBudget := mock_models.Budget()
	testReport := mock_models.MonthlyReport().WithBudget(*testBudget)
	testErr := errors.New("test error")

	t.Run("Success", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetMonthlyReport(Eq(ctx), Eq(testReport.ID)).Return(&testReport, nil)

		actualReport, err := resolver.Budget().CurrentMonth(ctx, testBudget)
		require.NoError(t, err)
		assert.True(t, actualReport == &testReport)
	})

	t.Run("Error", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetMonthlyReport(Eq(ctx), Eq(testReport.ID)).Return(nil, testErr)

		_, err := resolver.Budget().CurrentMonth(ctx, testBudget)
		require.Equal(t, testErr, err)
	})
}
