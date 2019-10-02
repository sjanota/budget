package models

import "go.mongodb.org/mongo-driver/bson/primitive"

func (i AmountInput) ToModel() *Amount {
	return &Amount{
		Integer: i.Integer,
		Decimal: i.Decimal,
	}
}

func (i ExpenseCategoryInput) ToModel() *ExpenseCategory {
	return &ExpenseCategory{
		CategoryID: i.CategoryID,
		Balance:    *i.Balance.ToModel(),
	}
}

func (i ExpenseInput) ToModel(budgetID primitive.ObjectID) *Expense {
	categories := make([]*ExpenseCategory, len(i.Entries))
	for i, entry := range i.Entries {
		categories[i] = entry.ToModel()
	}

	return &Expense{
		Categories: categories,
		Date:       i.Date,
		AccountID:  i.AccountID,
	}
}

func (a AccountInput) ToModel(budgetID primitive.ObjectID) *Account {
	return &Account{
		Name:     a.Name,
		BudgetID: budgetID,
	}
}

func (a EnvelopeInput) ToModel(budgetID primitive.ObjectID) *Envelope {
	return &Envelope{
		Name:     a.Name,
		BudgetID: budgetID,
	}
}
