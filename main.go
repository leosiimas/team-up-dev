package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
	"github.com/leosiimas/team-up-dev/graph/generated"
	"github.com/leosiimas/team-up-dev/resolvers"
)

const (
	port = ":8080"
)

func lineSeparator() {
	fmt.Println("========")
}

func startMessage() {
	lineSeparator()
	color.Green("Listening on localhost%s\n", port)
	color.Green("Visit `http://localhost%s/graphql` in your browser\n", port)
	lineSeparator()
}

func main() {
	lineSeparator()
	// Create the database `todos` manually within postgres
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456",
		Database: "todos",
	})
	defer db.Close()

	r := chi.NewRouter()

	// The base path that users would use is POST /graphql which is fairly
	// idiomatic.
	r.Route("/graphql", func(r chi.Router) {

		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolvers.Resolver{
				DB: db,
			},
		})

		srv := handler.NewDefaultServer(schema)
		srv.Use(extension.FixedComplexityLimit(300))

		r.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("api-gateway", "/graphql")
	r.Get("/", gqlPlayground)

	startMessage()
	panic(http.ListenAndServe(port, r))
}
