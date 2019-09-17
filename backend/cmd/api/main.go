package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/handlers"
	"github.com/sjanota/budget/backend/pkg/resolver"
	"github.com/sjanota/budget/backend/pkg/schema"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoUrl := os.Getenv("MONGODB_URL")
	if mongoUrl == "" {
		log.Fatal("Missing required MONGODB_URL env")
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	h := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
	)(
		handler.GraphQL(schema.NewExecutableSchema(schema.Config{Resolvers: &resolver.Resolver{}})),
	)
	http.Handle("/query", h)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
