package main

import (
	"log"
	"os"

	"github.com/bottlehub/unboard/configs/auth"
	"github.com/bottlehub/unboard/graph"
	"github.com/bottlehub/unboard/routes"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

// GraphQL handle helper
func graphqlHandler() gin.HandlerFunc {
	handle := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		handle.ServeHTTP(c.Writer, c.Request)
	}
}

// Redirects to fetching the graphql handle
func playgroundHandler() gin.HandlerFunc {
	handle := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		handle.ServeHTTP(c.Writer, c.Request)
	}
}

// Starts the server process
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware("phrase"))

	route := gin.Default()

	route.GET("/")
	routes.Route(route)

	route.POST("/query", graphqlHandler())
	route.GET("/graphql", playgroundHandler())

	log.Printf("Connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(route.Run(":" + port))
}
