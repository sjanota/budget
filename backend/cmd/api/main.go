package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sjanota/budget/backend/pkg/storage"

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

	mongoURI := "mongodb://localhost:32768/budget"
	//mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("Missing required MONGODB_URI env")
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))

	storage, err := storage.New(mongoURI)
	if err != nil {
		log.Fatalf("Couldn't create storate: %s", err)
	}

	err = storage.Init(context.Background())
	if err != nil {
		log.Fatalf("Couldn't init storate: %s", err)
	}

	resolver := &resolver.Resolver{Storage: storage}

	h := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
	)(
		handler.GraphQL(
			schema.NewExecutableSchema(schema.Config{Resolvers: resolver}),
			handler.WebsocketKeepAliveDuration(10*time.Second),
			handler.WebsocketUpgrader(
				websocket.Upgrader{
					ReadBufferSize:  1024,
					WriteBufferSize: 1024,
					CheckOrigin: func(r *http.Request) bool {
						return true
					},
				},
			),
		),
	)
	http.Handle("/query", h)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
