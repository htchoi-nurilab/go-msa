package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/htchoi-nurilab/go-msa/user-service/config/database"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/domain"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/handler"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/repository"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(".env not found, using system env")
	}

	db := database.NewDatabase()
	if err := db.DB().AutoMigrate(&domain.User{}); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db.DB())
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()
	userHandler.RegisterRoutes(r)

	r.Run(":8081")
}
