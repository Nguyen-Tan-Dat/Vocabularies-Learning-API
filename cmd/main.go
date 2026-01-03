package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/db"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/resolvers"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/middleware"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/repositories"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/services"
	"github.com/vektah/gqlparser/v2/ast"
)

//var db *gorm.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	db.ConnectDatabase()
	topicService := services.TopicService{Repo: repositories.TopicRepository{DB: db.GetDB()}}
	englishService := services.EnglishService{Repo: repositories.EnglishRepository{DB: db.GetDB()}}
	resolver := resolvers.Resolver{TopicService: topicService, EnglishService: englishService}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	// Add middleware
	http.Handle("/graphql", middleware.JWTMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
