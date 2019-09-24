package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *expensesRepository) session(budgetID primitive.ObjectID) *Expenses {
	return &Expenses{
		expensesRepository: r,
		budgetID:           budgetID,
	}
}

type Expenses struct {
	*expensesRepository
	budgetID primitive.ObjectID
}

func (r *Expenses) TotalBalanceForAccount(ctx context.Context, accountID primitive.ObjectID) (bson.Raw, error) {
	cursor, err := r.collection.Aggregate(ctx, []doc{
		{opMatch: doc{_id: r.budgetID}},
		{
			"$project": doc{
				"expenses": doc{
					"$filter": doc {
						"input": "$expenses",
						"as": "expense",
						"cond": doc{
							"$eq": []interface{}{"$$expense.accountid", accountID},
						},
					},
				},
			},
		},
		{
			"$project": doc{
				"totalInt": doc{
					"$sum": "$expenses.totalbalance.integer",
				},
				"totalDec": doc{
					"$sum": "$expenses.totalbalance.decimal",
				},
				_id: 0,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if !cursor.Next(ctx) {
		return nil, nil
	}
	return cursor.Current, nil
}

func (r *Expenses) FindAll(ctx context.Context) ([]*models.Expense, error) {
	result := &models.Budget{}
	opts := options.FindOne().SetProjection(doc{expenses: 1})
	err := r.findByID(ctx, r.budgetID, result, opts)
	return result.Expenses, err
}

func (r *Expenses) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	result := &models.Expense{}
	cursor, err := r.collection.Aggregate(ctx, []doc{
		{opMatch: doc{_id: r.budgetID, "expenses.id": id}},
		{opUnwind: "$expenses"},
		{opReplaceRoot: doc{"newRoot": "$expenses"}},
	})
	if err != nil {
		return nil, err
	}

	if !cursor.Next(ctx) {
		return nil, nil
	}

	err = cursor.Decode(result)
	return result, err
}

func (r *Expenses) Delete(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	budget := &models.Budget{}
	opts := options.FindOneAndUpdate().SetProjection(
		doc{expenses: doc{opElemMatch: doc{"id": id}}},
	)
	result := r.collection.FindOneAndUpdate(ctx,
		doc{_id: r.budgetID},
		doc{opPull: doc{expenses: doc{"id": id}}},
		opts,
	)
	err := result.Decode(budget)
	if err != nil {
		return nil, err
	}

	//r.notify(&models.ExpenseEvent{
	//	Type:    models.EventTypeDeleted,
	//	Expense: result,
	//})
	return budget.Expenses[0], err
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
	result := input.ToModel()
	_, err := r.collection.UpdateOne(ctx,
		doc{_id: r.budgetID},
		doc{opPush: doc{expenses: result}},
	)
	if err != nil {
		return nil, err
	}

	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeCreated,
		Expense: result,
	})
	return result, nil
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
