package resolver

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/sjanota/budget/backend/pkg/mocks"
	"github.com/stretchr/testify/require"
)

func TestBudgetResolver_CurrentMonth(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	storage := mocks.NewMockBudgetResolverStorage(mock)
	resolver := &budgetResolver{storage}
	ctx := context.TODO()
	budget := mocks.Budget()
	report := mocks.MonthlyReport().WithBudget(*budget)

	t.Run("Success", func(t *testing.T) {
		storage.EXPECT().
			GetMonthlyReport(gomock.Eq(ctx), gomock.Eq(report.ID)).
			Return(&report, nil)

		actualReport, err := resolver.CurrentMonth(ctx, budget)
		require.NoError(t, err)
		assert.True(t, actualReport == &report)
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("test error")
		storage.EXPECT().
			GetMonthlyReport(gomock.Eq(ctx), gomock.Eq(report.ID)).
			Return(nil, err)

		_, actualErr := resolver.CurrentMonth(ctx, budget)
		require.Equal(t, err, actualErr)
	})
}
