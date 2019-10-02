package models

import (
	"math"
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

func (b Budget) Category(name string) *Category {
	for _, category := range b.Categories {
		if category.Name == name {
			return category
		}
	}
	return nil
}

func (b Budget) Account(name string) *Account {
	for _, account := range b.Accounts {
		if account.Name == name {
			return account
		}
	}
	return nil
}

func (b Budget) Envelope(name string) *Envelope {
	for _, envelope := range b.Envelopes {
		if envelope.Name == name {
			return envelope
		}
	}
	return nil
}

type MonthlyBudget struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Month     time.Month         `json:"month"`
	Year      uint               `json:"year"`
	Expenses  []Expense          `json:"expenses"`
	Transfers []Transfer         `json:"transfers"`
	Plans     []Plan             `json:"plans"`
	BudgetID  primitive.ObjectID
}

type Expense struct {
	Categories []*ExpenseCategory `json:"categories"`
	Date       *time.Time         `json:"date"`
	AccountID  *primitive.ObjectID
}

func (e Expense) Balance() Amount {
	var sum = Amount{0, 0}
	for _, c := range e.Categories {
		sum = sum.Add(c.Balance)
	}
	return sum
}

type ExpenseCategory struct {
	Balance    Amount `json:"balance"`
	CategoryID primitive.ObjectID
	BudgetID   primitive.ObjectID
}

type Account struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Balance  Amount
	BudgetID primitive.ObjectID
}

type AccountInput struct {
	Name string `json:"name"`
}

type Transfer struct {
	Balance       Amount    `json:"balance"`
	Date          time.Time `json:"date"`
	FromAccountID primitive.ObjectID
	ToAccountID   primitive.ObjectID
}

type Envelope struct {
	Name     string `json:"name"`
	Limit    Amount `json:"Limit"`
	Balance  Amount
	BudgetID primitive.ObjectID
}

type EnvelopeInput struct {
	Name  string `json:"name"`
	Limit Amount `json:"Limit"`
}

type Plan struct {
	Balance        Amount    `json:"balance"`
	Executed       Amount    `json:"executed"`
	Date           time.Time `json:"date"`
	FromEnvelopeID primitive.ObjectID
	ToEnvelopeID   primitive.ObjectID
}

type Category struct {
	Name         string `json:"name"`
	EnvelopeName string
	BudgetID     primitive.ObjectID
}

type CategoryInput struct {
	Name         string `json:"name"`
	EnvelopeName string `json:"envelopeName"`
}

func (a Amount) Add(other Amount) Amount {
	decimal := a.Decimal + other.Decimal
	return Amount{
		Integer: a.Integer + other.Integer + decimal/100,
		Decimal: decimal % 100,
	}
}

func (a Amount) Sub(other Amount) Amount {
	decimal := a.Decimal - other.Decimal
	timesOverflown := int(math.Floor(float64(a.Decimal) / float64(100)))
	return Amount{
		Integer: decimal + timesOverflown*100,
		Decimal: a.Integer - other.Integer - timesOverflown,
	}
}

func (a Amount) IsBiggerThan(other Amount) bool {
	return a.Integer >= other.Integer && a.Decimal > other.Decimal
}
