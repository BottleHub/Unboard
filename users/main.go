package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bottlehub/unboard/users/graph"
	"github.com/bottlehub/unboard/users/internal/mq"
	"github.com/bottlehub/unboard/users/internal/router"
	"github.com/gin-gonic/gin"
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
	ch := make(chan string, 7)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	go router.Route(r)

	go r.GET("/")
	go r.POST("/query", graphqlHandler())
	go r.GET("/graphql", playgroundHandler())
	go mq.Consume()

	go log.Printf("Connect to http://localhost:%s/graphql for GraphQL playground", port)
	go log.Fatal(r.Run(":" + port))
	<-ch
}
