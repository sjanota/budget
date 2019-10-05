package resolver

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sjanota/budget/backend/pkg/mocks"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestMutationResolver_CreateBudget(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	storage := mocks.NewMockMutationResolverStorage(mock)
	now := time.Date(2020, time.March, 30, 0, 0, 0, 0, time.UTC)
	budgetID := primitive.NewObjectID()
	resolver := &mutationResolver{
		Storage:     storage,
		Now:         func() time.Time { return now },
		NewObjectID: func() primitive.ObjectID { return budgetID },
	}
	ctx := context.TODO()

	t.Run("Success", func(t *testing.T) {
		expectedMonth := models.Month{Month: now.Month(), Year: now.Year()}
		expectedInput := models.MonthlyReportInput{expectedMonth}
		expectedBudget := mocks.Budget()

		storage.EXPECT().
			CreateMonthlyReport(gomock.Eq(ctx), gomock.Eq(budgetID), gomock.Eq(&expectedInput)).
			Return(mocks.MonthlyReport(), nil)

		storage.EXPECT().
			CreateBudget(gomock.Eq(ctx), gomock.Eq(budgetID), gomock.Eq(expectedMonth)).
			Return(expectedBudget, nil)

		budget, err := resolver.CreateBudget(ctx)
		require.NoError(t, err)
		assert.Equal(t, expectedBudget, budget)
	})

	t.Run("CreateMonthlyReport error", func(t *testing.T) {
		expectedMonth := models.Month{Month: now.Month(), Year: now.Year()}
		expectedInput := models.MonthlyReportInput{expectedMonth}
		expectedErr := errors.New("test error")

		storage.EXPECT().
			CreateMonthlyReport(gomock.Eq(ctx), gomock.Eq(budgetID), gomock.Eq(&expectedInput)).
			Return(nil, expectedErr)

		_, err := resolver.CreateBudget(ctx)
		require.Equal(t, expectedErr, err)
	})

	t.Run("CreateBudget error", func(t *testing.T) {
		expectedMonth := models.Month{Month: now.Month(), Year: now.Year()}
		expectedInput := models.MonthlyReportInput{expectedMonth}
		expectedErr := errors.New("test error")

		storage.EXPECT().
			CreateMonthlyReport(gomock.Eq(ctx), gomock.Eq(budgetID), gomock.Eq(&expectedInput)).
			Return(mocks.MonthlyReport(), nil)

		storage.EXPECT().
			CreateBudget(gomock.Eq(ctx), gomock.Eq(budgetID), gomock.Eq(expectedMonth)).
			Return(nil, expectedErr)

		_, err := resolver.CreateBudget(ctx)
		require.Equal(t, expectedErr, err)
	})

}
