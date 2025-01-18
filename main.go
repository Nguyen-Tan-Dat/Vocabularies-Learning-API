// main.go
package main

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nguyen-Tan-Dat/Vocabualries-Learning-API/graph"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// For non-playground requests, use handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	r.GET("/", playground.Handler("GraphQL playground", "/query"))
	r.POST("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})))

	r.Run(":8080")
}
