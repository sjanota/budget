package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"time"

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
	return &budgetResolver{r}
}

type budgetResolver struct {
	*Resolver
}

func (r *budgetResolver) CurrentMonth(ctx context.Context, obj *models.Budget) (*models.MonthlyReport, error) {
	return r.Storage.GetMonthlyReport(ctx, models.MonthlyReportID{BudgetID: obj.ID, Month: obj.CurrentMonth})
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
	budgetID := primitive.NewObjectID()
	now := time.Now()
	input := &models.MonthlyReportInput{
		Month: models.Month{
			Year:  now.Year(),
			Month: now.Month(),
		},
	}
	monthlyReport, err := r.Storage.CreateMonthlyReport(ctx, budgetID, input)
	if err != nil {
		return nil, err
	}

	return r.Storage.CreateBudget(ctx, budgetID, monthlyReport.ID.Month)
}
