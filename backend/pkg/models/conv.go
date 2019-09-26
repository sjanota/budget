package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	var totalBalance MoneyAmount
	if i.TotalBalance != nil {
		totalBalance = *i.TotalBalance.ToModel()
	}

	return &Expense{
		Title:        i.Title,
		Location:     i.Location,
		Entries:      entries,
		TotalBalance: totalBalance,
		Date:         i.Date,
		AccountID:    i.AccountID,
		BudgetID:     budgetID,
	}
}

func (e *Expense) WithID(id primitive.ObjectID) *Expense {
	e.ID = id
	return e
}

func (a AccountInput) ToModel(budgetID primitive.ObjectID) *Account {
	return &Account{
		Name:     a.Name,
		BudgetID: budgetID,
	}
}

func (a *Account) WithID(id primitive.ObjectID) *Account {
	a.ID = id
	return a
}

func (a EnvelopeInput) ToModel(budgetID primitive.ObjectID) *Envelope {
	return &Envelope{
		Name:     a.Name,
		BudgetID: budgetID,
	}
}

func (a *Envelope) WithID(id primitive.ObjectID) *Envelope {
	a.ID = id
	return a
}

func (i CategoryInput) ToModel(budgetID primitive.ObjectID) *Category {
	return &Category{
		Name:     i.Name,
		BudgetID: budgetID,
	}
}

func (c *Category) WithID(id primitive.ObjectID) *Category {
	c.ID = id
	return c
}
