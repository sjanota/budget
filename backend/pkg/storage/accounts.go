package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage/collections"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountsRepository struct {
	*repository
	//watchers map[chan *models.ExpenseEvent]struct{}
}

func newAccountsRepository(db *mongo.Database) *AccountsRepository {
	return &AccountsRepository{
		repository: &repository{
			collection: db.Collection(collections.ACCOUNTS),
		},
		//watchers: make(map[chan *models.ExpenseEvent]struct{}),
	}
}

func (r *AccountsRepository) FindAll(ctx context.Context) ([]*models.Account, error) {
	var result []*models.Account
	err := r.find(ctx, Doc{}, func(d decodeFunc) error {
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

func (r *AccountsRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Account, error) {
	result := &models.Account{}
	err := r.findByID(ctx, id, result)
	return result, err
}

func (r *AccountsRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) (*models.Account, error) {
	result := &models.Account{}
	err := r.deleteByID(ctx, id, result)
	//r.notify(&models.ExpenseEvent{
	//	Type:    models.EventTypeDeleted,
	//	Expense: result,
	//})
	return result, err
}

func (r *AccountsRepository) InsertOne(ctx context.Context, input models.AccountInput) (*models.Account, error) {
	account := &models.Account{
		Name: input.Name,
		Available: &models.MoneyAmount{
			Integer: 0,
			Decimal: 0,
		},
	}
	result, err := r.collection.InsertOne(ctx, account)
	if err != nil {
		return nil, err
	}

	account.ID = result.InsertedID.(primitive.ObjectID)

	//r.notify(&models.ExpenseEvent{
	//	Type:    models.EventTypeCreated,
	//	Expense: expense,
	//})
	return account, nil
}

//func (r *AccountsRepository) Watch(ctx context.Context) (<-chan *models.ExpenseEvent, error) {
//	events := make(chan *models.ExpenseEvent)
//	r.watchers[events] = struct{}{}
//	go func() {
//		defer close(events)
//		defer func() {
//			delete(r.watchers, events)
//		}()
//		<-ctx.Done()
//	}()
//	return events, nil
//}
//
//func (r *AccountsRepository) notify(event *models.ExpenseEvent) {
//	for watcher := range r.watchers {
//		watcher <- event
//	}
//}
