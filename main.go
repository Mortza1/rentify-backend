package main

import (
    "log"
    // "os"
    "rentify-backend/graph"
    "rentify-backend/graph/generated"
    "rentify-backend/database"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    database.Connect()

    r := gin.Default()

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
    r.POST("/query", gin.WrapH(srv))
    r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

    log.Println("Server is running on http://localhost:8080")
    r.Run(":8080")
}
