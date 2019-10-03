package models

import (
	"encoding/json"
	"io"
	"math"
	"strconv"

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
