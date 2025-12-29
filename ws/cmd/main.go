package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}

	log.Println("Server started on :8081")
}
