package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type expenseCategoryResolver struct {
	*Resolver
}

func (r *expenseCategoryResolver) Category(ctx context.Context, obj *models.ExpenseCategory) (*models.Category, error) {
	return r.Storage.GetCategory(ctx, budgetFromContext(ctx), obj.CategoryID)
}
