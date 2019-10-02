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

type expenseCategoryResolver struct {
	*Resolver
}

func (r *expenseCategoryResolver) Category(ctx context.Context, obj *models.ExpenseCategory) (*models.Category, error) {
	return r.Storage.Categories(obj.BudgetID).FindByID(ctx, obj.CategoryID)
}

func (r *Resolver) ExpenseCategory() schema.ExpenseCategoryResolver {
	return &expenseCategoryResolver{}
}

func (r *Resolver) Envelope() schema.EnvelopeResolver {
	return &envelopeResolver{r, r.Storage.Playground()}
}

func (r *Resolver) Account() schema.AccountResolver {
	return &accountResolver{r, r.Storage.Playground()}
}

func (r *Resolver) Budget() schema.BudgetResolver {
	return &budgetResolver{r}
}

func (r *Resolver) Subscription() schema.SubscriptionResolver {
	return &subscriptionResolver{r}
}

func (r *Resolver) Mutation() schema.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Category() schema.CategoryResolver {
	return &categoryResolver{r}
}

func (r *Resolver) Expense() schema.ExpenseResolver {
	return &expenseResolver{r}
}

func (r *Resolver) Query() schema.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Category(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Category, error) {
	return r.Storage.Categories(budgetID).FindByID(ctx, id)
}

func (r *queryResolver) Categories(ctx context.Context, budgetID primitive.ObjectID) ([]*models.Category, error) {
	return r.Storage.Categories(budgetID).FindAll(ctx)
}

func (r *queryResolver) Envelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Envelope, error) {
	return r.Storage.Envelopes(budgetID).FindByID(ctx, id)
}

func (r *queryResolver) Envelopes(ctx context.Context, budgetID primitive.ObjectID) ([]*models.Envelope, error) {
	return r.Storage.Envelopes(budgetID).FindAll(ctx)
}

func (r *queryResolver) Expense(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Expense, error) {
	return r.Storage.Expenses(budgetID).FindByID(ctx, id)
}

func (r *queryResolver) Account(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Account, error) {
	return r.Storage.Accounts(budgetID).FindByID(ctx, id)
}

func (r *queryResolver) Accounts(ctx context.Context, budgetID primitive.ObjectID) ([]*models.Account, error) {
	return r.Storage.Accounts(budgetID).FindAll(ctx)
}

func (r *queryResolver) Expenses(ctx context.Context, budgetID primitive.ObjectID) ([]*models.Expense, error) {
	return r.Storage.Expenses(budgetID).FindAll(ctx)
}

func (r *queryResolver) Budget(ctx context.Context, id primitive.ObjectID) (*models.Budget, error) {
	return r.Storage.Budgets().FindByID(ctx, id)
}

func (r *queryResolver) Budgets(ctx context.Context) ([]*models.Budget, error) {
	return r.Storage.Budgets().FindAll(ctx)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCategory(ctx context.Context, budgetID primitive.ObjectID, input models.CategoryInput) (*models.Category, error) {
	return r.Storage.Categories(budgetID).Insert(ctx, input)
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, input models.CategoryInput) (*models.Category, error) {
	return r.Storage.Categories(budgetID).ReplaceByID(ctx, id, input)
}

func (r *mutationResolver) CreateEnvelope(ctx context.Context, budgetID primitive.ObjectID, input models.EnvelopeInput) (*models.Envelope, error) {
	return r.Storage.Envelopes(budgetID).Insert(ctx, input)
}

func (r *mutationResolver) UpdateEnvelope(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, input models.EnvelopeInput) (*models.Envelope, error) {
	return r.Storage.Envelopes(budgetID).ReplaceByID(ctx, id, input)
}

func (r *mutationResolver) CreateAccount(ctx context.Context, budgetID primitive.ObjectID, input models.AccountInput) (*models.Account, error) {
	return r.Storage.Accounts(budgetID).Insert(ctx, input)
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, input models.AccountInput) (*models.Account, error) {
	return r.Storage.Accounts(budgetID).ReplaceByID(ctx, id, input)
}

func (r *mutationResolver) CreateExpense(ctx context.Context, budgetID primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses(budgetID).Insert(ctx, input)
}

func (r *mutationResolver) DeleteExpense(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID) (*models.Expense, error) {
	return r.Storage.Expenses(budgetID).DeleteByID(ctx, id)
}

func (r *mutationResolver) UpdateExpense(ctx context.Context, budgetID primitive.ObjectID, id primitive.ObjectID, input models.ExpenseInput) (*models.Expense, error) {
	return r.Storage.Expenses(budgetID).ReplaceByID(ctx, id, input)
}

func (r *mutationResolver) CreateBudget(ctx context.Context, name string) (*models.Budget, error) {
	return r.Storage.Budgets().Insert(ctx, name)
}
