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

func (r MonthlyReport) Expense(id primitive.ObjectID) *Expense {
	for _, expense := range r.Expenses {
		if expense.ID == id {
			return expense
		}
	}
	return nil
}

func (r MonthlyReport) Transfer(id primitive.ObjectID) *Transfer {
	for _, transfer := range r.Transfers {
		if transfer.ID == id {
			return transfer
		}
	}
	return nil
}

func (r MonthlyReport) Plan(id primitive.ObjectID) *Plan {
	for _, plan := range r.Plans {
		if plan.ID == id {
			return plan
		}
	}
	return nil
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
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title      string             `json:"title"`
	Categories []*ExpenseCategory `json:"categories"`
	Date       Date               `json:"date"`
	AccountID  primitive.ObjectID
}

func (e Expense) TotalAmount() Amount {
	var sum Amount
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
}

func (c ExpenseCategory) WithAmount(amount Amount) *ExpenseCategory {
	c.Amount = amount
	return &c
}

type Account struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Balance Amount
}

type Transfer struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount        Amount             `json:"balance"`
	Title         string             `json:"title"`
	Date          Date               `json:"date"`
	FromAccountID *primitive.ObjectID
	ToAccountID   primitive.ObjectID
}

type Envelope struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Limit   *Amount            `json:"Limit"`
	Balance Amount
}

type Plan struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title          string             `json:"title"`
	Amount         Amount             `json:"balance"`
	FromEnvelopeID *primitive.ObjectID
	ToEnvelopeID   primitive.ObjectID
}

type Category struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name"`
	EnvelopeID primitive.ObjectID
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

func (u CategoryUpdate) Changes() Changes {
	result := make(map[string]interface{})
	if u.Name != nil {
		result["name"] = *u.Name
	}
	if u.EnvelopeID != nil {
		result["envelopeid"] = *u.EnvelopeID
	}
	return result
}

func (u CategoryUpdate) WithEnvelope(envelopeID primitive.ObjectID) *CategoryUpdate {
	u.EnvelopeID = &envelopeID
	return &u
}

func (u ExpenseUpdate) Changes() Changes {
	result := make(map[string]interface{})
	if u.Title != nil {
		result["title"] = *u.Title
	}
	if u.AccountID != nil {
		result["accountid"] = *u.AccountID
	}
	if u.Date != nil {
		result["date"] = *u.Date
	}
	if u.Categories != nil {
		result["categories"] = u.Categories
	}
	return result
}

func (u TransferUpdate) Changes() Changes {
	result := make(map[string]interface{})
	if u.Title != nil {
		result["title"] = *u.Title
	}
	if u.FromAccountID != nil {
		result["fromaccountid"] = *u.FromAccountID
	}
	if u.Date != nil {
		result["date"] = *u.Date
	}
	if u.ToAccountID != nil {
		result["toaccountid"] = *u.ToAccountID
	}
	if u.Amount != nil {
		result["amount"] = *u.Amount
	}
	return result
}

func (u PlanUpdate) Changes() Changes {
	result := make(map[string]interface{})
	if u.Title != nil {
		result["title"] = *u.Title
	}
	if u.FromEnvelopeID != nil {
		result["fromenvelopeid"] = *u.FromEnvelopeID
	}
	if u.ToEnvelopeID != nil {
		result["toenvelopeid"] = *u.ToEnvelopeID
	}
	if u.Amount != nil {
		result["amount"] = *u.Amount
	}
	return result
}