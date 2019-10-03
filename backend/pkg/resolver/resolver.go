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

func (r *Resolver) Category() schema.CategoryResolver {
	return &categoryResolver{r}
}

type categoryResolver struct {
	*Resolver
}

func (r *categoryResolver) Envelope(ctx context.Context, obj *models.Category) (*models.Envelope, error) {
	return r.Storage.GetEnvelope(ctx, obj.BudgetID, obj.EnvelopeID)
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
	return &mutationResolver{r}
}

type mutationResolver struct {
	*Resolver
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in map[string]interface{}) (*models.Account, error) {
	return r.Storage.UpdateAccount(ctx, budgetID, id, in)
}

func (r *mutationResolver) CreateCategory(ctx context.Context, budgetID primitive.ObjectID, in models.CategoryInput) (*models.Category, error) {
	return r.Storage.CreateCategory(ctx, budgetID, &in)
}

func (r *mutationResolver) CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, in models.EnvelopeInput) (*models.Envelope, error) {
	return r.Storage.CreateEnvelope(ctx, budgetID, &in)
}

func (r *mutationResolver) CreateAccount(ctx context.Context, budgetID primitive.ObjectID, in models.AccountInput) (*models.Account, error) {
	return r.Storage.CreateAccount(ctx, budgetID, &in)
}

func (r *mutationResolver) CreateBudget(ctx context.Context) (*models.Budget, error) {
	return r.Storage.CreateBudget(ctx)
}
