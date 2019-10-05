package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const Budgets = "budgets"
const MonthlyReports = "monthly_reports"

type Storage struct {
	db             *mongo.Database
	budgets        *mongo.Collection
	monthlyReports *mongo.Collection
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
		db:             database,
		budgets:        database.Collection(Budgets),
		monthlyReports: database.Collection(MonthlyReports),
	}

	return storage, nil
}

func (s *Storage) Drop(ctx context.Context) error {
	return s.db.Drop(ctx)
}

func (s *Storage) Init(ctx context.Context) error {
	return s.createMonthlyReportIndexes(ctx)
}
