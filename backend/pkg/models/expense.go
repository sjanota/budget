package models

type Expense struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Location  *string         `json:"location"`
	Entries   []*ExpenseEntry `json:"entries"`
	Total     *MoneyAmount    `json:"total"`
	Date      *string         `json:"date"`
	AccountID *string         `json:"account"`
}
