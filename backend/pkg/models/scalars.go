package models

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
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

type Amount int64

func (a *Amount) UnmarshalGQL(v interface{}) error {
	s, ok := v.(json.Number)
	if !ok {
		return errors.New("Amount must be a number")
	}

	i, err := s.Int64()
	if err != nil {
		return err
	}
	*a = Amount(i)
	return nil
}

func (a Amount) MarshalGQL(w io.Writer) {
	_ = json.NewEncoder(w).Encode(a)
}

func (a Amount) Add(other Amount) Amount {
	return a + other
}

func (a Amount) Sub(other Amount) Amount {
	return a - other
}

func (a Amount) IsBiggerThan(other Amount) bool {
	return a > other
}

func (a Amount) IsNegative() bool {
	return a < 0
}

func NewAmount() Amount {
	return 0
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) ToMonth() Month {
	return Month{
		Year:  d.Year,
		Month: d.Month,
	}
}

func (d *Date) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return ErrMalformedDate
	}
	parsed, err := time.Parse("2006-01-02", s)
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

func (m Month) Previous() Month {
	if m.Month == time.January {
		return Month{
			Year:  m.Year - 1,
			Month: time.December,
		}
	}
	return Month{
		Year:  m.Year,
		Month: m.Month - 1,
	}
}
