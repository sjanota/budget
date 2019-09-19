package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Expense struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Location  *string            `json:"location"`
	Entries   []*ExpenseEntry    `json:"entries"`
	Total     MoneyAmount        `json:"total"`
	Date      *string            `json:"date"`
	AccountID *primitive.ObjectID
}

type Category struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	EnvelopeID  primitive.ObjectID
}

type ExpenseEntry struct {
	Title      string      `json:"title"`
	Amount     MoneyAmount `json:"amount"`
	CategoryID primitive.ObjectID
}

func (i MoneyAmountInput) ToMoneyAmount() *MoneyAmount {
	return &MoneyAmount{
		Integer: i.Integer,
		Decimal: i.Decimal,
	}
}

func (i ExpenseEntryInput) ToExpenseEntry() *ExpenseEntry {
	return &ExpenseEntry{
		Title:      i.Title,
		CategoryID: i.CategoryID,
		Amount:     *i.Amount.ToMoneyAmount(),
	}
}

func (i ExpenseInput) ToExpense(id primitive.ObjectID) *Expense {
	entries := make([]*ExpenseEntry, len(i.Entries))
	for i, entry := range i.Entries {
		entries[i] = entry.ToExpenseEntry()
	}

	return &Expense{
		ID:        id,
		Title:     i.Title,
		Location:  i.Location,
		Entries:   entries,
		Total:     *i.Total.ToMoneyAmount(),
		Date:      i.Date,
		AccountID: i.AccountID,
	}
}