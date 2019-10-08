package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Budget struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	Accounts     []*Account         `json:"accounts"`
	Envelopes    []*Envelope        `json:"envelopes"`
	Categories   []*Category        `json:"categories"`
	CurrentMonth Month
}

type Changes map[string]interface{}

func (b Budget) Category(id primitive.ObjectID) *Category {
	for _, category := range b.Categories {
		if category.ID == id {
			return category
		}
	}
	return nil
}

func (b Budget) Account(id primitive.ObjectID) *Account {
	for _, account := range b.Accounts {
		if account.ID == id {
			return account
		}
	}
	return nil
}

func (b Budget) Envelope(id primitive.ObjectID) *Envelope {
	for _, envelope := range b.Envelopes {
		if envelope.ID == id {
			return envelope
		}
	}
	return nil
}

func (b Budget) CurrentMonthID() MonthlyReportID {
	return MonthlyReportID{
		Month:    b.CurrentMonth,
		BudgetID: b.ID,
	}
}

type MonthlyReport struct {
	ID        MonthlyReportID `bson:"_id,omitempty"`
	Expenses  []*Expense      `json:"expenses"`
	Transfers []*Transfer     `json:"transfers"`
	Plans     []*Plan         `json:"plans"`
}

func (r MonthlyReport) Month() Month {
	return r.ID.Month
}

func (r MonthlyReport) WithBudget(budget Budget) MonthlyReport {
	r.ID = *r.ID.WithBudget(budget).WithMonth(budget.CurrentMonth)
	return r
}

type MonthlyReportID struct {
	Month    Month
	BudgetID primitive.ObjectID
}

func (id MonthlyReportID) WithBudget(budget Budget) *MonthlyReportID {
	id.BudgetID = budget.ID
	return &id
}

func (id MonthlyReportID) WithMonth(month Month) *MonthlyReportID {
	id.Month = month
	return &id
}

type Expense struct {
	Title      *string            `json:"title"`
	Categories []*ExpenseCategory `json:"categories"`
	Date       Date               `json:"date"`
	AccountID  primitive.ObjectID
	BudgetID   primitive.ObjectID
}

func (e Expense) TotalAmount() Amount {
	var sum = Amount{0, 0}
	for _, c := range e.Categories {
		sum = sum.Add(c.Amount)
	}
	return sum
}

func (e Expense) WithCategories(categories ...*ExpenseCategory) *Expense {
	e.Categories = categories
	return &e
}

type ExpenseCategory struct {
	Amount     Amount `json:"balance"`
	CategoryID primitive.ObjectID
	BudgetID   primitive.ObjectID
}

func (c ExpenseCategory) WithAmount(amount Amount) *ExpenseCategory {
	c.Amount = amount
	return &c
}

type Account struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Balance  Amount
	BudgetID primitive.ObjectID
}

func (a Account) WithBudget(budgetID primitive.ObjectID) *Account {
	a.BudgetID = budgetID
	return &a
}

type Transfer struct {
	Amount        Amount `json:"balance"`
	Title         string `json:"title"`
	FromAccountID primitive.ObjectID
	ToAccountID   primitive.ObjectID
}

type Envelope struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Limit    *Amount            `json:"Limit"`
	Balance  Amount
	BudgetID primitive.ObjectID
}

func (e Envelope) WithBudget(budgetID primitive.ObjectID) *Envelope {
	e.BudgetID = budgetID
	return &e
}

type Plan struct {
	Title          string `json:"title"`
	Amount         Amount `json:"balance"`
	Executed       Amount `json:"executed"`
	FromEnvelopeID primitive.ObjectID
	ToEnvelopeID   primitive.ObjectID
}

type Category struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name"`
	EnvelopeID primitive.ObjectID
	BudgetID   primitive.ObjectID
}

func (c Category) WithBudget(budgetID primitive.ObjectID) Category {
	c.BudgetID = budgetID
	return c
}

func (c Category) WithEnvelope(envelopeID primitive.ObjectID) Category {
	c.EnvelopeID = envelopeID
	return c
}

type CategoryInput struct {
	Name       string             `json:"name"`
	EnvelopeID primitive.ObjectID `json:"envelopeID"`
}

func (i ExpenseInput) WithDate(date Date) *ExpenseInput {
	i.Date = date
	return &i
}

func (i ExpenseInput) WithAccount(accountID primitive.ObjectID) *ExpenseInput {
	i.AccountID = accountID
	return &i
}

func (i ExpenseInput) WithCategories(categories ...*ExpenseCategoryInput) *ExpenseInput {
	i.Categories = categories
	return &i
}

func (i ExpenseCategoryInput) WithCategory(categoryID primitive.ObjectID) *ExpenseCategoryInput {
	i.CategoryID = categoryID
	return &i
}
