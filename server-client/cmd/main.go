package main

import (
	"log"
	"server-client/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	r := gin.Default()

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
