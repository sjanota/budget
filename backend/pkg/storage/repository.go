package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type decodeFunc func(interface{}) error

type repository struct {
	storage    *Storage
	collection *mongo.Collection
}

func (r *repository) find(ctx context.Context, filter doc, consumer func(decodeFunc) error, findOptions ...*options.FindOptions) error {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer func() { _ = cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		err := consumer(cursor.Decode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) findOne(ctx context.Context, filter doc, v interface{}, opts ...*options.FindOneOptions) error {
	result := r.collection.FindOne(ctx, filter, opts...)
	if err := result.Err(); err == mongo.ErrNoDocuments {
		return err
	}
	return result.Decode(v)
}

func (r *repository) findByID(ctx context.Context, id primitive.ObjectID, v interface{}, opts ...*options.FindOneOptions) error {
	return r.findOne(ctx, doc{_id: id}, v, opts...)
}

func (r *repository) deleteOne(ctx context.Context, filter doc, v interface{}) error {
	result := r.collection.FindOneAndDelete(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (r *repository) deleteByID(ctx context.Context, id primitive.ObjectID, v interface{}) error {
	return r.deleteOne(ctx, doc{_id: id}, v)
}

func (r *repository) replaceOne(ctx context.Context, filter doc, replacement interface{}, v interface{}) error {
	result := r.collection.FindOneAndReplace(ctx, filter, replacement, options.FindOneAndReplace().SetReturnDocument(options.After))
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (r *repository) replaceByID(ctx context.Context, id primitive.ObjectID, replacement interface{}, v interface{}) error {
	return r.replaceOne(ctx, doc{_id: id}, replacement, v)
}

func (r *repository) insertOne(ctx context.Context, v interface{}) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, v)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

