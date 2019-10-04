// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountInput struct {
	Name string `json:"name"`
}

type EnvelopeInput struct {
	Name  string  `json:"name"`
	Limit *Amount `json:"limit"`
}

type ExpenseCategoryInput struct {
	CategoryID primitive.ObjectID `json:"categoryID"`
	Amount     Amount             `json:"amount"`
}

type ExpenseInput struct {
	Title       *string                 `json:"title"`
	Categories  []*ExpenseCategoryInput `json:"categories"`
	AccountID   primitive.ObjectID      `json:"accountID"`
	TotalAmount Amount                  `json:"totalAmount"`
}

type PlanInput struct {
	Title          *string            `json:"title"`
	FromEnvelopeID primitive.ObjectID `json:"fromEnvelopeID"`
	ToEnvelopeID   primitive.ObjectID `json:"toEnvelopeID"`
	Amount         Amount             `json:"amount"`
}

type TransferInput struct {
	Title         *string            `json:"title"`
	FromAccountID primitive.ObjectID `json:"fromAccountID"`
	ToAccountID   primitive.ObjectID `json:"toAccountID"`
	Amount        Amount             `json:"amount"`
}
