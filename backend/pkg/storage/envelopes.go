package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, input *models.EnvelopeInput) (*models.Envelope, error) {
	if exists, err := s.doesEnvelopeExist(ctx, budgetID, input.Name); err != nil {
		return nil, err
	} else if exists {
		return nil, ErrEnvelopeAlreadyExists
	}

	toInsert := &models.Envelope{Name: input.Name, Limit: input.Limit}
	find := doc{
		"_id": budgetID,
	}
	update := doc{
		"$push": doc{
			"envelopes": toInsert,
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

func (s *Storage) GetEnvelope(ctx context.Context, budgetID primitive.ObjectID, envelopeName string) (*models.Envelope, error) {
	find := doc{
		"_id":            budgetID,
		"envelopes.name": envelopeName,
	}
	project := doc{
		"envelopes.$": 1,
	}
	res := s.db.Collection(budgets).FindOne(ctx, find, options.FindOne().SetProjection(project))
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := &models.Budget{}
	err := res.Decode(result)
	if err != nil {
		return nil, err
	}
	account := result.Envelopes[0]
	account.BudgetID = budgetID
	return account, nil
}

func (s *Storage) doesEnvelopeExist(ctx context.Context, budgetID primitive.ObjectID, envelopeName string) (bool, error) {
	find := doc{
		"_id":            budgetID,
		"envelopes.name": envelopeName,
	}
	res := s.db.Collection(budgets).FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
