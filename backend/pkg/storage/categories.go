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

	toInsert := &models.Category{Name: input.Name, EnvelopeID: input.EnvelopeID, ID: primitive.NewObjectID()}
	if err := s.pushEntityToBudget(ctx, budgetID, "categories", toInsert); err != nil {
		return nil, err
	}
	toInsert.BudgetID = budgetID
	return toInsert, nil
}

func (s *Storage) GetCategory(ctx context.Context, budgetID, id primitive.ObjectID) (*models.Category, error) {
	budget, err := s.getBudgetByEntityID(ctx, budgetID, "categories", id)
	if err != nil {
		return nil, err
	}
	if len(budget.Categories) == 0 {
		return nil, nil
	}

	category := budget.Categories[0]
	category.BudgetID = budgetID
	return category, nil
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
				"_id": input.EnvelopeID,
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
		return ErrInvalidReference
	}
	if len(result.Categories) == 1 {
		return ErrAlreadyExists
	}
	return nil
}
