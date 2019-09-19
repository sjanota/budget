package models

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"strconv"
)

type MoneyAmount float64

func (ma *MoneyAmount) UnmarshalGQL(v interface{}) error {
	f, err := graphql.UnmarshalFloat(v)
	if err != nil {
		return err
	}
	*ma = MoneyAmount(f)
	return nil
}

func (ma MoneyAmount) MarshalGQL(w io.Writer) {
	graphql.MarshalFloat(float64(ma)).MarshalGQL(w)
}

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
