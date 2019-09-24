package main

import (
	"context"
	"fmt"
	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/sjanota/budget/backend/pkg/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"sync"
	"time"
)

var (
	accountsConfig = map[int]int{
		10: 2000,
		20: 200,
	}
)

func main() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("Missing required MONGODB_URI env")
	}

	storage, err := storage.New(mongoURI)
	if err != nil {
		log.Fatalf("Couldn't create storate: %s", err)
	}

	ctx := context.Background()
	storage.Drop(ctx)
	storage.Init(ctx)

	budget, err := storage.Budgets().Create(ctx, "test")
	if err != nil {
		log.Fatalf("cannot create budget: %s", err)
	}

	for nOfAccounts, nOfExpenses := range accountsConfig {
		for i := 0; i < nOfAccounts; i++ {
			accountID := primitive.NewObjectID()
			wg := sync.WaitGroup{}
			wg.Add(nOfExpenses)
			for j := 0; j < nOfExpenses; j++ {
				go func() {
					defer wg.Done()

					_, err = storage.Expenses(budget.ID).Insert(ctx, models.ExpenseInput{
						Title:    generateName(),
						Location: nil,
						Entries:  nil,
						TotalBalance: &models.MoneyAmountInput{
							Integer: 1,
							Decimal: 0,
						},
						Date:      nil,
						AccountID: &accountID,
					})
					if err != nil {
						log.Printf("cannot create expense: %s", err)
					}
				}()
			}
			wg.Wait()
			took := measure(func() {
				rsp, err := storage.Expenses(budget.ID).TotalBalanceForAccount(ctx, accountID)
				if err != nil {
					log.Fatalf("cannot calc total: %s", err)
				}
				log.Println(rsp)
			})
			fmt.Println(took)
		}
	}

}

func generateName() string {
	return primitive.NewObjectID().Hex()
}

func measure(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Now()
	return end.Sub(start)

}
