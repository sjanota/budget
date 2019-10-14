package resolver

import (
	"context"
	"github.com/sjanota/budget/backend/pkg/models"
)

type transferResolver struct {
	*Resolver
}

func (r *transferResolver) FromAccount(ctx context.Context, obj *models.Transfer) (*models.Account, error) {
	return r.Storage.GetAccount(ctx, budgetFromContext(ctx), obj.FromAccountID)
}

func (r *transferResolver) ToAccount(ctx context.Context, obj *models.Transfer) (*models.Account, error) {
	return r.Storage.GetAccount(ctx, budgetFromContext(ctx), obj.ToAccountID)
}
