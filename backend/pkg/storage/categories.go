package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type categoriesRepository struct {
	*repository
	storage *Storage
}

func newCategoriesRepository(s *Storage) *categoriesRepository {
	return &categoriesRepository{
		repository: &repository{
			storage:    s,
			collection: s.db.Collection("categories"),
		},
	}
}

type Categories struct {
	*categoriesRepository
	budgetID primitive.ObjectID
}

func (r categoriesRepository) session(budgetID primitive.ObjectID) *Categories {
	return &Categories{
		categoriesRepository: &r,
		budgetID:            budgetID,
	}
}

func (r *Categories) FindAll(ctx context.Context) ([]*models.Category, error) {
	result := make([]*models.Category, 0)
	err := r.find(ctx, doc{budgetID: r.budgetID}, func(d decodeFunc) error {
		e := &models.Category{}
		err := d(e)
		if err != nil {
			return err
		}
		result = append(result, e)
		return nil
	})
	return result, err
}

func (r *Categories) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error) {
	result := &models.Category{}
	err := r.findByID(ctx, id, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Categories) ReplaceByID(ctx context.Context, id primitive.ObjectID, input models.CategoryInput) (*models.Category, error) {
	result := &models.Category{}
	replacement := input.ToModel(r.budgetID)
	err := r.replaceOne(ctx, doc{budgetID: r.budgetID}, replacement, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Categories) Insert(ctx context.Context, input models.CategoryInput) (*models.Category, error) {
	if err := r.expectBudget(ctx, r.budgetID); err != nil {
		return nil, err
	}
	category := input.ToModel(r.budgetID)
	result, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	return category.WithID(result.InsertedID.(primitive.ObjectID)), nil
}
