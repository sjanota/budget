package resolver

import (
	"context"

	"github.com/sjanota/budget/backend/pkg/models"
)

type envelopeResolver struct {
	*Resolver
}

func (e envelopeResolver) Balance(ctx context.Context, obj *models.Envelope) (*models.MoneyAmount, error) {
	panic("implement me")
}
