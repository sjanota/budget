package resolver

import (
	"context"
	"time"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/mutation_resolver_storage.go -package=mocks github.com/sjanota/budget/backend/pkg/resolver MutationResolverStorage
type MutationResolverStorage interface {
	UpdateCategory(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Category, error)
	UpdateEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Envelope, error)
	UpdateAccount(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Account, error)
	CreateCategory(ctx context.Context, budgetID primitive.ObjectID, in *models.CategoryInput) (*models.Category, error)
	CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, in *models.EnvelopeInput) (*models.Envelope, error)
	CreateAccount(ctx context.Context, budgetID primitive.ObjectID, in *models.AccountInput) (*models.Account, error)
	CreateMonthlyReport(ctx context.Context, budgetID primitive.ObjectID, in *models.MonthlyReportInput) (*models.MonthlyReport, error)
	CreateBudget(ctx context.Context, id primitive.ObjectID, currentMonth models.Month) (*models.Budget, error)
}

type mutationResolver struct {
	Storage MutationResolverStorage
	Now func() time.Time
	NewObjectID func() primitive.ObjectID
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
	budgetID := r.NewObjectID()
	now := r.Now()
	month := models.Month{
		Year:  now.Year(),
		Month: now.Month(),
	}
	input := &models.MonthlyReportInput{
		Month: month,
	}
	_, err := r.Storage.CreateMonthlyReport(ctx, budgetID, input)
	if err != nil {
		return nil, err
	}

	return r.Storage.CreateBudget(ctx, budgetID, month)
}
