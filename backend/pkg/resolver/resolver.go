package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/storage"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=mocks/generated_storage.go -source=resolver.go

type Storage interface {
	CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, in *models.EnvelopeInput) (*models.Envelope, error)
	GetEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Envelope, error)
	UpdateEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Envelope, error)

	CreateMonthlyReport(ctx context.Context, budgetID primitive.ObjectID, month models.Month, plans []*models.PlanInput) (*models.MonthlyReport, error)
	GetMonthlyReport(ctx context.Context, id models.MonthlyReportID) (*models.MonthlyReport, error)

	CreateCategory(ctx context.Context, budgetID primitive.ObjectID, in *models.CategoryInput) (*models.Category, error)
	UpdateCategory(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.CategoryUpdate) (*models.Category, error)

	CreateAccount(ctx context.Context, budgetID primitive.ObjectID, in *models.AccountInput) (*models.Account, error)
	UpdateAccount(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, in models.Changes) (*models.Account, error)

	CreateBudget(ctx context.Context, name string, currentMonth models.Month) (*models.Budget, error)
	GetBudget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error)
	ListBudgets(ctx context.Context) ([]*models.Budget, error)

	CreateExpense(ctx context.Context, reportID models.MonthlyReportID, in *models.ExpenseInput) (*models.Expense, error)
	GetExpensesTotalForAccount(ctx context.Context, reportID models.MonthlyReportID, accountID primitive.ObjectID) (models.Amount, error)
	GetExpensesTotalForEnvelope(ctx context.Context, reportID models.MonthlyReportID, envelopeID primitive.ObjectID) (models.Amount, error)
	GetCategory(ctx context.Context, budgetID primitive.ObjectID, categoryID primitive.ObjectID) (*models.Category, error)
	GetAccount(ctx context.Context, budgetID primitive.ObjectID, accountID primitive.ObjectID) (*models.Account, error)
	UpdateExpense(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID, update storage.ChangeSet) (*models.Expense, error)
	CreateTransfer(ctx context.Context, reportID models.MonthlyReportID, in *models.TransferInput) (*models.Transfer, error)
	UpdateTransfer(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID, update models.TransferUpdate) (*models.Transfer, error)
	GetTransfersTotalForAccount(ctx context.Context, reportID models.MonthlyReportID, accountID primitive.ObjectID) (models.Amount, error)
	CreatePlan(ctx context.Context, reportID models.MonthlyReportID, in *models.PlanInput) (*models.Plan, error)
	UpdatePlan(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID, update models.PlanUpdate) (*models.Plan, error)
	GetPlansTotalForEnvelope(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID) (models.Amount, error)
	ReplaceBudget(ctx context.Context, budget *models.Budget) (*models.Budget, error)
	DeleteExpense(ctx context.Context, reportID models.MonthlyReportID, id primitive.ObjectID) (*models.Expense, error)
}

var _ schema.ResolverRoot = &Resolver{}

type Resolver struct {
	Storage Storage
}

func (r *Resolver) MonthlyReport() schema.MonthlyReportResolver {
	return &monthlyReportResolver{r}
}

func (r *Resolver) Account() schema.AccountResolver {
	return &accountResolver{r}
}

func (r *Resolver) Envelope() schema.EnvelopeResolver {
	return &envelopeResolver{r}
}

func (r *Resolver) Budget() schema.BudgetResolver {
	return &budgetResolver{r}
}

func (r *Resolver) Expense() schema.ExpenseResolver {
	return &expenseResolver{r}
}

func (r *Resolver) ExpenseCategory() schema.ExpenseCategoryResolver {
	return &expenseCategoryResolver{r}
}

func (r *Resolver) Plan() schema.PlanResolver {
	return &planResolver{r}
}

func (r *Resolver) Transfer() schema.TransferResolver {
	return &transferResolver{r}
}

func (r *Resolver) Category() schema.CategoryResolver {
	return &categoryResolver{r}
}

func (r *Resolver) Query() schema.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() schema.MutationResolver {
	return &mutationResolver{Resolver: r}
}
