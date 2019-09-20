package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/schema"
	"github.com/sjanota/budget/backend/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resolver struct {
	Storage *storage.Storage
}

func (r *Resolver) Subscription() schema.SubscriptionResolver {
	return &subscriptionResolver{r}
}

func (r *Resolver) ExpenseEntry() schema.ExpenseEntryResolver {
	return &expenseEntryResolver{r}
}

func (r *Resolver) Mutation() schema.MutationResolver {
	return &mutationResolver{r}
}

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
	return r.Storage.Expenses().FindAll(ctx)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAccount(ctx context.Context, input models.AccountInput) (*models.Account, error) {
	return r.Storage.Accounts().InsertOne(ctx, input)
}

func (r *mutationResolver) UpdateExpense(ctx context.Context, id primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses().ReplaceByID(ctx, id, input)
}

func (r *mutationResolver) DeleteExpense(ctx context.Context, id primitive.ObjectID) (*models.Expense, error) {
	return r.Storage.Expenses().DeleteByID(ctx, id)
}

func (r *mutationResolver) CreateExpense(ctx context.Context, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses().InsertOne(ctx, input)
}

type subscriptionResolver struct {
	*Resolver
}

func (r *subscriptionResolver) ExpenseEvents(ctx context.Context) (<-chan *models.ExpenseEvent, error) {
	return r.Storage.Expenses().Watch(ctx)
}


