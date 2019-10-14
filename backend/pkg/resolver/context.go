package resolver

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func budgetFromContext(ctx context.Context) primitive.ObjectID {
	resolverCtx := graphql.GetResolverContext(ctx)
	var budgetID primitive.ObjectID
	var ok bool
	for resolverCtx != nil {
		budgetID, ok = resolverCtx.Args["budgetID"].(primitive.ObjectID)
		if ok {
			break
		}
		resolverCtx = resolverCtx.Parent
	}
	return budgetID
}
