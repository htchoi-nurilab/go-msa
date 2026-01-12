package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "user-service",
			"status":  "ok",
		})
	})

	log.Println("user-service started on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
