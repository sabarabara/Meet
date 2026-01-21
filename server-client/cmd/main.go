package main

import (
	"log"
	"server-client/pkg/db"

	gen "server-client/internal/presenter/dto/gql/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	r := gin.Default()

	app := InitializeApp(db.DB)

	srv := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{
		Resolvers: app.Resolver,
	}))
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
	log.Println("Server started on :8080")
}
