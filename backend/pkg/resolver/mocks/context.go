package mock_resolver

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MockContext(budgetID primitive.ObjectID) context.Context {
	rctx := &graphql.ResolverContext{Args: map[string]interface{}{"budgetID": budgetID}}
	return graphql.WithResolverContext(context.TODO(), rctx)
}
