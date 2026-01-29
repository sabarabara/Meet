package main

import (
	"log"
	"server-client/pkg/db"

	gen "server-client/internal/presenter/dto/gql/graph/generated"
	"server-client/internal/presenter/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	db.InitRedis()
	rClient := db.InitRedis()

	app, err := InitializeApp(db.SqlDB, db.GormDB, rClient)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	srv := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{
		Resolvers: app.Resolver,
	}))

	r := gin.Default()
	r.Use(middleware.ProxyHeaderMiddleware())

	auth := r.Group("/auth")
	{
		auth.GET("/login", app.LoginHandler.Login)
		auth.GET("/callback", app.LoginHandler.Callback)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	protected := r.Group("/")
	protected.Use(app.AuthMiddleware.AuthenticateMiddleware())
	{
		protected.POST("/query", func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)
		})

		protected.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
		})
	}

	if err := r.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
	log.Println("Server started on :8080")
}
