package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	time "time"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/schema"
	"github.com/sjanota/budget/backend/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resolver struct {
	Storage *storage.Storage
}

var _ schema.ResolverRoot = &Resolver{}

func (r *Resolver) Budget() schema.BudgetResolver {
	return &budgetResolver{r.Storage}
}

func (r *Resolver) Expense() schema.ExpenseResolver {
	panic("implement me")
}

func (r *Resolver) ExpenseCategory() schema.ExpenseCategoryResolver {
	panic("implement me")
}

func (r *Resolver) Plan() schema.PlanResolver {
	panic("implement me")
}

func (r *Resolver) Transfer() schema.TransferResolver {
	panic("implement me")
}

func (r *Resolver) Category() schema.CategoryResolver {
	return &categoryResolver{r.Storage}
}

func (r *Resolver) Query() schema.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) Budgets(ctx context.Context) ([]*models.Budget, error) {
	return r.Storage.ListBudgets(ctx)
}

func (r *queryResolver) Budget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	return r.Storage.GetBudget(ctx, id)
}

func (r *Resolver) Mutation() schema.MutationResolver {
	return &mutationResolver{Storage: r.Storage, Now: time.Now, NewObjectID: primitive.NewObjectID}
}
