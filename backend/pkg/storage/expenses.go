package storage

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
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