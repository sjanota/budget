package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) Accounts(ctx context.Context, budgetID primitive.ObjectID) ([]*models.Account, error) {
	budget, err := r.Storage.GetBudget(ctx, budgetID)
	if err != nil {
		return nil, err
	}
	return budget.Accounts, nil
}

func (r *queryResolver) Envelopes(ctx context.Context, budgetID primitive.ObjectID) ([]*models.Envelope, error) {
	budget, err := r.Storage.GetBudget(ctx, budgetID)
	if err != nil {
		return nil, err
	}
	return budget.Envelopes, nil
}

func (r *queryResolver) Categories(ctx context.Context, budgetID primitive.ObjectID) ([]*models.Category, error) {
	budget, err := r.Storage.GetBudget(ctx, budgetID)
	if err != nil {
		return nil, err
	}
	return budget.Categories, nil
}

func (r *queryResolver) Budgets(ctx context.Context) ([]*models.Budget, error) {
	return r.Storage.ListBudgets(ctx)
}

func (r *queryResolver) Budget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	return r.Storage.GetBudget(ctx, id)
}

