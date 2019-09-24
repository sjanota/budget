package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type budgetsRepository struct {
	storage *Storage
	*repository
}

func newBudgetsRepository(ctx context.Context, storage *Storage) (*budgetsRepository, error) {
	collection := storage.db.Collection("budgets")
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    doc{"name": 1},
		Options: options.Index().SetUnique(true),
	})

	return &budgetsRepository{
		storage: storage,
		repository: &repository{
			collection: collection,
		},
	}, err
}

func (r *budgetsRepository) session() *Budgets {
	return &Budgets{r}
}

type Budgets struct {
	*budgetsRepository
}

func (r *Budgets) Insert(ctx context.Context, name string) (budget *models.Budget, err error) {
	budget = &models.Budget{
		Name:     name,
		Expenses: make([]*models.Expense, 0),
	}
	budget.ID, err = r.insertOne(ctx, budget)
	return
}

func (r *Budgets) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	result := &models.Budget{}
	err := r.findByID(ctx, id, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Budgets) FindAll(ctx context.Context) ([]*models.Budget, error) {
	result := make([]*models.Budget, 0)
	err := r.find(ctx, nil, func(d decodeFunc) error {
		entry := &models.Budget{}
		err := d(entry)
		if err != nil {
			return err
		}
		result = append(result, entry)
		return nil
	})
	return result, err
}

func (r *Budgets) DeleteByID(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	result := &models.Budget{}
	err := r.deleteByID(ctx, id, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}
