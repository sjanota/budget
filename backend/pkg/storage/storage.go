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
	db       *mongo.Database
	expenses *expensesRepository
	budgets  *budgetsRepository
}

func (s *Storage) Expenses(budgetID primitive.ObjectID) *Expenses {
	return s.expenses.session(budgetID)
}

func (s *Storage) Budgets() *Budgets {
	return s.budgets.session()
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
		db: database,
	}

	return storage, nil
}

func (s *Storage) Drop(ctx context.Context) error {
	return s.db.Drop(ctx)
}

func (s *Storage) Init(ctx context.Context) error {
	var err error
	s.budgets, err = newBudgetsRepository(ctx, s)
	if err != nil {
		return err
	}
	s.expenses = newExpensesRepository(s)
	return nil
}

