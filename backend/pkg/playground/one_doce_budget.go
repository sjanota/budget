package playground

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ID primitive.ObjectID
type Date time.Time
type Amount float64

func (a Amount) Add(other Amount) Amount {
	return a + other
}

func (a Amount) Sub(other Amount) Amount {
	return a - other
}

func (a Amount) IsBiggerThan(other Amount) bool {
	return a > other
}

type Budget struct {
	ID             ID
	Accounts       []*Account
	Envelopes      []*Envelope
	Categories     []*Category
	CurrentMonthID ID
}

func (b Budget) Category(id ID) *Category {
	for _, category := range b.Categories {
		if category.ID == id {
			return category
		}
	}
	return nil
}

func (b Budget) Account(id ID) *Account {
	for _, account := range b.Accounts {
		if account.ID == id {
			return account
		}
	}
	return nil
}

func (b Budget) Envelope(id ID) *Envelope {
	for _, envelope := range b.Envelopes {
		if envelope.ID == id {
			return envelope
		}
	}
	return nil
}

type MonthlyBudget struct {
	BudgetID  ID
	ID        ID
	Month     time.Month
	Year      uint
	Expenses  []Expense
	Transfers []Transfer
	Plans     []Plan
}

type Expense struct {
	ID         ID
	AccountID  ID
	Categories []ExpenseCategory
	Date       Date
}

func (e Expense) Balance() Amount {
	var sum Amount = 0
	for _, c := range e.Categories {
		sum = sum.Add(c.Balance)
	}
	return sum
}

type ExpenseCategory struct {
	ID         ID
	CategoryID ID
	Balance    Amount
}

type Account struct {
	ID      ID
	Name    string
	Balance Amount
}

type Transfer struct {
	ID          ID
	FromAccount ID
	ToAccount   ID
	Balance     Amount
	Date        Date
}

type Envelope struct {
	ID      ID
	Name    string
	Limit   Amount
	Balance Amount
}

type Plan struct {
	ID           ID
	FromEnvelope ID
	ToEnvelope   ID
	Balance      Amount
	Executed     Amount
	Date         Date
}

type Category struct {
	ID          ID
	EnvelopeID  ID
	Name        string
	Description string
}

type Storage interface {
	GetMonthlyBudget(budgetID ID, monthlyBudgetID ID) (*MonthlyBudget, error)
	GetBudget(budgetID ID) (*Budget, error)
	GetAccount(budgetID ID, accountID ID) (*Account, error)
	GetEnvelope(budgetID ID, envelopeID ID) (*Envelope, error)
	GetCurrentMonthlyBudget(budgetID ID) (*MonthlyBudget, error)
	GetCurrentExpensesForAccount(budgetID ID, monthlyBudgetID ID, accountID ID) ([]*Expense, error)
	EnsureMonthlyBudget(budgetID ID, month time.Month, year uint) (*MonthlyBudget, error)
	UpdateMonthlyBudget(budgetID ID, monthlyBudget *MonthlyBudget) (MonthlyBudget, error)
	UpdateBudget(budget *Budget) (*Budget, error)
}

// Better to be done in memory
func CloseMonthlyBudget(budgetID ID, monthlyBudgetID ID, storage Storage) error {
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
		fromAccount := budget.Account(transfer.FromAccount)
		fromAccount.Balance = fromAccount.Balance.Sub(transfer.Balance)

		toAccount := budget.Account(transfer.ToAccount)
		toAccount.Balance = toAccount.Balance.Add(transfer.Balance)
	}

	// processPlans
	for _, plan := range monthlyBudget.Plans {
		fromEnvelope := budget.Envelope(plan.FromEnvelope)
		fromEnvelope.Balance = fromEnvelope.Balance.Sub(plan.Balance)

		toEnvelope := budget.Envelope(plan.ToEnvelope)
		toEnvelope.Balance = toEnvelope.Balance.Add(plan.Balance)
		plan.Executed = plan.Balance
		if toEnvelope.Balance.IsBiggerThan(toEnvelope.Limit) {
			plan.Executed = plan.Balance - (toEnvelope.Balance - toEnvelope.Limit)
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
func GetAccount(budgetID ID, accountID ID, storage Storage) (*Account, error) {
	account, err := storage.GetAccount(budgetID, accountID)
	if err != nil {
		return nil, err
	}

	monthlyBudget, err := storage.GetCurrentMonthlyBudget(budgetID)
	if err != nil {
		return nil, err
	}

	for _, expense := range monthlyBudget.Expenses {
		if expense.AccountID == accountID {
			account.Balance = account.Balance.Sub(expense.Balance())
		}
	}

	for _, transfer := range monthlyBudget.Transfers {
		if transfer.FromAccount == accountID {
			account.Balance = account.Balance.Sub(transfer.Balance)
		}
		if transfer.ToAccount == accountID {
			account.Balance = account.Balance.Add(transfer.Balance)
		}
	}
	return account, nil
}


// Should be possible to do as one DB operation
func GetEnvelope(budgetID ID, envelopeID ID, storage Storage) (*Envelope, error) {
	budget, err := storage.GetBudget(budgetID)
	if err != nil {
		return nil, err
	}

	envelope, err := storage.GetEnvelope(budgetID, envelopeID)
	if err != nil {
		return nil, err
	}

	monthlyBudget, err := storage.GetCurrentMonthlyBudget(budgetID)
	if err != nil {
		return nil, err
	}

	for _, expense := range monthlyBudget.Expenses {
		for _, expenseCategory := range expense.Categories {
			if budget.Category(expenseCategory.CategoryID).EnvelopeID == envelopeID {
				envelope.Balance = envelope.Balance.Sub(expenseCategory.Balance)
			}
		}
	}

	for _, plan := range monthlyBudget.Plans {
		if plan.FromEnvelope == envelopeID {
			envelope.Balance = envelope.Balance.Sub(plan.Balance)
		}
		if plan.ToEnvelope == envelopeID {
			envelope.Balance = envelope.Balance.Add(plan.Balance)
		}
	}
	return envelope, nil
}
