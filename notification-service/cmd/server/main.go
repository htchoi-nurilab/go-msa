package main

import (
	"flag"
	"fmt"
	"net"

	"log"

	"github.com/htchoi-nurilab/go-msa/notification-service/config/database"
	"github.com/htchoi-nurilab/go-msa/notification-service/internal/domain"
	"github.com/htchoi-nurilab/go-msa/notification-service/internal/grpcserver"
	"github.com/htchoi-nurilab/go-msa/notification-service/internal/repository"
	"github.com/htchoi-nurilab/go-msa/notification-service/internal/service"
	"github.com/htchoi-nurilab/go-msa/notification-service/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(".env not found, using system env")
	}

	db := database.NewDatabase()
	if err := db.DB().AutoMigrate(&domain.Notification{}); err != nil {
		log.Fatal(err)
	}

	notificationRepo := repository.NewNotificationRepository(db.DB())
	notificationSvc := service.NewNotificationService(notificationRepo)
	notificationServer := grpcserver.NewServer(notificationSvc)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	notificationpb.RegisterNotificationServiceServer(grpcSrv, notificationServer)

	log.Printf("notification gRPC server listening at %v", lis.Addr())
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
