package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Storage struct {
	db *mongo.Database
	expenses *ExpensesRepository
	accounts *AccountsRepository
}

func (s *Storage) Expenses() *ExpensesRepository {
	return s.expenses
}

func (s *Storage) Accounts() *AccountsRepository {
	return s.accounts
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
		expenses: newExpensesRepository(database),
		accounts: newAccountsRepository(database),
	}, nil
}

func (s *Storage) Drop(ctx context.Context) error {
	return s.db.Drop(ctx)
}

func (s *Storage) Init(ctx context.Context) error {
	return nil
}
