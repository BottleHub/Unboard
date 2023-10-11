package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bottlehub/unboard/boards/graph"
	"github.com/bottlehub/unboard/boards/internal/mq"
	"github.com/bottlehub/unboard/boards/internal/router"
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8090"

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
	ch := make(chan bool)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//gin.SetMode(gin.ReleaseMode)

	route := gin.Default()

	go router.Route(route)

	go route.GET("/")
	go route.POST("/query", graphqlHandler())
	go route.GET("/graphql", playgroundHandler())
	go mq.Consume("TestQueue", func(s string) {
		fmt.Println(s)
	})
	// go mq.Consume("CreateBoard", func(s string) {
	// 	// Marshal string to json
	// 	t, err := json.Marshal(s)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// Get desc if available
	// 	desc := ""
	// 	if string(t[2]) != "" {
	// 		desc = string(t[2])
	// 	}

	// 	// Create new model
	// 	d := model.NewChatboard{
	// 		Name:        string(t[0]),
	// 		ImageURL:    string(t[1]),
	// 		Description: &desc,
	// 	}

	// 	// Create new context and resolver
	// 	ctx := new(context.Context)
	// 	res := graph.Resolver{}
	// 	// Create new request using context and model
	// 	_, err = res.Mutation().CreateChatboard(*ctx, d)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// })

	go log.Printf("Connect to http://localhost:%s/graphql for GraphQL playground", port)
	go log.Fatal(route.Run(":" + port))
	<-ch
}
