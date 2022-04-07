package main

import (
	"log"
	"net/http"

	db "graphql-golang/internal/pkg/db/mysql"
	"graphql-golang/schema"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/graphql-go/handler"
)

const defaultPort = "8080"

func main() {
	port := defaultPort
	db.InitDB()
	db.Migrate()

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
