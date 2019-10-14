package resolver

import (
	"testing"

	mock_resolver "github.com/sjanota/budget/backend/pkg/resolver/mocks"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"

	. "github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCategoryResolver_Envelope(t *testing.T) {
	testBudget := mock_models.Budget()
	ctx := mock_resolver.MockContext(testBudget.ID)
	testEnvelope := mock_models.Envelope()
	testCategory := mock_models.Category().WithEnvelope(testEnvelope.ID)
	testErr := errors.New("test error")

	t.Run("Success", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetEnvelope(Eq(ctx), Eq(testBudget.ID), Eq(testEnvelope.ID)).Return(testEnvelope, nil)

		envelope, err := resolver.Category().Envelope(ctx, &testCategory)
		require.NoError(t, err)
		assert.Equal(t, testEnvelope, envelope)
	})

	t.Run("Error", func(t *testing.T) {
		resolver, storageExpect, after := before(t)
		defer after()

		storageExpect.GetEnvelope(Eq(ctx), Eq(testBudget.ID), Eq(testEnvelope.ID)).Return(nil, testErr)

		_, err := resolver.Category().Envelope(ctx, &testCategory)
		require.Equal(t, testErr, err)
	})
}
