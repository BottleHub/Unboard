package main

import (
	"log"
	"os"

	"github.com/bottlehub/unboard/boards/graph"
	"github.com/bottlehub/unboard/boards/internal/mq"
	"github.com/bottlehub/unboard/boards/internal/routes"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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
	go mq.Consume("TestQueue")
	handle := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		handle.ServeHTTP(c.Writer, c.Request)
	}
}

// Starts the server process
func main() {
	ch := make(chan bool)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//gin.SetMode(gin.ReleaseMode)

	route := gin.Default()

	go routes.Route(route)

	go route.GET("/")
	go route.POST("/query", graphqlHandler())
	go route.GET("/graphql", playgroundHandler())

	go log.Printf("Connect to http://localhost:%s/graphql for GraphQL playground", port)
	go log.Fatal(route.Run(":" + port))
	<-ch
}
