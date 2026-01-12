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
			"service": "notification-service",
			"status":  "ok",
		})
	})

	log.Println("notification-service started on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
