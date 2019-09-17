package main

import (
	"github.com/sjanota/budget/backend/pkg/storage"
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

	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {
		log.Fatal("Missing required MONGODB_URI env")
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	storage, err := storage.New(mongoUri)
	if err != nil {
		log.Fatalf("Couldn't create storate: %s", err)
	}

	resolver := &resolver.Resolver{Storage: storage}

	h := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
	)(
		handler.GraphQL(schema.NewExecutableSchema(schema.Config{Resolvers: resolver})),
	)
	http.Handle("/query", h)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
