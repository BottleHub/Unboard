package main

import (
	"log"
	"os"

	"github.com/bottlehub/unboard/boards/graph"
	"github.com/bottlehub/unboard/boards/internals/mq"
	"github.com/bottlehub/unboard/boards/internals/routes"
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
	handle := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		handle.ServeHTTP(c.Writer, c.Request)
	}
}

// Starts the server process
func main() {
	ch := make(chan string, 3)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//gin.SetMode(gin.ReleaseMode)

	//router := chi.NewRouter()
	//router.Use(auth.Middleware("phrase"))

	route := gin.Default()

	routes.Route(route)

	route.GET("/")
	route.POST("/query", graphqlHandler())
	route.GET("/graphql", playgroundHandler())
	go mq.Consume()

	go log.Printf("Connect to http://localhost:%s/graphql for GraphQL playground", port)
	go log.Fatal(route.Run(":" + port))
	<-ch
}
