package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, input *models.EnvelopeInput) (*models.Envelope, error) {
	if exists, err := s.budgetEntityExistsByName(ctx, budgetID, "envelopes", input.Name); err != nil {
		return nil, err
	} else if exists {
		return nil, ErrEnvelopeAlreadyExists
	}

	toInsert := &models.Envelope{Name: input.Name, Limit: input.Limit, ID: primitive.NewObjectID()}
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

func (s *Storage) GetEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Envelope, error) {
	budget, err := s.getBudgetByEntityID(ctx, budgetID, "envelopes", id)
	if err != nil {
		return nil, err
	}
	if len(budget.Envelopes) == 0 {
		return nil, nil
	}

	envelope := budget.Envelopes[0]
	envelope.BudgetID = budgetID
	return envelope, nil
}
