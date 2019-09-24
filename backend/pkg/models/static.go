package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Budget struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Expenses []*Expense
}

type Expense struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string             `json:"title"`
	Location     *string            `json:"location"`
	Entries      []*ExpenseEntry    `json:"entries"`
	TotalBalance MoneyAmount        `json:"total"`
	Date         *string            `json:"date"`
	AccountID    *primitive.ObjectID
	BudgetID     primitive.ObjectID
}

type Category struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	EnvelopeID  primitive.ObjectID
}

type ExpenseEntry struct {
	Title      string      `json:"title"`
	Balance    MoneyAmount `json:"amount"`
	CategoryID primitive.ObjectID
}

func (i MoneyAmountInput) ToModel() *MoneyAmount {
	return &MoneyAmount{
		Integer: i.Integer,
		Decimal: i.Decimal,
	}
}

func (i ExpenseEntryInput) ToModel() *ExpenseEntry {
	return &ExpenseEntry{
		Title:      i.Title,
		CategoryID: i.CategoryID,
		Balance:    *i.Balance.ToModel(),
	}
}

func (i ExpenseInput) ToModel(budgetID primitive.ObjectID) *Expense {
	entries := make([]*ExpenseEntry, len(i.Entries))
	for i, entry := range i.Entries {
		entries[i] = entry.ToModel()
	}

	return &Expense{
		Title:        i.Title,
		Location:     i.Location,
		Entries:      entries,
		TotalBalance: *i.TotalBalance.ToModel(),
		Date:         i.Date,
		AccountID:    i.AccountID,
		BudgetID:     budgetID,
	}
}

func (e *Expense) WithID(id primitive.ObjectID) *Expense {
	e.ID = id
	return e
}
