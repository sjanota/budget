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

func (r *Resolver) Budget() schema.BudgetResolver {
	return &BudgetResolver{Resolver: r}
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

func (r *queryResolver) Budget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	return r.Storage.Budgets().FindByID(ctx, id)
}

func (r *queryResolver) Budgets(ctx context.Context) ([]*models.Budget, error) {
	return r.Storage.Budgets().FindAll(ctx)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateBudget(ctx context.Context, name string) (*models.Budget, error) {
	return r.Storage.Budgets().Insert(ctx, name)
}

func (r *mutationResolver) Budget(ctx context.Context, id primitive.ObjectID) (models.BudgetMutation, error) {
	return &BudgetResolver{
		Resolver: r.Resolver,
	}, nil
}
