package main

import (
	"flag"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/htchoi-nurilab/go-msa/notification-service/proto"
	"github.com/htchoi-nurilab/go-msa/user-service/config/database"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/client"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/domain"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/handler"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/repository"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(".env not found, using system env")
	}

	db := database.NewDatabase()
	if err := db.DB().AutoMigrate(&domain.User{}); err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	conn, err := grpc.NewClient(
		*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	pbClient := notificationpb.NewNotificationServiceClient(conn)
	notiClient := client.NewGrpcNotificationClient(pbClient, 2*time.Second)

	userRepo := repository.NewUserRepository(db.DB())
	userSvc := service.NewUserService(userRepo, notiClient)
	userHandler := handler.NewUserHandler(userSvc)

	r := gin.Default()
	userHandler.RegisterRoutes(r)

	r.Run(":8081")
}
