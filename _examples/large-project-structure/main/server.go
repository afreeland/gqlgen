package main

import (
	"fmt"
	"github.com/99designs/gqlgen/_examples/large-project-structure/integration"
	public_graph "github.com/99designs/gqlgen/_examples/large-project-structure/main/public/graph"
	"log"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"syscall"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/djang0man/sdui/copyschemas"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/ast"

	private_int_graph "github.com/99designs/gqlgen/_examples/large-project-structure/integration/private/graph"
	private_graph "github.com/99designs/gqlgen/_examples/large-project-structure/main/private/graph"
)

const defaultPort = "8080"
const privatePort = "8081"

var publicExecutableSchema graphql.ExecutableSchema

func main() {
	copiedFiles, err := copyschemas.CopyGraphqlSchemas("../")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	for _, path := range copiedFiles {
		fmt.Println(path)
	}

	publicPort := os.Getenv("PORT")
	if publicPort == "" {
		publicPort = defaultPort
	}

	tmpDir := os.TempDir()
	inputDir, err := os.MkdirTemp(tmpDir, "input-*")
	if err != nil {
		panic(err)
	}
	// Ensure the file is cleaned up
	defer os.RemoveAll(inputDir)

	// Set up a channel to listen for termination signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle signals
	go func() {
		<-sigs
		log.Println("Signal received, cleaning up...")
		os.RemoveAll(inputDir)
		os.Exit(0)
	}()

	fmt.Printf("Looking for INPUT FILES: %s\n", inputDir)

	// Create a new logger instance
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.WithFields(logrus.Fields{
		"namespace": "main",
	})

	resolver := public_graph.NewResolver(public_graph.NewExecutableSchema(public_graph.Config{
		Resolvers: &public_graph.Resolver{
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

	publicExecutableSchema = resolver.GetExecutableSchema()

	if publicExecutableSchema == nil {
		log.Fatal("Executable schema is nil")
	}

	parsedSchema := resolver.GetParsedSchema()
	if parsedSchema == nil {
		log.Fatal("Parsed schema is nil")
	} else {
		log.Println("Parsed Schema: ", parsedSchema)
	}

	var AppConnectorMethods []string
	var AppConnectorInterfaceTypes []string

	for key, obj := range parsedSchema.Types {
		for _, iface := range obj.Interfaces {
			if iface == "AppConnector" {
				AppConnectorInterfaceTypes = append(AppConnectorInterfaceTypes, key)
				break
			}
		}
	}

	for _, obj := range parsedSchema.Query.Fields {
		if slices.Contains(AppConnectorInterfaceTypes, obj.Type.NamedType) {
			AppConnectorMethods = append(AppConnectorMethods, obj.Name)
		}
	}

	if len(AppConnectorMethods) > 0 {
		public_graph.PublicAppConnectorMethods = AppConnectorMethods
	}

	// Create a new executable schema with the composed resolver
	publicSrv := handler.NewDefaultServer(publicExecutableSchema)

	publicSrv.AddTransport(transport.Options{})
	publicSrv.AddTransport(transport.GET{})
	publicSrv.AddTransport(transport.POST{})

	publicSrv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	publicSrv.Use(extension.Introspection{})
	publicSrv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Create a new servemux for the public server
	publicMux := http.NewServeMux()
	publicMux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	publicMux.Handle("/query", publicSrv)

	// Create a new executable schema with the composed resolver
	privateSrv := handler.NewDefaultServer(private_graph.NewExecutableSchema(private_graph.Config{
		Resolvers: &private_graph.Resolver{
			PrivateIntegration: &private_int_graph.Resolver{
				InputDir: inputDir,
			},
		},
	}))

	privateSrv.AddTransport(transport.Options{})
	privateSrv.AddTransport(transport.GET{})
	privateSrv.AddTransport(transport.POST{})

	privateSrv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	privateSrv.Use(extension.Introspection{})
	privateSrv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Create a new servemux for the private server
	privateMux := http.NewServeMux()
	privateMux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	privateMux.Handle("/query", privateSrv)

	// start our private server
	go func() {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", privatePort)
		log.Fatal(http.ListenAndServe(":"+privatePort, privateMux))
	}()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", publicPort)
	log.Fatal(http.ListenAndServe(":"+publicPort, publicMux))
}
