package resolver

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sjanota/budget/backend/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCategoryResolver_Envelope(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	storage := mocks.NewMockCategoryResolverStorage(mock)
	resolver := &categoryResolver{storage}
	ctx := context.TODO()
	budget := mocks.Budget()
	envelope := mocks.Envelope()
	category := mocks.Category().WithBudget(budget.ID).WithEnvelope(envelope.ID)

	t.Run("Success", func(t *testing.T) {
		storage.EXPECT().
			GetEnvelope(gomock.Eq(ctx), gomock.Eq(budget.ID), gomock.Eq(envelope.ID)).
			Return(envelope, nil)

		actualEnvelope, err := resolver.Envelope(ctx, &category)
		require.NoError(t, err)
		assert.True(t, actualEnvelope == envelope)
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("test error")
		storage.EXPECT().
			GetEnvelope(gomock.Eq(ctx), gomock.Eq(budget.ID), gomock.Eq(envelope.ID)).
			Return(nil, err)

		_, actualErr := resolver.Envelope(ctx, &category)
		require.Equal(t, err, actualErr)
	})
}
