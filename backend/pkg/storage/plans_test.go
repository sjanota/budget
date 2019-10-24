package storage_test

import (
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage_DeletePlan(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)
	envelope1 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	envelope2 := whenSomeEnvelopeExists(t, ctx, budget.ID)
	plan1 := whenSomePlanExists(t, ctx, envelope1.ID, envelope2.ID, report)
	plan2 := whenSomePlanExists(t, ctx, envelope1.ID, envelope2.ID, report)

	t.Run("Success", func(t *testing.T) {
		deleted, err := testStorage.DeletePlan(ctx, report.ID, plan1.ID)
		require.NoError(t, err)
		assert.Equal(t, plan1, deleted)

		report, err := testStorage.GetMonthlyReport(ctx, report.ID)
		require.NoError(t, err)
		assert.ElementsMatch(t, []*models.Plan{plan2}, report.Plans)
	})

}
