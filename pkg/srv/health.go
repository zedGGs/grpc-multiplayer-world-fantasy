package srv

import (
	"context"
	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type healthService struct {
	pb.UnimplementedHealthServiceServer
}

func NewHealthServiceServer() pb.HealthServiceServer {
	return &healthService{}
}

func (s *healthService) Health(context.Context, *emptypb.Empty) (*pb.HealthMessage, error) {
	return &pb.HealthMessage{Status: "ok"}, nil
}
