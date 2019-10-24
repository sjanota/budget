package storage_test

import (
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage_DeleteTransfer(t *testing.T) {
	ctx := before()
	budget := whenSomeBudgetExists(t, ctx)
	report := whenSomeMonthlyReportExists(t, ctx, budget.ID)
	account1 := whenSomeAccountExists(t, ctx, budget.ID)
	account2 := whenSomeAccountExists(t, ctx, budget.ID)
	transfer1 := whenSomeTransferExists(t, ctx, account1.ID, account2.ID, report)
	transfer2 := whenSomeTransferExists(t, ctx, account1.ID, account2.ID, report)

	t.Run("Success", func(t *testing.T) {
		deleted, err := testStorage.DeleteTransfer(ctx, report.ID, transfer1.ID)
		require.NoError(t, err)
		assert.Equal(t, transfer1, deleted)

		report, err := testStorage.GetMonthlyReport(ctx, report.ID)
		require.NoError(t, err)
		assert.ElementsMatch(t, []*models.Transfer{transfer2}, report.Transfers)
	})

}
