package resolver

import (
	"context"
	"github.com/sjanota/budget/pkg/models"
	"github.com/sjanota/budget/pkg/schema"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Category() schema.CategoryResolver {
	return &categoryResolver{r}
}

func (r *Resolver) Expense() schema.ExpenseResolver {
	return &expenseResolver{r}
}

func (r *Resolver) Query() schema.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Expenses(ctx context.Context, since *string, until *string) ([]*models.Expense, error) {
	account := "account"
	return []*models.Expense{
		{
			ID:        "bla",
			Title:     "bla",
			Location:  nil,
			Entries:   nil,
			Total:     nil,
			Date:      nil,
			AccountID: &account,
		},
	}, nil
}


