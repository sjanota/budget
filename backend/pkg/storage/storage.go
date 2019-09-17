package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/storage/collections"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Storage struct {
	db *mongo.Database
	expenses *ExpensesRepository
}

func (s *Storage) Expenses() *ExpensesRepository {
	return s.expenses
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
	return &Storage{
		db: database,
		expenses: &ExpensesRepository{
			repository: &repository{
				collection: database.Collection(collections.EXPENSES),
			},
		},
	}, nil
}


