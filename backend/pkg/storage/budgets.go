package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) CreateBudget(ctx context.Context) (*models.Budget, error) {
	budget := &models.Budget{
		Accounts:   []*models.Account{},
		Envelopes:  []*models.Envelope{},
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

func (s *Storage) budgetEntityExistsByName(ctx context.Context, budgetID primitive.ObjectID, arrayField, name string) (bool, error) {
	find := doc{
		"_id":                budgetID,
		arrayField + ".name": name,
	}
	res := s.db.Collection(budgets).FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) budgetEntityExistsByID(ctx context.Context, budgetID primitive.ObjectID, arrayField string, id primitive.ObjectID) (bool, error) {
	find := doc{
		"_id":               budgetID,
		arrayField + "._id": id,
	}
	res := s.db.Collection(budgets).FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) getBudgetByEntityID(ctx context.Context, budgetID primitive.ObjectID, arrayField string, id primitive.ObjectID) (*models.Budget, error) {
	find := doc{
		"_id":          budgetID,
	}
	project := doc{
		arrayField: doc{
			"$elemMatch": doc{
				"_id": id,
			},
		},
	}
	res := s.db.Collection(budgets).FindOne(ctx, find, options.FindOne().SetProjection(project))
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, ErrNoBudget
	} else if err != nil {
		return nil, err
	}

	result := &models.Budget{}
	err := res.Decode(result)
	return result, err
}

func (s *Storage) pushEntityToBudgetByName(ctx context.Context, budgetID primitive.ObjectID, arrayField, name string, input interface{}) error {
	if exists, err := s.budgetEntityExistsByName(ctx, budgetID, arrayField, name); err != nil {
		return err
	} else if exists {
		return ErrAlreadyExists
	}

	return s.pushEntityToBudget(ctx, budgetID, arrayField, input)
}

func (s *Storage) pushEntityToBudget(ctx context.Context, budgetID primitive.ObjectID, arrayField string, input interface{}) error {
	find := doc{
		"_id": budgetID,
	}
	update := doc{
		"$push": doc{
			arrayField: input,
		},
	}
	res, err := s.db.Collection(budgets).UpdateOne(ctx, find, update)
	if err != nil {
		return err
	} else if res.MatchedCount == 0 {
		return ErrNoBudget
	}
	return nil
}