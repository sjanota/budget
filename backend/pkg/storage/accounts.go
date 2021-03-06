package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type accountsRepository struct {
	*repository
	storage *Storage
}

func newAccountsRepository(s *Storage) *accountsRepository {
	return &accountsRepository{
		repository: &repository{
			storage:    s,
			collection: s.db.Collection("accounts"),
		},
	}
}

type Accounts struct {
	*accountsRepository
	budgetID primitive.ObjectID
}

func (r accountsRepository) session(budgetID primitive.ObjectID) *Accounts {
	return &Accounts{
		accountsRepository: &r,
		budgetID:           budgetID,
	}
}

func (r *Accounts) FindAll(ctx context.Context) ([]*models.Account, error) {
	result := make([]*models.Account, 0)
	err := r.find(ctx, doc{budgetID: r.budgetID}, func(d decodeFunc) error {
		e := &models.Account{}
		err := d(e)
		if err != nil {
			return err
		}
		result = append(result, e)
		return nil
	})
	return result, err
}

func (r *Accounts) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Account, error) {
	result := &models.Account{}
	err := r.findByID(ctx, id, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Accounts) ReplaceByID(ctx context.Context, id primitive.ObjectID, input models.AccountInput) (*models.Account, error) {
	result := &models.Account{}
	replacement := input.ToModel(r.budgetID)
	err := r.replaceOne(ctx, doc{budgetID: r.budgetID}, replacement, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Accounts) Insert(ctx context.Context, input models.AccountInput) (*models.Account, error) {
	if err := r.expectBudget(ctx, r.budgetID); err != nil {
		return nil, err
	}
	account := &models.Account{
		Name:     input.Name,
		BudgetID: r.budgetID,
	}
	result, err := r.collection.InsertOne(ctx, account)
	if err != nil {
		return nil, err
	}

	account.ID = result.InsertedID.(primitive.ObjectID)
	return account, nil
}
