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
	BudgetID    primitive.ObjectID
}

type ExpenseEntry struct {
	Title      string      `json:"title"`
	Balance    MoneyAmount `json:"amount"`
	CategoryID primitive.ObjectID
}

type Account struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	BudgetID primitive.ObjectID
}

type Envelope struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	BudgetID primitive.ObjectID
}
