package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Budget struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	Accounts     []*Account         `json:"accounts"`
	Envelopes    []*Envelope        `json:"envelopes"`
	Categories   []*Category        `json:"categories"`
	CurrentMonth Month
}

type Changes map[string]interface{}

func (ch Changes) Has(key string) bool {
	_, ok := ch[key]
	return ok
}

func (ch Changes) GetID(key string) primitive.ObjectID {
	return ch[key].(primitive.ObjectID)
}

func (ch Changes) GetDate(key string) Date {
	return ch[key].(Date)
}

func (b Budget) Category(id primitive.ObjectID) *Category {
	for _, category := range b.Categories {
		if category.ID == id {
			return category
		}
	}
	return nil
}

func (b Budget) Account(id primitive.ObjectID) *Account {
	for _, account := range b.Accounts {
		if account.ID == id {
			return account
		}
	}
	return nil
}

func (b Budget) Envelope(id primitive.ObjectID) *Envelope {
	for _, envelope := range b.Envelopes {
		if envelope.ID == id {
			return envelope
		}
	}
	return nil
}

func (b Budget) CurrentMonthID() MonthlyReportID {
	return MonthlyReportID{
		Month:    b.CurrentMonth,
		BudgetID: b.ID,
	}
}

type MonthlyReport struct {
	ID        MonthlyReportID `bson:"_id,omitempty"`
	Expenses  []*Expense      `json:"expenses"`
	Transfers []*Transfer     `json:"transfers"`
	Plans     []*Plan         `json:"plans"`
}

func (r MonthlyReport) Expense(id primitive.ObjectID) *Expense {
	for _, expense := range r.Expenses {
		if expense.ID == id {
			return expense
		}
	}
	return nil
}

func (r MonthlyReport) Transfer(id primitive.ObjectID) *Transfer {
	for _, transfer := range r.Transfers {
		if transfer.ID == id {
			return transfer
		}
	}
	return nil
}

func (r MonthlyReport) Plan(id primitive.ObjectID) *Plan {
	for _, plan := range r.Plans {
		if plan.ID == id {
			return plan
		}
	}
	return nil
}

func (r MonthlyReport) Month() Month {
	return r.ID.Month
}

func (r MonthlyReport) WithBudget(budget Budget) MonthlyReport {
	r.ID = *r.ID.WithBudget(budget).WithMonth(budget.CurrentMonth)
	return r
}

type MonthlyReportID struct {
	Month    Month
	BudgetID primitive.ObjectID
}

func (id MonthlyReportID) WithBudget(budget Budget) *MonthlyReportID {
	id.BudgetID = budget.ID
	return &id
}

func (id MonthlyReportID) WithMonth(month Month) *MonthlyReportID {
	id.Month = month
	return &id
}

type Expense struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title      string             `json:"title"`
	Categories []*ExpenseCategory `json:"categories"`
	Date       Date               `json:"date"`
	AccountID  primitive.ObjectID
}

func (e Expense) TotalAmount() Amount {
	var sum Amount
	for _, c := range e.Categories {
		sum = sum.Add(c.Amount)
	}
	return sum
}

func (e Expense) WithCategories(categories ...*ExpenseCategory) *Expense {
	e.Categories = categories
	return &e
}

type ExpenseCategory struct {
	Amount     Amount `json:"balance"`
	CategoryID primitive.ObjectID
}

func (c ExpenseCategory) WithAmount(amount Amount) *ExpenseCategory {
	c.Amount = amount
	return &c
}

type Account struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Balance Amount
}

type Transfer struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount        Amount             `json:"amount"`
	Title         string             `json:"title"`
	Date          Date               `json:"date"`
	FromAccountID *primitive.ObjectID
	ToAccountID   primitive.ObjectID
}

type Envelope struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Limit   *Amount            `json:"limit"`
	Balance Amount
}

type Plan struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title           string             `json:"title"`
	CurrentAmount   Amount             `json:"currentAmount"`
	RecurringAmount *Amount            `json:"recurringAmount"`
	FromEnvelopeID  *primitive.ObjectID
	ToEnvelopeID    primitive.ObjectID
}

type Category struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name"`
	EnvelopeID primitive.ObjectID
}

func (c Category) WithEnvelope(envelopeID primitive.ObjectID) Category {
	c.EnvelopeID = envelopeID
	return c
}

