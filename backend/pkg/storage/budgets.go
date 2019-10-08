package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Storage) CreateBudget(ctx context.Context, name string, currentMonth models.Month) (*models.Budget, error) {
	budget := &models.Budget{
		Name:         name,
		Accounts:     []*models.Account{},
		Envelopes:    []*models.Envelope{},
		Categories:   []*models.Category{},
		CurrentMonth: currentMonth,
	}
	result, err := s.budgets.InsertOne(ctx, budget)
	if err != nil {
		return nil, err
	}

	budget.ID = result.InsertedID.(primitive.ObjectID)
	return budget, nil
}

func (s *Storage) ListBudgets(ctx context.Context) ([]*models.Budget, error) {
	cursor, err := s.budgets.Find(ctx, doc{})
	if err != nil {
		return nil, err
	}

	result := make([]*models.Budget, 0)
	err = cursor.All(ctx, &result)
	return result, err
}

func (s *Storage) GetBudget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	budget, err := s.budgets.FindOneByID(ctx, id)
	if err == ErrNoBudget {
		return nil, nil
	}
	return budget, err
}

func (s *Storage) budgetEntityExistsByName(ctx context.Context, budgetID primitive.ObjectID, arrayField, name string) (bool, error) {
	find := doc{
		"_id":                budgetID,
		arrayField + ".name": name,
	}
	res := s.budgets.FindOne(ctx, find)
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
	res := s.budgets.FindOne(ctx, find)
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) getBudgetByEntityID(ctx context.Context, budgetID primitive.ObjectID, arrayField string, id primitive.ObjectID) (*models.Budget, error) {
	project := doc{
		arrayField: doc{
			"$elemMatch": doc{
				"_id": id,
			},
		},
	}
	return s.budgets.FindOneByID(ctx, budgetID, options.FindOne().SetProjection(project))
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
	res, err := s.budgets.UpdateOne(ctx, find, update)
	if err != nil {
		return err
	} else if res.MatchedCount == 0 {
		return ErrNoBudget
	}
	return nil
}

func (s *Storage) updateAndVerifyEntityInBudget(ctx context.Context, budgetID, id primitive.ObjectID, arrayField string, changes models.Changes) (*models.Budget, error) {
	if exists, err := s.budgetEntityExistsByID(ctx, budgetID, arrayField, id); err != nil {
		return nil, err
	} else if !exists {
		return nil, ErrDoesNotExists
	}

	return s.updateEntityInBudget(ctx, budgetID, id, arrayField, changes)
}

func (s *Storage) updateEntityInBudget(ctx context.Context, budgetID, id primitive.ObjectID, arrayField string, changes models.Changes) (*models.Budget, error) {
	find := doc{
		"_id":               budgetID,
		arrayField + "._id": id,
	}
	project := doc{
		arrayField: doc{
			"$elemMatch": doc{
				"_id": id,
			},
		},
	}
	updateFields := doc{}
	for field, value := range changes {
		updateFields[arrayField+".$."+field] = value
	}
	update := doc{
		"$set": updateFields,
	}
	res := s.budgets.FindOneAndUpdate(ctx, find, update, options.FindOneAndUpdate().SetProjection(project).SetReturnDocument(options.After))
	if err := res.Err(); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	result := &models.Budget{}
	err := res.Decode(result)
	return result, err
}
