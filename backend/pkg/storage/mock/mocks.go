package mock

import (
	"math/rand"
	"time"

	"github.com/sjanota/budget/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Name() string {
	return primitive.NewObjectID().Hex()
}

func Amount() *models.Amount {
	return &models.Amount{
		Integer: rand.Int(),
		Decimal: rand.Int() % 100,
	}
}

func Month() time.Month {
	return time.Month(rand.Int()%12 + 1)
}

func Year() int {
	return rand.Int()%50 + 1990
}
