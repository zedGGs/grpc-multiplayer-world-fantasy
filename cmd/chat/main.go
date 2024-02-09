package main

import (
	"context"
	chat "github.com/zedGGs/grpc-multiplayer-world-fantasy/cmd/chat/app"
	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/helpers"
	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/pb"
	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/srv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/config"
	"github.com/uptrace/uptrace-go/uptrace"
)

var (
	conf *config.GlobalConfig
)

func init() {
	helpers.SetupLogger()
	conf = config.NewGlobalConfig()
}

func main() {
	ctx := context.Background()
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(conf.Uptrace.DSN),
		uptrace.WithServiceName(chat.ServiceName),
		uptrace.WithServiceVersion(conf.Version),
	)

	server := chat.NewServerContext(ctx, conf)
	grpcServer, gwmux := helpers.InitServerDefaults()
	address := server.GlobalConfig.Chat.Local.Address()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	pb.RegisterHealthServiceServer(grpcServer, srv.NewHealthServiceServer())
	err := pb.RegisterHealthServiceHandlerFromEndpoint(ctx, gwmux, address, opts)
	helpers.Check(ctx, err, "register health service handler endpoint")

	srvService, err := srv.NewChatServiceServer(ctx, server)
	helpers.Check(ctx, err, "create chat service")
	pb.RegisterChatServiceServer(grpcServer, srvService)
	err = pb.RegisterChatServiceHandlerFromEndpoint(ctx, gwmux, address, opts)
	helpers.Check(ctx, err, "register chat service handler endpoint")

	helpers.StartServer(ctx, grpcServer, gwmux, server.GlobalConfig.Chat.Local.Address())
}
