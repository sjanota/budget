package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	collection *mongo.Collection
}

type decodeFunc func(interface{}) error

func (r *repository) find(ctx context.Context, filter doc, consumer func(decodeFunc) error) error {
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

func (r *repository) findOne(ctx context.Context, filter doc, v interface{}) error {
	result := r.collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (r *repository) findByID(ctx context.Context, id primitive.ObjectID, v interface{}) error {
	return r.findOne(ctx, doc{_id: id}, v)
}