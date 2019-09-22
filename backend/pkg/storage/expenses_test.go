package storage_test

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestExpensesRepository_InsertOne(t *testing.T) {
	drop(t)
	ctx := context.Background()

	in := expenseInput(
		"title",
		"location",
		[]*models.ExpenseEntryInput{
			expenseEntryInput("food", primitive.NewObjectID(), &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 28,
			}),
			expenseEntryInput("sweets", primitive.NewObjectID(), &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 28,
			}),
		},
		&models.MoneyAmountInput{
			Integer: 3,
			Decimal: 28,
		},
		"12032019",
		primitive.NewObjectID(),
	)
	inserted, err := testStorage.Expenses().InsertOne(ctx, *in)
	require.NoError(t, err)
	assertExpenseMatchesInput(t, in, inserted)
}

func TestExpensesRepository_FindOne(t *testing.T) {
	drop(t)
	ctx := context.Background()

	in := expenseInput(
		"title",
		"location",
		[]*models.ExpenseEntryInput{
			expenseEntryInput("food", primitive.NewObjectID(), &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 28,
			}),
			expenseEntryInput("sweets", primitive.NewObjectID(), &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 28,
			}),
		},
		&models.MoneyAmountInput{
			Integer: 3,
			Decimal: 28,
		},
		"12032019",
		primitive.NewObjectID(),
	)
	inserted, err := testStorage.Expenses().InsertOne(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Expenses().FindByID(ctx, inserted.ID)
	require.NoError(t, err)
	assertExpenseMatchesInput(t, in, found)
}

func assertExpenseMatchesInput(t assert.TestingT, expected *models.ExpenseInput, actual *models.Expense) {
	assert.NotNil(t, actual.ID)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Location, actual.Location)
	assert.Len(t, actual.Entries, len(expected.Entries))
	for i := range actual.Entries {
		assert.Equal(t, expected.Entries[i].Title, actual.Entries[i].Title)
		assert.Equal(t, expected.Entries[i].CategoryID, actual.Entries[i].CategoryID)
		assert.Equal(t, expected.Entries[i].Amount.Integer, actual.Entries[i].Amount.Integer)
		assert.Equal(t, expected.Entries[i].Amount.Decimal, actual.Entries[i].Amount.Decimal)
	}
	assert.Equal(t, expected.Total.Integer, actual.Total.Integer)
	assert.Equal(t, expected.Total.Decimal, actual.Total.Decimal)
	assert.Equal(t, expected.Date, actual.Date)
	assert.Equal(t, expected.AccountID, actual.AccountID)
}

func expenseInput(title, location string, entries []*models.ExpenseEntryInput, total *models.MoneyAmountInput, date string, accountID primitive.ObjectID) *models.ExpenseInput {
	return &models.ExpenseInput{
		Title:     title,
		Location:  &location,
		Entries:   entries,
		Total:     total,
		Date:      &date,
		AccountID: &accountID,
	}
}

func expenseEntryInput(title string, categoryID primitive.ObjectID, amount *models.MoneyAmountInput) *models.ExpenseEntryInput {
	return &models.ExpenseEntryInput{
		Title:      title,
		CategoryID: categoryID,
		Amount:     amount,
	}
}
