package storage

import (
	"context"
	"fmt"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage/collections"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpensesRepository struct {
	*repository
	watchers map[chan models.ExpenseEvent]struct{}
}

func newExpensesRepository(db *mongo.Database) *ExpensesRepository {
	return &ExpensesRepository{
		repository: &repository{
			collection: db.Collection(collections.EXPENSES),
		},
		watchers: make(map[chan models.ExpenseEvent]struct{}),
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

func (r *ExpensesRepository) InsertOne(ctx context.Context, input *models.ExpenseInput) (*models.Expense, error) {
	result, err := r.collection.InsertOne(ctx, input)
	if err != nil {
		return nil, err
	}

	entries := make([]*models.ExpenseEntry, len(input.Entries))
	for i, entry := range input.Entries {
		entries[i] = &models.ExpenseEntry{
			Title:      entry.Title,
			CategoryID: entry.CategoryID,
			Amount:     entry.Amount,
		}
	}
	expense := &models.Expense{
		ID:        result.InsertedID.(primitive.ObjectID),
		Title:     input.Title,
		Location:  input.Location,
		Entries:   entries,
		Total:     input.Total,
		Date:      input.Date,
		AccountID: input.AccountID,
	}
	r.notify(models.ExpenseAdded{
		ID:      expense.ID,
		Type:    models.EventTypeAdded,
		Expense: expense,
	})
	return expense, nil
}

func (r *ExpensesRepository) Watch(ctx context.Context) (<-chan models.ExpenseEvent, error) {
	events := make(chan models.ExpenseEvent)
	r.watchers[events] = struct{}{}
	go func() {
		defer close(events)
		defer func() {
			delete(r.watchers, events)
		}()
		<-ctx.Done()
		fmt.Println("<-ctx.Done()")
	}()
	return events, nil
}

func (r *ExpensesRepository) notify(event models.ExpenseEvent) {
	for watcher := range r.watchers {
		watcher <- event
	}
}