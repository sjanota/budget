package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

const Budgets = "budgets"
const MonthlyReports = "monthly_reports"

type Storage struct {
	db             *mongo.Database
	budgets        *budgetsCollection
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
		budgets:        &budgetsCollection{database.Collection(Budgets)},
		monthlyReports: database.Collection(MonthlyReports),
	}

	return storage, nil
}

func (s *Storage) Drop(ctx context.Context) error {
	return s.db.Drop(ctx)
}

func (s *Storage) Init(ctx context.Context) error {
	return nil
}

type budgetsCollection struct {
	*mongo.Collection
}

func (coll *budgetsCollection) FindOneByID(ctx context.Context, id primitive.ObjectID,	opts ...*options.FindOneOptions) (*models.Budget, error) {
	res := coll.FindOne(ctx, doc{"_id": id}, opts...)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, ErrNoBudget
	} else if err != nil {
		return nil, err
	}

	result := &models.Budget{}
	err := res.Decode(result)
	return result, err
}
