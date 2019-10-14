package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, input *models.EnvelopeInput) (*models.Envelope, error) {
	toInsert := &models.Envelope{Name: input.Name, Limit: input.Limit, ID: primitive.NewObjectID()}
	if err := s.pushEntityToBudgetByName(ctx, budgetID, "envelopes", input.Name, toInsert); err != nil {
		return nil, err
	}
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
	return envelope, nil
}

func (s *Storage) UpdateEnvelope(ctx context.Context, budgetID, id primitive.ObjectID, changes models.Changes) (*models.Envelope, error) {
	budget, err := s.updateAndVerifyEntityInBudget(ctx, budgetID, id, "envelopes", changes)
	if err != nil {
		return nil, err
	}
	if len(budget.Envelopes) == 0 {
		return nil, nil
	}

	envelope := budget.Envelopes[0]
	return envelope, nil
}
