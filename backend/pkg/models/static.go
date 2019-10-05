package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Budget struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Accounts       []*Account         `json:"accounts"`
	Envelopes      []*Envelope        `json:"envelopes"`
	Categories     []*Category        `json:"categories"`
	CurrentMonthID primitive.ObjectID
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

type MonthlyReport struct {
	ID        MonthlyReportID `bson:"_id,omitempty"`
	Expenses  []*Expense         `json:"expenses"`
	Transfers []*Transfer        `json:"transfers"`
	Plans     []*Plan            `json:"plans"`
}

func (r MonthlyReport) Month() time.Month {
	return r.ID.Month
}

func (r MonthlyReport) Year() int {
	return r.ID.Year
}

type MonthlyReportID struct {
	Month    time.Month
	Year     int
	BudgetID primitive.ObjectID
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

type ExpenseCategory struct {
	Amount     Amount `json:"balance"`
	CategoryID primitive.ObjectID
	BudgetID   primitive.ObjectID
}

type Account struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Balance  Amount
	BudgetID primitive.ObjectID
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

type CategoryInput struct {
	Name       string             `json:"name"`
	EnvelopeID primitive.ObjectID `json:"envelopeID"`
}
