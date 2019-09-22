package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	collection *mongo.Collection
}

type decodeFunc func(interface{}) error

func (r *repository) find(ctx context.Context, filter Doc, consumer func(decodeFunc) error) error {
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

func (r *repository) findOne(ctx context.Context, filter Doc, v interface{}) error {
	result := r.collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (r *repository) FindOneRaw(ctx context.Context, filter Doc) (bson.Raw, error) {
	result := r.collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, err
	}
	return result.DecodeBytes()
}

func (r *repository) findByID(ctx context.Context, id primitive.ObjectID, v interface{}) error {
	return r.findOne(ctx, Doc{_id: id}, v)
}

func (r *repository) deleteOne(ctx context.Context, filter Doc, v interface{}) error {
	result := r.collection.FindOneAndDelete(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (r *repository) deleteByID(ctx context.Context, id primitive.ObjectID, v interface{}) error {
	return r.deleteOne(ctx, Doc{_id: id}, v)
}

func (r *repository) replaceOne(ctx context.Context, filter Doc, replacement interface{}, v interface{}) error {
	result := r.collection.FindOneAndReplace(ctx, filter, replacement, options.FindOneAndReplace().SetReturnDocument(options.After))
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (r *repository) replaceByID(ctx context.Context, id primitive.ObjectID, replacement interface{}, v interface{}) error {
	return r.replaceOne(ctx, Doc{_id: id}, replacement, v)
}
