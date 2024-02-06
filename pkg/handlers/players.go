package srv

import (
	characters "github.com/zedGGs/grpc-multiplayer-world-fantasy/cmd/players/app"
	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/pb"
)

type playersServiceServer struct {
	pb.UnimplementedCharacterServiceServer
	server *characters.CharactersServerContext
}