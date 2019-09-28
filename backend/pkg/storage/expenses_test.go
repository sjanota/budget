package storage_test

import (
	"context"
	"testing"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExpenses_Insert(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.Equal(t, expected, inserted)
	assertEvent(t, watch, models.EventTypeCreated, expected)
}

func TestExpenses_Insert_BudgetNotExist(t *testing.T) {
	ctx := context.Background()

	in := expenseInput1
	_, err := testStorage.Expenses(primitive.NewObjectID()).Insert(ctx, *in)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestExpenses_Watch_BudgetNotExist(t *testing.T) {
	ctx := context.Background()

	_, err := testStorage.Expenses(primitive.NewObjectID()).Watch(ctx)
	require.EqualError(t, err, storage.ErrNoBudget.Error())
}

func TestExpenses_DeleteByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	deleted, err := testStorage.Expenses(budget.ID).DeleteByID(ctx, inserted.ID)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.Equal(t, expected, deleted)
	assertEvent(t, watch, models.EventTypeDeleted, expected)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, deleted.ID)
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestExpenses_DeleteByID_NotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	deleted, err := testStorage.Expenses(budget.ID).DeleteByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	assert.Nil(t, deleted)
	assertNoEvent(t, watch)
}

func TestExpenses_FindByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, inserted.ID)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, found)
	assert.Equal(t, expected, found)
}

