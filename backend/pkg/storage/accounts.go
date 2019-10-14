package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storage) CreateAccount(ctx context.Context, budgetID primitive.ObjectID, input *models.AccountInput) (*models.Account, error) {
	toInsert := &models.Account{Name: input.Name, ID: primitive.NewObjectID()}
	if err := s.pushEntityToBudgetByName(ctx, budgetID, "accounts", input.Name, toInsert); err != nil {
		return nil, err
	}
	return toInsert, nil
}

func (s *Storage) GetAccount(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Account, error) {
	budget, err := s.getBudgetByEntityID(ctx, budgetID, "accounts", id)
	if err != nil {
		return nil, err
	}
	if len(budget.Accounts) == 0 {
		return nil, nil
	}

	account := budget.Accounts[0]
	return account, nil
}

func (s *Storage) UpdateAccount(ctx context.Context, budgetID, id primitive.ObjectID, changes models.Changes) (*models.Account, error) {
	budget, err := s.updateAndVerifyEntityInBudget(ctx, budgetID, id, "accounts", changes)
	if err != nil {
		return nil, err
	}

	account := budget.Accounts[0]
	return account, nil
}
