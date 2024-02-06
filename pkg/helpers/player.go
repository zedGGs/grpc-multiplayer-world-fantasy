package helpers

import (
	"context"
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/model"
	"github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/pb"
)

func GetCharacterIdFromTarget(
	ctx context.Context,
	charactersServiceClient pb.CharacterServiceClient,
	target *pb.CharacterTarget,
) (uint, error) {
	if target == nil {
		return 0, fmt.Errorf("target cannot be nil")
	}

	targetCharacterId := uint(0)
	switch t := target.Type.(type) {
	case *pb.CharacterTarget_Name:
		targetChar, err := charactersServiceClient.GetCharacter(ctx, target)
		if err != nil {
			return 0, err
		}
		targetCharacterId = uint(targetChar.Id)

	case *pb.CharacterTarget_Id:
		targetCharacterId = uint(t.Id)

	default:
		log.WithContext(ctx).Errorf("target type unknown: %+v", target)
		return 0, model.ErrHandleRequest
	}

	return targetCharacterId, nil
}

func GetCharacterNameFromTarget(
	ctx context.Context,
	charactersServiceClient pb.CharacterServiceClient,
	target *pb.CharacterTarget,
) (string, error) {
	targetCharacterName := ""
	switch t := target.Type.(type) {
	case *pb.CharacterTarget_Name:
		targetCharacterName = t.Name

	case *pb.CharacterTarget_Id:
		targetChar, err := charactersServiceClient.GetCharacter(ctx, target)
		if err != nil {
			return "", err
		}
		targetCharacterName = targetChar.Name

	default:
		log.WithContext(ctx).Errorf("target type unknown: %s", reflect.TypeOf(target.Type).Name())
		return "", model.ErrHandleRequest
	}

	return targetCharacterName, nil
}