func TestExpenses_FindByID_NotExists(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	_, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	found, err := testStorage.Expenses(budget.ID).FindByID(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	require.Nil(t, found)
}

func TestExpenses_ReplaceByID(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput1
	inserted, err := testStorage.Expenses(budget.ID).Insert(ctx, *in)
	require.NoError(t, err)

	watch, err := testStorage.Expenses(budget.ID).Watch(ctx)
	require.NoError(t, err)

	in = expenseInput2
	replaced, err := testStorage.Expenses(budget.ID).ReplaceByID(ctx, inserted.ID, *in)
	require.NoError(t, err)

	expected := in.ToModel(budget.ID).WithID(inserted.ID)
	assert.NotNil(t, replaced)
	assert.Equal(t, expected, replaced)
	assertEvent(t, watch, models.EventTypeUpdated, replaced)
}

func TestExpenses_ReplaceByID_NotExist(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	in := expenseInput2
	replaced, err := testStorage.Expenses(budget.ID).ReplaceByID(ctx, primitive.NewObjectID(), *in)
	require.NoError(t, err)
	require.Nil(t, replaced)
}

func TestExpenses_FindAll(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	inserted1, err := testStorage.Expenses(budget.ID).Insert(ctx, *expenseInput1)
	require.NoError(t, err)
	inserted2, err := testStorage.Expenses(budget.ID).Insert(ctx, *expenseInput2)
	require.NoError(t, err)

	found, err := testStorage.Expenses(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 2)
	assert.Contains(t, found, inserted1)
	assert.Contains(t, found, inserted2)
}

func TestExpenses_FindAll_None(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	found, err := testStorage.Expenses(budget.ID).FindAll(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Len(t, found, 0)
}

func TestExpenses_TotalBalanceForAccount(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	_, err := testStorage.Expenses(budget.ID).Insert(ctx, *expenseInput1)
	require.NoError(t, err)
	_, err = testStorage.Expenses(budget.ID).Insert(ctx, *expenseInput2)
	require.NoError(t, err)

	balance, err := testStorage.Expenses(budget.ID).TotalBalanceForAccount(ctx, accountID)
	require.NoError(t, err)
	require.NotNil(t, balance)
	assert.Equal(t, 17, balance.Integer)
	assert.Equal(t, 10, balance.Decimal)
}

func TestExpenses_TotalBalanceForEnvelope(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	envelopeIn := models.EnvelopeInput{Name: "food"}
	envelope, err := testStorage.Envelopes(budget.ID).Insert(ctx, envelopeIn)
	require.NoError(t, err)

	envelopeInOther := models.EnvelopeInput{Name: "travels"}
	envelopeOther, err := testStorage.Envelopes(budget.ID).Insert(ctx, envelopeInOther)
	require.NoError(t, err)

	categoryIn1 := models.CategoryInput{
		Name:       "lunch",
		EnvelopeID: envelope.ID,
	}
	category1, err := testStorage.Categories(budget.ID).Insert(ctx, categoryIn1)
	require.NoError(t, err)
	categoryIn2 := models.CategoryInput{
		Name:       "restaurants",
		EnvelopeID: envelope.ID,
	}
	category2, err := testStorage.Categories(budget.ID).Insert(ctx, categoryIn2)
	require.NoError(t, err)

	categoryInOther := models.CategoryInput{
		Name:       "airplanes",
		EnvelopeID: envelopeOther.ID,
	}
	categoryOther, err := testStorage.Categories(budget.ID).Insert(ctx, categoryInOther)
	require.NoError(t, err)


	_, err = testStorage.Expenses(budget.ID).Insert(ctx, models.ExpenseInput{
		Title:        "burger",
		Entries:      []*models.ExpenseEntryInput{
			{
				CategoryID: category1.ID,
				Balance:    &models.MoneyAmountInput{10,60},
			},
			{
				CategoryID: primitive.NewObjectID(),
				Balance:    &models.MoneyAmountInput{10,10},
			},
		},
	})
	require.NoError(t, err)

	_, err = testStorage.Expenses(budget.ID).Insert(ctx, models.ExpenseInput{
		Title:        "lasagne",
		Entries:      []*models.ExpenseEntryInput{
			{
				CategoryID: category2.ID,
				Balance:    &models.MoneyAmountInput{10,60},
			},
			{
				CategoryID: primitive.NewObjectID(),
				Balance:    &models.MoneyAmountInput{10,10},
			},
		},
	})
	require.NoError(t, err)

	_, err = testStorage.Expenses(budget.ID).Insert(ctx, models.ExpenseInput{
		Title:        "Flight to Canada",
		Entries:      []*models.ExpenseEntryInput{
			{
				CategoryID: categoryOther.ID,
				Balance:    &models.MoneyAmountInput{10,60},
			},
			{
				CategoryID: primitive.NewObjectID(),
				Balance:    &models.MoneyAmountInput{10,10},
			},
		},
	})
	require.NoError(t, err)

	balance, err := testStorage.Expenses(budget.ID).TotalBalanceForEnvelope(ctx, envelope.ID)
	require.NoError(t, err)
	require.NotNil(t, balance)
	assert.Equal(t, 21, balance.Integer)
	assert.Equal(t, 20, balance.Decimal)
}

func TestExpenses_TotalBalanceForAccount_NotFound(t *testing.T) {
	ctx, budget, after := beforeWithBudget(t)
	defer after()

	_, err := testStorage.Expenses(budget.ID).Insert(ctx, *expenseInput1)
	require.NoError(t, err)
	_, err = testStorage.Expenses(budget.ID).Insert(ctx, *expenseInput2)
	require.NoError(t, err)

	balance, err := testStorage.Expenses(budget.ID).TotalBalanceForAccount(ctx, primitive.NewObjectID())
	require.NoError(t, err)
	require.NotNil(t, balance)
	assert.Equal(t, 0, balance.Integer)
	assert.Equal(t, 0, balance.Decimal)
}

var accountID = primitive.NewObjectID()
var expenseInput1 = &models.ExpenseInput{
	Title:    "title",
	Location: strPtr("location"),
	Entries: []*models.ExpenseEntryInput{
		{
			Title:      "food",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 5,
				Decimal: 50,
			},
		},
		{
			Title:      "sweets",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 12,
				Decimal: 80,
			},
		},
	},
	TotalBalance: &models.MoneyAmountInput{
		Integer: 6,
		Decimal: 40,
	},
	Date:      strPtr("12032019"),
	AccountID: &accountID,
}

var expenseInput2 = &models.ExpenseInput{
	Title:    "title",
	Location: strPtr("location"),
	Entries: []*models.ExpenseEntryInput{
		{
			Title:      "food",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 1,
				Decimal: 50,
			},
		},
		{
			Title:      "cat food",
			CategoryID: primitive.ObjectID{},
			Balance: &models.MoneyAmountInput{
				Integer: 2,
				Decimal: 30,
			},
		},
		{
			Title:      "medicine",
			CategoryID: primitive.NewObjectID(),
			Balance: &models.MoneyAmountInput{
				Integer: 3,
				Decimal: 60,
			},
		},
	},
	TotalBalance: &models.MoneyAmountInput{
		Integer: 10,
		Decimal: 70,
	},
	Date:      strPtr("12032019"),
	AccountID: &accountID,
}

func assertEvent(t *testing.T, ch <-chan *models.ExpenseEvent, eventType models.EventType, expense *models.Expense) {
	select {
	case event := <-ch:
		assert.Equal(t, &models.ExpenseEvent{Type: eventType, Expense: expense}, event)
	default:
		t.Error("no event received")
	}
}

func assertNoEvent(t *testing.T, ch <-chan *models.ExpenseEvent) {
	select {
	case event := <-ch:
		t.Errorf("unexpected event: %#v", event)
	default:
	}
}
