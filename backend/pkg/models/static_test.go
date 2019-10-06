package models_test

import (
	"testing"

	mock_models "github.com/sjanota/budget/backend/pkg/models/mocks"
	"github.com/stretchr/testify/assert"
)

func TestExpense_TotalAmount(t *testing.T) {
	category1 := mock_models.ExpenseCategory()
	category2 := mock_models.ExpenseCategory()
	expense := mock_models.Expense().WithCategories(category1, category2)
	expectedTotal := category1.Amount.Add(category2.Amount)

	assert.Equal(t, expectedTotal, expense.TotalAmount())
}
