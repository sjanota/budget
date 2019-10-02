package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) CreateAccount(ctx context.Context, budgetID primitive.ObjectID, input *models.AccountInput) (*models.Account, error) {
	if exists, err := s.doesAccountWithTheSameNameExist(ctx, budgetID, input.Name); err != nil {
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

func (s *Storage) GetAccount(ctx context.Context, budgetID primitive.ObjectID, accountID primitive.ObjectID) (*models.Account, error) {
	find := doc{
		"_id":          budgetID,
		"accounts._id": accountID,
	}
	project := doc{
		"accounts.$": 1,
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
	account := result.Accounts[0]
	account.BudgetID = budgetID
	return account, nil
}

func (s *Storage) UpdateAccount(ctx context.Context, budgetID primitive.ObjectID, accountID primitive.ObjectID, changes models.Changes) (*models.Account, error) {
	if exists, err := s.doesAccountExist(ctx, budgetID, accountID); err != nil {
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

func (s *Storage) doesAccountWithTheSameNameExist(ctx context.Context, budgetID primitive.ObjectID, accountName string) (bool, error) {
	find := doc{
		"_id":           budgetID,
		"accounts.name": accountName,
	}
	res := s.db.Collection(budgets).FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) doesAccountExist(ctx context.Context, budgetID, accountID primitive.ObjectID) (bool, error) {
	find := doc{
		"_id":          budgetID,
		"accounts._id": accountID,
	}
	res := s.db.Collection(budgets).FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
