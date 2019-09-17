package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Expense struct {
	ID        primitive.ObjectID `json:"id"`
	Title     string             `json:"title"`
	Location  *string            `json:"location"`
	Entries   []*ExpenseEntry    `json:"entries"`
	Total     *MoneyAmount       `json:"total"`
	Date      *string            `json:"date"`
	AccountID *primitive.ObjectID
}

type Category struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	EnvelopeID  primitive.ObjectID
}
