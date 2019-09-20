// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Available *MoneyAmount       `json:"available"`
}

type AccountInput struct {
	Name string `json:"name"`
}

type BudgetPlan struct {
	ID     primitive.ObjectID `json:"id"`
	Date   *string            `json:"date"`
	From   *Envelope          `json:"from"`
	To     *Account           `json:"to"`
	Amount *MoneyAmount       `json:"amount"`
}

type Envelope struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Available   *MoneyAmount       `json:"available"`
	Expenses    []*Expense         `json:"expenses"`
	BudgetPlans []*BudgetPlan      `json:"budgetPlans"`
}

type ExpenseEntryInput struct {
	Title      string             `json:"title"`
	CategoryID primitive.ObjectID `json:"categoryID"`
	Amount     *MoneyAmountInput  `json:"amount"`
}

type ExpenseEvent struct {
	Type    EventType `json:"type"`
	Expense *Expense  `json:"expense"`
}

type ExpenseInput struct {
	Title     string               `json:"title"`
	Location  *string              `json:"location"`
	Entries   []*ExpenseEntryInput `json:"entries"`
	Total     *MoneyAmountInput    `json:"total"`
	Date      *string              `json:"date"`
	AccountID *primitive.ObjectID  `json:"accountID"`
}

type MoneyAmount struct {
	Integer int `json:"integer"`
	Decimal int `json:"decimal"`
}

type MoneyAmountInput struct {
	Integer int `json:"integer"`
	Decimal int `json:"decimal"`
}

type Transfer struct {
	ID     primitive.ObjectID `json:"id"`
	Date   *string            `json:"date"`
	From   *Account           `json:"from"`
	To     *Account           `json:"to"`
	Amount *MoneyAmount       `json:"amount"`
}

type Direction string

const (
	DirectionIn   Direction = "IN"
	DirectionOut  Direction = "OUT"
	DirectionBoth Direction = "BOTH"
)

var AllDirection = []Direction{
	DirectionIn,
	DirectionOut,
	DirectionBoth,
}

func (e Direction) IsValid() bool {
	switch e {
	case DirectionIn, DirectionOut, DirectionBoth:
		return true
	}
	return false
}

func (e Direction) String() string {
	return string(e)
}

func (e *Direction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Direction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Direction", str)
	}
	return nil
}

func (e Direction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventType string

const (
	EventTypeCreated EventType = "CREATED"
	EventTypeDeleted EventType = "DELETED"
	EventTypeUpdated EventType = "UPDATED"
)

var AllEventType = []EventType{
	EventTypeCreated,
	EventTypeDeleted,
	EventTypeUpdated,
}

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeCreated, EventTypeDeleted, EventTypeUpdated:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}

func (e *EventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventType", str)
	}
	return nil
}

func (e EventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
