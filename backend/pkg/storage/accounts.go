package storage

import (
	"context"

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
			collection: s.db.Collection("accounts"),
		},
	}
}

func (r *accountsRepository) session(budgetID primitive.ObjectID) *Accounts {
	return &Accounts{
		accountsRepository: r,
		budgetID:           budgetID,
	}
}

type Accounts struct {
	*accountsRepository
	budgetID primitive.ObjectID
}

func (r *Accounts) FindAll(ctx context.Context) ([]*models.Account, error) {
	var result []*models.Account
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
	return result, err
}

func (r *Accounts) DeleteByID(ctx context.Context, id primitive.ObjectID) (*models.Account, error) {
	result := &models.Account{}
	err := r.deleteByID(ctx, id, result)
	return result, err
}

func (r *Accounts) InsertOne(ctx context.Context, input models.AccountInput) (*models.Account, error) {
	account := &models.Account{
		Name: input.Name,
		Balance: &models.MoneyAmount{
			Integer: 0,
			Decimal: 0,
		},
	}
	result, err := r.collection.InsertOne(ctx, account)
	if err != nil {
		return nil, err
	}

	account.ID = result.InsertedID.(primitive.ObjectID)
	return account, nil
}
