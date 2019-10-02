package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreateAccount(ctx context.Context, budgetID primitive.ObjectID, input *models.AccountInput) (*models.Account, error) {
	if exists, err := s.budgetEntityExistsByName(ctx, budgetID, "accounts", input.Name); err != nil {
		return nil, err
	} else if exists {
		return nil, ErrAccountAlreadyExists
	}

	toInsert := &models.Account{Name: input.Name, ID: primitive.NewObjectID()}
	find := doc{
		"_id": budgetID,
	}
	update := doc{
		"$push": doc{
			"accounts": toInsert,
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
		return nil, ErrAccountDoesNotExists
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
