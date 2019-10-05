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

func (d *Date) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return ErrMalformedDate
	}
	parsed, err := time.Parse("\"2006-01-02\"", s)
	if err != nil {
		return ErrMalformedDate
	}
	d.Year, d.Month, d.Day = parsed.Date()
	return nil
}

func (d Date) MarshalGQL(w io.Writer) {
	s := fmt.Sprintf("\"%04d-%02d-%02d\"", d.Year, int(d.Month), d.Day)
	_, _ = w.Write([]byte(s))
}

type Month struct {
	Year  int
	Month time.Month
}

func (m *Month) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return ErrMalformedMonth
	}
	parsed, err := time.Parse("\"2006-01\"", s)
	if err != nil {
		return ErrMalformedMonth
	}
	m.Year, m.Month, _ = parsed.Date()
	return nil
}

func (m Month) MarshalGQL(w io.Writer) {
	s := fmt.Sprintf("\"%04d-%02d\"", m.Year, int(m.Month))
	_, _ = w.Write([]byte(s))
}

func (m Month) Contains(d Date) bool {
	return m.Month == d.Month && m.Year == d.Year
}

func (m Month) Next() Month {
	if m.Month == time.December {
		return Month{
			Year:  m.Year + 1,
			Month: time.January,
		}
	}
	return Month{
		Year:  m.Year,
		Month: m.Month + 1,
	}
}
