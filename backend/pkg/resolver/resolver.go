package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	account := primitive.NewObjectID()
	return []*models.Expense{
		{
			ID:        primitive.NewObjectID(),
			Title:     "bla",
			Location:  nil,
			Entries:   nil,
			Total:     nil,
			Date:      nil,
			AccountID: &account,
		},
	}, nil
}


