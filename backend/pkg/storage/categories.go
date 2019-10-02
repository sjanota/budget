package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreateCategory(ctx context.Context, budgetID primitive.ObjectID, input *models.CategoryInput) (*models.Category, error) {
	if err := s.verifyCategoryInput(ctx, budgetID, input); err != nil {
		return nil, err
	}

	toInsert := &models.Category{Name: input.Name, EnvelopeName: input.EnvelopeName}
	find := doc{
		"_id": budgetID,
	}
	update := doc{
		"$push": doc{
			"categories": toInsert,
		},
	}
	res, err := s.db.Collection(budgets).UpdateOne(ctx, find, update)
	if err != nil {
		return nil, err
	} else if res.MatchedCount == 0 {
		return nil, ErrNoBudget
	}
	toInsert.BudgetID = budgetID
	return toInsert, nil
}

func (s *Storage) GetCategory(ctx context.Context, budgetID primitive.ObjectID, input models.CategoryInput) (*models.Category, error) {
	return &models.Category{}, nil
}

func (s *Storage) verifyCategoryInput(ctx context.Context, budgetID primitive.ObjectID, input *models.CategoryInput) error {
	find := doc{
		"_id": budgetID,
	}
	project := doc{
		"categories": doc{
			"$elemMatch": doc{
				"name": input.Name,
			},
		},
		"envelopes": doc{
			"$elemMatch": doc{
				"name": input.EnvelopeName,
			},
		},
	}
	res := s.db.Collection(budgets).FindOne(ctx, find, options.FindOne().SetProjection(project))
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return ErrNoBudget
	} else if err != nil {
		return err
	}

	result := &models.Budget{}
	err := res.Decode(result)
	if err != nil {
		return err
	}

	if len(result.Envelopes) == 0 {
		return ErrEnvelopeDoesNotExists
	}
	if len(result.Categories) == 1 {
		return ErrCategoryAlreadyExists
	}
	return nil
}
