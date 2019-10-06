package resolver

import (
	"context"
	"time"

	"github.com/sjanota/budget/backend/pkg/schema"
	"github.com/sjanota/budget/backend/pkg/storage"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ schema.MutationResolver = &mutationResolver{}

type mutationResolver struct {
	*Resolver
	Now         func() time.Time
	NewObjectID func() primitive.ObjectID
}

func (r *mutationResolver) CreateExpense(ctx context.Context, budgetID primitive.ObjectID, in models.ExpenseInput) (*models.Expense, error) {
	budget, err := r.Query().Budget(ctx, budgetID)
	if err != nil {
		return nil, err
	}

	reportID := models.MonthlyReportID{
		Month:    budget.CurrentMonth,
		BudgetID: budgetID,
	}
	expense, err := r.Storage.CreateExpense(ctx, reportID, &in)
	if err == storage.ErrNoReport {
		_, err = r.Storage.CreateMonthlyReport(ctx, budgetID, budget.CurrentMonth)
		if err == storage.ErrAlreadyExists || err == nil {
			expense, err = r.Storage.CreateExpense(ctx, reportID, &in)
		}
	}

	return expense, err
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in map[string]interface{}) (*models.Category, error) {
	return r.Storage.UpdateCategory(ctx, budgetID, id, in)
}

func (r *mutationResolver) UpdateEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in map[string]interface{}) (*models.Envelope, error) {
	return r.Storage.UpdateEnvelope(ctx, budgetID, id, in)
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
	now := r.Now()
	month := models.Month{
		Year:  now.Year(),
		Month: now.Month(),
	}
	return r.Storage.CreateBudget(ctx, month)
}