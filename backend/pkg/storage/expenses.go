package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage/collections"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpensesRepository struct {
	*repository
	watchers map[chan *models.ExpenseEvent]struct{}
}

func newExpensesRepository(db *mongo.Database) *ExpensesRepository {
	return &ExpensesRepository{
		repository: &repository{
			collection: db.Collection(collections.EXPENSES),
		},
		watchers: make(map[chan *models.ExpenseEvent]struct{}),
	}
}

func (r *ExpensesRepository) FindAll(ctx context.Context) ([]*models.Expense, error) {
	var result []*models.Expense
	err := r.find(ctx, doc{}, func(d decodeFunc) error {
		e := &models.Expense{}
		err := d(e)
		if err != nil {
			return err
		}
		result = append(result, e)
		return nil
	})
	return result, err
}

func (r *ExpensesRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	result := &models.Expense{}
	err := r.findByID(ctx, id, result)
	return result, err
}

func (r *ExpensesRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	result := &models.Expense{}
	err := r.deleteByID(ctx, id, result)
	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeDeleted,
		Expense: result,
	})
	return result, err
}

func (r *ExpensesRepository) InsertOne(ctx context.Context, input *models.ExpenseInput) (*models.Expense, error) {
	result, err := r.collection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}


	expense := input.ToExpense(result.InsertedID.(primitive.ObjectID))
	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeCreated,
		Expense: expense,
	})
	return expense, nil
}

func (r *ExpensesRepository) Watch(ctx context.Context) (<-chan *models.ExpenseEvent, error) {
	events := make(chan *models.ExpenseEvent)
	r.watchers[events] = struct{}{}
	go func() {
		defer close(events)
		defer func() {
			delete(r.watchers, events)
		}()
		<-ctx.Done()
	}()
	return events, nil
}

func (r *ExpensesRepository) notify(event *models.ExpenseEvent) {
	for watcher := range r.watchers {
		watcher <- event
	}
}