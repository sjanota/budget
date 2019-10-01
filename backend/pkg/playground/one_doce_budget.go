package playground

import (
	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Storage interface {
	GetMonthlyBudget(budgetID primitive.ObjectID, monthlyBudgetID primitive.ObjectID) (*models.MonthlyBudget, error)
	GetBudget(budgetID primitive.ObjectID) (*models.Budget, error)
	GetAccount(budgetID primitive.ObjectID, accountID primitive.ObjectID) (*models.Account, error)
	GetEnvelope(budgetID primitive.ObjectID, envelopeID primitive.ObjectID) (*models.Envelope, error)
	GetCurrentMonthlyBudget(budgetID primitive.ObjectID) (*models.MonthlyBudget, error)
	GetCurrentExpensesForAccount(budgetID primitive.ObjectID, monthlyBudgetID primitive.ObjectID, accountID primitive.ObjectID) ([]*models.Expense, error)
	EnsureMonthlyBudget(budgetID primitive.ObjectID, month time.Month, year uint) (*models.MonthlyBudget, error)
	UpdateMonthlyBudget(budgetID primitive.ObjectID, monthlyBudget *models.MonthlyBudget) (*models.MonthlyBudget, error)
	UpdateBudget(budget *models.Budget) (*models.Budget, error)
}

// Better to be done in memory
func CloseMonthlyBudget(budgetID primitive.ObjectID, monthlyBudgetID primitive.ObjectID, storage Storage) error {
	budget, err := storage.GetBudget(budgetID)
	if err != nil {
		return err
	}

	monthlyBudget, err := storage.GetMonthlyBudget(budgetID, monthlyBudgetID)
	if err != nil {
		return err
	}

	// processExpenses
	for _, expense := range monthlyBudget.Expenses {
		account := budget.Account(expense.AccountID)
		for _, expenseCategory := range expense.Categories {
			category := budget.Category(expenseCategory.CategoryID)
			envelope := budget.Envelope(category.EnvelopeID)

			account.Balance = account.Balance.Sub(expenseCategory.Balance)
			envelope.Balance = envelope.Balance.Sub(expenseCategory.Balance)
		}
	}

	// processTransfers
	for _, transfer := range monthlyBudget.Transfers {
		fromAccount := budget.Account(transfer.FromAccountID)
		fromAccount.Balance = fromAccount.Balance.Sub(transfer.Balance)

		toAccount := budget.Account(transfer.ToAccountID)
		toAccount.Balance = toAccount.Balance.Add(transfer.Balance)
	}

	// processPlans
	for _, plan := range monthlyBudget.Plans {
		fromEnvelope := budget.Envelope(plan.FromEnvelopeID)
		fromEnvelope.Balance = fromEnvelope.Balance.Sub(plan.Balance)

		toEnvelope := budget.Envelope(plan.ToEnvelopeID)
		toEnvelope.Balance = toEnvelope.Balance.Add(plan.Balance)
		plan.Executed = plan.Balance
		if toEnvelope.Balance.IsBiggerThan(toEnvelope.Limit) {
			plan.Executed = plan.Balance.Sub(toEnvelope.Balance).Add(toEnvelope.Limit)
			toEnvelope.Balance = toEnvelope.Limit
		}
	}

	// getNextMonth
	month := monthlyBudget.Month + 1
	year := monthlyBudget.Year
	if monthlyBudget.Month == time.December {
		month = time.January
		year += 1
	}

	// createNextMonthlyBudget
	nextMonthlyBudget, err := storage.EnsureMonthlyBudget(budget.ID, month, year)
	if err != nil {
		return err
	}
	budget.CurrentMonthID = nextMonthlyBudget.ID

	// commitCurrentMonthlyBudget
	_, err = storage.UpdateMonthlyBudget(budget.ID, monthlyBudget)
	if err != nil {
		return err
	}

	// commitBudget
	_, err = storage.UpdateBudget(budget)
	return err
}

// Should be possible to do as one DB operation
func GetAccount(budgetID primitive.ObjectID, accountID primitive.ObjectID, storage Storage) (*models.Account, error) {

}
