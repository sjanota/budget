package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpensesRepository struct {
	*repository
}

func (r *ExpensesRepository) FindAll(ctx context.Context) ([]*models.Expense, error) {
	var result []*models.Expense
	err := r.Find(ctx, doc{}, func(d decodeFunc) error {
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

func (r *ExpensesRepository) InsertOne(ctx context.Context, expense *models.ExpenseInput) (*models.Expense, error) {
	result, err := r.collection.InsertOne(ctx, expense)
	if err != nil {
		return nil, err
	}

	entries := make([]*models.ExpenseEntry, len(expense.Entries))
	for i, entry := range expense.Entries {
		entries[i] = &models.ExpenseEntry{
			Title:      entry.Title,
			CategoryID: entry.CategoryID,
			Amount:     entry.Amount,
		}
	}
	return &models.Expense{
		ID:        result.InsertedID.(primitive.ObjectID),
		Title:     expense.Title,
		Location:  expense.Location,
		Entries:   entries,
		Total:     expense.Total,
		Date:      expense.Date,
		AccountID: expense.AccountID,
	}, nil

}