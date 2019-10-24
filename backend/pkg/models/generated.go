// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Problem interface {
	IsProblem()
}

type AccountInput struct {
	Name string `json:"name"`
}

type CategoryInput struct {
	Name       string             `json:"name"`
	EnvelopeID primitive.ObjectID `json:"envelopeID"`
}

type CategoryUpdate struct {
	Name       *string             `json:"name"`
	EnvelopeID *primitive.ObjectID `json:"envelopeID"`
}

type EnvelopeInput struct {
	Name  string  `json:"name"`
	Limit *Amount `json:"limit"`
}

type EnvelopeOverLimit struct {
	Severity Severity           `json:"severity"`
	ID       primitive.ObjectID `json:"id"`
}

func (EnvelopeOverLimit) IsProblem() {}

type ExpenseCategoryInput struct {
	CategoryID primitive.ObjectID `json:"categoryID"`
	Amount     Amount             `json:"amount"`
}

type ExpenseInput struct {
	Title      string                  `json:"title"`
	Categories []*ExpenseCategoryInput `json:"categories"`
	AccountID  primitive.ObjectID      `json:"accountID"`
	Date       Date                    `json:"date"`
}

type ExpenseUpdate struct {
	Title      *string                 `json:"title"`
	Categories []*ExpenseCategoryInput `json:"categories"`
	AccountID  *primitive.ObjectID     `json:"accountID"`
	Date       *Date                   `json:"date"`
}

type Misplanned struct {
	Severity    Severity `json:"severity"`
	Overplanned bool     `json:"overplanned"`
}

func (Misplanned) IsProblem() {}

type MonthStillInProgress struct {
	Severity Severity `json:"severity"`
}

func (MonthStillInProgress) IsProblem() {}

type NegativeBalanceOnAccount struct {
	Severity Severity           `json:"severity"`
	ID       primitive.ObjectID `json:"id"`
}

func (NegativeBalanceOnAccount) IsProblem() {}

type NegativeBalanceOnEnvelope struct {
	Severity Severity           `json:"severity"`
	ID       primitive.ObjectID `json:"id"`
}

func (NegativeBalanceOnEnvelope) IsProblem() {}

type PlanInput struct {
	Title           string              `json:"title"`
	FromEnvelopeID  *primitive.ObjectID `json:"fromEnvelopeID"`
	ToEnvelopeID    primitive.ObjectID  `json:"toEnvelopeID"`
	CurrentAmount   Amount              `json:"currentAmount"`
	RecurringAmount *Amount             `json:"recurringAmount"`
}

type TransferInput struct {
	Title         string              `json:"title"`
	FromAccountID *primitive.ObjectID `json:"fromAccountID"`
	ToAccountID   primitive.ObjectID  `json:"toAccountID"`
	Amount        Amount              `json:"amount"`
	Date          Date                `json:"date"`
}

type Severity string

const (
	SeverityError   Severity = "ERROR"
	SeverityWarning Severity = "WARNING"
	SeverityInfo    Severity = "INFO"
)

var AllSeverity = []Severity{
	SeverityError,
	SeverityWarning,
	SeverityInfo,
}

func (e Severity) IsValid() bool {
	switch e {
	case SeverityError, SeverityWarning, SeverityInfo:
		return true
	}
	return false
}

func (e Severity) String() string {
	return string(e)
}

func (e *Severity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Severity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Severity", str)
	}
	return nil
}

func (e Severity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
