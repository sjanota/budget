package main

import (
	"context"
	"encoding/json"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sjanota/budget/backend/pkg/schema"
	"github.com/sjanota/budget/backend/pkg/storage"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/sjanota/budget/backend/pkg/resolver"
)

const defaultPort = "8080"

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

const (
	issuer   = "https://damp-pond-6290.eu.auth0.com/"
	audience = "https://backend/api"
)

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

	r := mux.NewRouter()
	r.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"authorization", "content-type"}),
	))

	authDisabled := os.Getenv("INSECURE_AUTH_DISABLED")
	if authDisabled != "true" {
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				// Verify 'aud' claim
				checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
				if !checkAud {
					return token, errors.New("Invalid audience.")
				}
				// Verify 'iss' claim
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
				if !checkIss {
					return token, errors.New("Invalid issuer.")
				}

				// Do not check iat as auth0 server is in different timezone
				delete(map[string]interface{}(token.Claims.(jwt.MapClaims)), "iat")

				cert, err := getPemCert(token)
				if err != nil {
					panic(err.Error())
				}

				result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
				return result, nil
			},
			SigningMethod: jwt.SigningMethodRS256,
		})
		r.Use(jwtMiddleware.Handler)
	} else {
		log.Print("Running with disabled authentication")
	}

	gqlHandler := handler.GraphQL(
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
	)
	r.Handle("/query", gqlHandler)
	r.Handle("/", handler.Playground("Budget", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(issuer + ".well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
