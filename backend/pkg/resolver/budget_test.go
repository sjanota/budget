package resolver

import (
	"context"
	"errors"
	"testing"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	mock_resolver "github.com/sjanota/budget/backend/pkg/resolver/mocks"

	"github.com/stretchr/testify/assert"

	. "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBudgetResolver_CurrentMonth(t *testing.T) {
	mock := NewController(t)
	defer mock.Finish()

	storage := mock_resolver.NewMockStorage(mock)
	resolver := &budgetResolver{Resolver: &Resolver{Storage: storage}}
	ctx := context.TODO()
	budget := mock_models.Budget()
	report := mock_models.MonthlyReport().WithBudget(*budget)

	t.Run("Success", func(t *testing.T) {
		storage.EXPECT().GetMonthlyReport(Eq(ctx), Eq(report.ID)).Return(&report, nil)

		actualReport, err := resolver.CurrentMonth(ctx, budget)
		require.NoError(t, err)
		assert.True(t, actualReport == &report)
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("test error")
		storage.EXPECT().GetMonthlyReport(Eq(ctx), Eq(report.ID)).Return(nil, err)

		_, actualErr := resolver.CurrentMonth(ctx, budget)
		require.Equal(t, err, actualErr)
	})
}
