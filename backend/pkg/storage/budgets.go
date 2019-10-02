package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) CreateBudget(ctx context.Context) (*models.Budget, error) {
	budget := &models.Budget{
		Accounts: []*models.Account{},
		Envelopes: []*models.Envelope{},
		Categories: []*models.Category{},
	}
	result, err := s.db.Collection(budgets).InsertOne(ctx, budget)
	if err != nil {
		return nil, err
	}

	budget.ID = result.InsertedID.(primitive.ObjectID)
	return budget, nil
}

func (s *Storage) GetBudget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	find := doc{
		"_id": id,
	}

	res := s.db.Collection(budgets).FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := &models.Budget{}
	err := res.Decode(result)
	return result, err
}
