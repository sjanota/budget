package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type subscriptionResolver struct {
	*Resolver
}

func (r *subscriptionResolver) ExpenseEvent(ctx context.Context, budgetID primitive.ObjectID) (<-chan *models.ExpenseEvent, error) {
	return r.Storage.Budget(budgetID).Expenses().Watch(ctx)
}
