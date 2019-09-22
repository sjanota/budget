package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type expensesRepository struct {
	*Storage
	watchers map[chan *models.ExpenseEvent]struct{}
}

func newExpensesRepository(storage *Storage) *expensesRepository {
	return &expensesRepository{
		Storage:  storage,
		watchers: make(map[chan *models.ExpenseEvent]struct{}),
	}
}

func (r *expensesRepository) ForBudget(budgetID primitive.ObjectID) *Expenses {
	return &Expenses{
		expensesRepository: r,
		budgetID:           budgetID,
	}
}

type Expenses struct {
	*expensesRepository
	budgetID primitive.ObjectID
}

func (r *Expenses) FindAll(ctx context.Context) ([]*models.Expense, error) {
	var result []*models.Expense
	err := r.find(ctx, Doc{}, func(d decodeFunc) error {
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

func (r *Expenses) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	result := &models.Expense{}
	err := r.findByID(ctx, id, result)
	return result, err
}

func (r *Expenses) DeleteByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	result := &models.Expense{}
	err := r.deleteByID(ctx, id, result)
	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeDeleted,
		Expense: result,
	})
	return result, err
}

func (r *Expenses) ReplaceByID(ctx context.Context, id primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	result := &models.Expense{}
	replacement := input.ToModel(id)
	err := r.replaceByID(ctx, id, replacement, result)
	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeUpdated,
		Expense: result,
	})
	return result, err
}

func (r *Expenses) Insert(ctx context.Context, input models.ExpenseInput) (*models.Expense, error) {
	result, err := r.collection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	expense := input.ToModel(result.InsertedID.(primitive.ObjectID))
	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeCreated,
		Expense: expense,
	})
	return expense, nil
}

func (r *Expenses) Watch(ctx context.Context) (<-chan *models.ExpenseEvent, error) {
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

func (r *Expenses) notify(event *models.ExpenseEvent) {
	for watcher := range r.watchers {
		watcher <- event
	}
}
