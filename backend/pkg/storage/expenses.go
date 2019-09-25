package storage

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type expensesRepository struct {
	*repository
	storage  *Storage
	watchers map[chan *models.ExpenseEvent]struct{}
}

func newExpensesRepository(storage *Storage) *expensesRepository {
	return &expensesRepository{
		watchers: make(map[chan *models.ExpenseEvent]struct{}),
		repository: &repository{
			storage:    storage,
			collection: storage.db.Collection("expenses"),
		},
	}
}

type Expenses struct {
	*expensesRepository
	budgetID primitive.ObjectID
}

func (r *expensesRepository) session(budgetID primitive.ObjectID) *Expenses {
	return &Expenses{
		expensesRepository: r,
		budgetID:           budgetID,
	}
}

func (r *Expenses) TotalBalanceForAccount(ctx context.Context, accountID primitive.ObjectID) (*models.MoneyAmount, error) {
	cursor, err := r.collection.Aggregate(ctx, []doc{
		{"$match": doc{"budgetid": r.budgetID, "accountid": accountID}},
		{
			"$group": doc{
				"_id": nil,
				"integer": doc{
					"$sum": "$totalbalance.integer",
				},
				"decimal": doc{
					"$sum": "$totalbalance.decimal",
				},
			},
		},
		{
			"$project": doc{
				"_id": 0,
				"integer": doc{
					"$sum": list{
						"$integer",
						doc{
							"$floor": doc{
								"$divide": list{"$decimal", 100},
							},
						},
					},
				},
				"decimal": doc{
					"$mod": list{"$decimal", 100},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if !cursor.Next(ctx) {
		return &models.MoneyAmount{}, nil
	}

	result := &models.MoneyAmount{}
	err = cursor.Decode(result)
	return result, err
}

func (r *Expenses) FindAll(ctx context.Context) ([]*models.Expense, error) {
	result := make([]*models.Expense, 0)
	err := r.find(ctx, doc{budgetID: r.budgetID}, func(d decodeFunc) error {
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
	err := r.findOne(ctx, doc{"budgetid": r.budgetID, _id: id}, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return result, err
}

func (r *Expenses) DeleteByID(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	result := &models.Expense{}
	err := r.deleteOne(ctx, doc{"budgetid": r.budgetID, _id: id}, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeDeleted,
		Expense: result,
	})
	return result, err
}

func (r *Expenses) ReplaceByID(ctx context.Context, id primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	result := &models.Expense{}
	replacement := input.ToModel(r.budgetID)
	err := r.replaceOne(ctx, doc{"budgetid": r.budgetID, _id: id}, replacement, result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeUpdated,
		Expense: result,
	})
	return result, err
}

func (r *Expenses) Insert(ctx context.Context, input models.ExpenseInput) (*models.Expense, error) {
	if err := r.expectBudget(ctx, r.budgetID); err != nil {
		return nil, err
	}
	result := input.ToModel(r.budgetID)
	id, err := r.insertOne(ctx, result)
	if err != nil {
		return nil, err
	}

	r.notify(&models.ExpenseEvent{
		Type:    models.EventTypeCreated,
		Expense: result,
	})
	return result.WithID(id), nil
}

func (r *Expenses) Watch(ctx context.Context) (<-chan *models.ExpenseEvent, error) {
	if err := r.expectBudget(ctx, r.budgetID); err != nil {
		return nil, err
	}
	events := make(chan *models.ExpenseEvent, 1)
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
