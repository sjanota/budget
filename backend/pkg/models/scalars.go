package models

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UnmarshalID(v interface{}) (primitive.ObjectID, error) {
	s, ok := v.(string)
	if !ok {
		return primitive.ObjectID{}, errors.New("ID must be string")
	}

	oid, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return primitive.ObjectID{}, errors.Wrap(err, "cannot unmarshal id")
	}

	return oid, nil
}

func MarshalID(id primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(id.Hex()))
	})
}

func UnmarshalMonth(v interface{}) (time.Month, error) {
	month, ok := v.(int)
	if !ok || month < 1 || month > 12 {
		return 0, errors.New("ID must be an intager in a range <1;12>")
	}

	return time.Month(month), nil
}

func MarshalMonth(month time.Month) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Itoa(int(month)))
	})
}

type Amount struct {
	Integer int `json:"integer"`
	Decimal int `json:"decimal"`
}

func (a *Amount) UnmarshalGQL(v interface{}) error {
	return mapstructure.Decode(v, a)
}

func (a Amount) MarshalGQL(w io.Writer) {
	_ = json.NewEncoder(w).Encode(a)
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

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

var errInvalidDate = errors.New("Date must be ISO 8601 date string")

func (a *Date) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return errInvalidDate
	}
	parsed, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return errInvalidDate
	}
	a.Year, a.Month, a.Day = parsed.Date()
	return nil
}

func (a Date) MarshalGQL(w io.Writer) {
	s := fmt.Sprintf("%v-%v-%v", a.Year, int(a.Month), a.Day)
	_ = json.NewEncoder(w).Encode(s)
}
