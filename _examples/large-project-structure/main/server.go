package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/_examples/large-project-structure/main/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/_examples/large-project-structure/integration"
	_ "github.com/99designs/gqlgen/_examples/large-project-structure/shared"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create a new logger instance
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.WithFields(logrus.Fields{
		"namespace": "main",
	})

	// Create a new executable schema with the composed resolver
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			ExternalQueryResolver: &integration.Resolver{
				Logger: func(log *logrus.Logger) *logrus.Entry {
					// Clone the mainLoggerEntry with a different namespace
					return log.WithFields(logrus.Fields{
						"namespace": "integration",
					})
				}(log),
			},
			// Add other team resolvers here
		},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
