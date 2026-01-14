package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/htchoi-nurilab/go-msa/notification-service/config/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(".env not found, using system env")
	}

	db := database.NewDatabase()

	r := gin.Default()

	r.Run(":8082")
}