func (i ExpenseInput) WithDate(date Date) *ExpenseInput {
	i.Date = date
	return &i
}

func (i ExpenseInput) WithAccount(accountID primitive.ObjectID) *ExpenseInput {
	i.AccountID = accountID
	return &i
}

func (i ExpenseInput) WithCategories(categories ...*ExpenseCategoryInput) *ExpenseInput {
	i.Categories = categories
	return &i
}

func (i ExpenseCategoryInput) WithCategory(categoryID primitive.ObjectID) *ExpenseCategoryInput {
	i.CategoryID = categoryID
	return &i
}

func (u CategoryUpdate) Changes() Changes {
	result := make(map[string]interface{})
	if u.Name != nil {
		result["name"] = *u.Name
	}
	if u.EnvelopeID != nil {
		result["envelopeid"] = *u.EnvelopeID
	}
	return result
}

func (u CategoryUpdate) WithEnvelope(envelopeID primitive.ObjectID) *CategoryUpdate {
	u.EnvelopeID = &envelopeID
	return &u
}

func (u ExpenseUpdate) Changes() Changes {
	result := make(map[string]interface{})
	if u.Title != nil {
		result["title"] = *u.Title
	}
	if u.AccountID != nil {
		result["accountid"] = *u.AccountID
	}
	if u.Date != nil {
		result["date"] = *u.Date
	}
	if u.Categories != nil {
		result["categories"] = u.Categories
	}
	return result
}

type PlanUpdate Changes
func NewPlanUpdate(changes map[string]interface{}) (PlanUpdate, error) {
	var err error
	result := make(map[string]interface{})
	for key, value := range changes {
		switch key {
		case "fromEnvelopeID":
			result["fromenvelopeid"], err = MaybeUnmarshalID(value)
			break
		case "toEnvelopeID":
			result["toenvelopeid"], err = UnmarshalID(value)
			break
		default:
			result[strings.ToLower(key)] = value
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func TransferChanges(changes map[string]interface{}) (Changes, error) {
	var err error
	result := make(map[string]interface{})
	for key, value := range changes {
		switch key {
		case "fromEnvelopeID":
			result["fromaccountid"], err = MaybeUnmarshalID(value)
			break
		case "toEnvelopeID":
			result["toaccountid"], err = UnmarshalID(value)
			break
		default:
			result[strings.ToLower(key)] = value
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (r MonthlyReport) TotalPlannedAmount() Amount {
	var sum Amount
	for _, p := range r.Plans {
		if p.FromEnvelopeID == nil {
			sum = sum.Add(p.CurrentAmount)
		}
	}
	return sum
}

func (r MonthlyReport) TotalIncomeAmount() Amount {
	var sum Amount
	for _, t := range r.Transfers {
		if t.FromAccountID == nil {
			sum = sum.Add(t.Amount)
		}
	}
	return sum
}

func (r MonthlyReport) TotalExpenseAmount() Amount {
	var sum Amount
	for _, e := range r.Expenses {
		sum = sum.Add(e.TotalAmount())
	}
	return sum
}

func (r MonthlyReport) ApplyTo(budget *Budget) {
	for _, expense := range r.Expenses {
		account := budget.Account(expense.AccountID)
		for _, expenseCategory := range expense.Categories {
			category := budget.Category(expenseCategory.CategoryID)
			envelope := budget.Envelope(category.EnvelopeID)

			account.Balance = account.Balance.Sub(expenseCategory.Amount)
			envelope.Balance = envelope.Balance.Sub(expenseCategory.Amount)
		}
	}

	for _, transfer := range r.Transfers {
		if transfer.FromAccountID != nil {
			fromAccount := budget.Account(*transfer.FromAccountID)
			fromAccount.Balance = fromAccount.Balance.Sub(transfer.Amount)
		}

		toAccount := budget.Account(transfer.ToAccountID)
		toAccount.Balance = toAccount.Balance.Add(transfer.Amount)
	}

	for _, plan := range r.Plans {
		if plan.FromEnvelopeID != nil {
			fromEnvelope := budget.Envelope(*plan.FromEnvelopeID)
			fromEnvelope.Balance = fromEnvelope.Balance.Sub(plan.CurrentAmount)
		}

		toEnvelope := budget.Envelope(plan.ToEnvelopeID)
		toEnvelope.Balance = toEnvelope.Balance.Add(plan.CurrentAmount)
	}
}