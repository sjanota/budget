package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=mocks/generated_storage.go github.com/sjanota/budget/backend/pkg/resolver Storage
type Storage interface {
	CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, in *models.EnvelopeInput) (*models.Envelope, error)
	GetEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Envelope, error)
	UpdateEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Envelope, error)

	CreateMonthlyReport(ctx context.Context, budgetID primitive.ObjectID, month models.Month) (*models.MonthlyReport, error)
	GetMonthlyReport(ctx context.Context, id models.MonthlyReportID) (*models.MonthlyReport, error)

	CreateCategory(ctx context.Context, budgetID primitive.ObjectID, in *models.CategoryInput) (*models.Category, error)
	UpdateCategory(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Category, error)

	CreateAccount(ctx context.Context, budgetID primitive.ObjectID, in *models.AccountInput) (*models.Account, error)
	UpdateAccount(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Account, error)

	CreateBudget(ctx context.Context, currentMonth models.Month) (*models.Budget, error)
	GetBudget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error)
	ListBudgets(ctx context.Context) ([]*models.Budget, error)

	CreateExpense(ctx context.Context, reportID models.MonthlyReportID, in *models.ExpenseInput) (*models.Expense, error)
}

type Resolver struct {
	Storage Storage
}

var _ schema.ResolverRoot = &Resolver{}

func (r *Resolver) Budget() schema.BudgetResolver {
	return &budgetResolver{r}
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
	return &mutationResolver{Resolver: r}
}
