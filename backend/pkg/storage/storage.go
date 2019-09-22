package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const collectionName = "budgets"

type Storage struct {
	db         *mongo.Database
	collection *mongo.Collection
	expenses   *expensesRepository
}

func (s *Storage) Expenses(budgetID primitive.ObjectID) *Expenses {
	return s.expenses.ForBudget(budgetID)
}

func New(uri string) (*Storage, error) {
	opts := options.Client().ApplyURI(uri).SetRetryWrites(false)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	cs, err := connstring.Parse(uri)
	if err != nil {
		return nil, err
	}

	database := client.Database(cs.Database)
	storage := &Storage{
		db:         database,
		collection: database.Collection(collectionName),
	}
	storage.expenses = newExpensesRepository(storage)
	return storage, nil
}

func (s *Storage) Drop(ctx context.Context) error {
	return s.db.Drop(ctx)
}

func (s *Storage) Init(ctx context.Context) error {
	return nil
}

type decodeFunc func(interface{}) error

func (s *Storage) find(ctx context.Context, filter Doc, consumer func(decodeFunc) error) error {
	cursor, err := s.collection.Find(ctx, filter)
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

func (s *Storage) findOne(ctx context.Context, filter Doc, v interface{}) error {
	result := s.collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (s *Storage) findByID(ctx context.Context, id primitive.ObjectID, v interface{}) error {
	return s.findOne(ctx, Doc{_id: id}, v)
}

func (s *Storage) deleteOne(ctx context.Context, filter Doc, v interface{}) error {
	result := s.collection.FindOneAndDelete(ctx, filter)
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (s *Storage) deleteByID(ctx context.Context, id primitive.ObjectID, v interface{}) error {
	return s.deleteOne(ctx, Doc{_id: id}, v)
}

func (s *Storage) replaceOne(ctx context.Context, filter Doc, replacement interface{}, v interface{}) error {
	result := s.collection.FindOneAndReplace(ctx, filter, replacement, options.FindOneAndReplace().SetReturnDocument(options.After))
	if err := result.Err(); err != nil {
		return err
	}
	return result.Decode(v)
}

func (s *Storage) replaceByID(ctx context.Context, id primitive.ObjectID, replacement interface{}, v interface{}) error {
	return s.replaceOne(ctx, Doc{_id: id}, replacement, v)
}
