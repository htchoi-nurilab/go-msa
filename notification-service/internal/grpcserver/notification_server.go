package grpcserver

import (
	"context"

	"github.com/htchoi-nurilab/go-msa/notification-service/internal/service"
	pb "github.com/htchoi-nurilab/go-msa/notification-service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
	notificationSvc *service.NotificationService
}

func NewServer(notificationSvc *service.NotificationService) *Server {
	return &Server{notificationSvc: notificationSvc}
}

func (s *Server) CreateNotification(ctx context.Context, request *pb.NotificationCreateRequest) (*emptypb.Empty, error) {
	if err := s.notificationSvc.CreateWelcomeNotification(
		ctx,
		uint(request.UserId),
		request.Name,
	); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
