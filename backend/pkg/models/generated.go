// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountInput struct {
	Name string `json:"name"`
}

type CategoriesFilter struct {
	EnvelopeID *primitive.ObjectID `json:"envelopeID"`
}

type EnvelopeInput struct {
	Name  string  `json:"name"`
	Limit *Amount `json:"limit"`
}
