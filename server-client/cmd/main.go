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
	db.InitRedis()
	rClient := db.InitRedis()
	r := gin.Default()

	app, err := InitializeApp(db.SqlDB, db.GormDB, rClient)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

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

	//OIDC認証用エンドポイント
	r.GET("/auth/login", func(c *gin.Context) {
		app.LoginHandler.Login(c.Writer, c.Request)
	})

	r.GET("/auth/callback", func(c *gin.Context) {
		app.LoginHandler.Callback(c.Writer, c.Request)
	})

	if err := r.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
	log.Println("Server started on :8080")
}
