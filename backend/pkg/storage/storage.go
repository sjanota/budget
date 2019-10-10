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
	budgets        *budgets
	monthlyReports *monthlyReports
}

type ChangeSet interface {
	Changes() models.Changes
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
		budgets:        &budgets{&collectionExtension{database.Collection(Budgets)}},
		monthlyReports: &monthlyReports{&collectionExtension{database.Collection(MonthlyReports)}},
	}

	return storage, nil
}

func (s *Storage) Drop(ctx context.Context) error {
	return s.db.Drop(ctx)
}

func (s *Storage) Init(ctx context.Context) error {
	return nil
}

type collectionExtension struct {
	*mongo.Collection
}

func (coll *collectionExtension) FindOneByID(ctx context.Context, id interface{}, into interface{}, opts ...*options.FindOneOptions) error {
	res := coll.FindOne(ctx, doc{"_id": id}, opts...)
	if err := res.Err(); err != err {
		return err
	}
	return res.Decode(into)
}

type budgets struct {
	*collectionExtension
}

func (coll *budgets) FindOneByID(ctx context.Context, id primitive.ObjectID, opts ...*options.FindOneOptions) (*models.Budget, error) {
	result := &models.Budget{}
	err := coll.collectionExtension.FindOneByID(ctx, id, result, opts...)
	if err == mongo.ErrNoDocuments {
		return nil, ErrNoBudget
	}
	return result, err
}

type monthlyReports struct {
	*collectionExtension
}

func (coll *monthlyReports) FindOneByID(ctx context.Context, id models.MonthlyReportID, opts ...*options.FindOneOptions) (*models.MonthlyReport, error) {
	result := &models.MonthlyReport{}
	err := coll.collectionExtension.FindOneByID(ctx, id, result, opts...)
	if err == mongo.ErrNoDocuments {
		return nil, ErrNoReport
	}
	return result, err
}
