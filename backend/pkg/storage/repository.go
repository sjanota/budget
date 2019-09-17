package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	collection *mongo.Collection
}

type decodeFunc func(interface{}) error

func (r *repository) Find(ctx context.Context, filter doc, consumer func(decodeFunc) error) error {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer func() { _= cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		err := consumer(cursor.Decode)
		if err != nil {
			return err
		}
	}

	return nil
}