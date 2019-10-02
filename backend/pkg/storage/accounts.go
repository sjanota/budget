package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreateAccount(ctx context.Context, budgetID primitive.ObjectID, input *models.AccountInput) (*models.Account, error) {
	toInsert := &models.Account{Name: input.Name, ID: primitive.NewObjectID()}
	if err := s.pushEntityToBudgetByName(ctx, budgetID, "accounts", input.Name, toInsert); err != nil {
		return nil, err
	}
	toInsert.BudgetID = budgetID
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
	account.BudgetID = budgetID
	return account, nil
}

func (s *Storage) UpdateAccount(ctx context.Context, budgetID primitive.ObjectID, accountID primitive.ObjectID, changes models.Changes) (*models.Account, error) {
	if exists, err := s.budgetEntityExistsByID(ctx, budgetID, "accounts", accountID); err != nil {
		return nil, err
	} else if !exists {
		return nil, ErrDoesNotExists
	}

	find := doc{
		"_id":          budgetID,
		"accounts._id": accountID,
	}
	project := doc{
		"accounts": doc{
			"$elemMatch": doc{
				"_id": accountID,
			},
		},
	}
	updateFields := doc{}
	for field, value := range changes {
		updateFields["accounts.$."+field] = value
	}
	update := doc{
		"$set": updateFields,
	}
	res := s.db.Collection(budgets).FindOneAndUpdate(ctx, find, update, options.FindOneAndUpdate().SetProjection(project).SetReturnDocument(options.After))
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
	account := result.Accounts[0]
	account.BudgetID = budgetID
	return account, nil
}
