package main

import (
	"log"
	"net/http"
	"os"

	"github.com/LEEDASILVA/graphQLServer/go/internal/auth"

	"github.com/99designs/gqlgen/handler"
	hackernews "github.com/LEEDASILVA/graphQLServer/go/graph/generated"
	db "github.com/LEEDASILVA/graphQLServer/go/internal/pkg/db/mysql"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter()

	r.Use(auth.Middleware())

	db.InitDB()
	db.Migrate()

	server := handler.GraphQL(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}}))
	r.Handle("/", handler.Playground("GraphQL playground", "/query"))
	r.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
