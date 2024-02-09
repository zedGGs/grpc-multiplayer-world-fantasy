package players

import (
	"context"

	"github.com/prometheus/alertmanager/config"
)

var (
	ServiceName = "players"
)

type CharactersServerContext struct {
	GlobalConfig     *config.GlobalConfig
	CharacterService service.CharacterService
}

func NewServerContext(ctx context.Context, conf *config.GlobalConfig) *CharactersServerContext {
	server := &CharactersServerContext{
		GlobalConfig: conf,
	}

	return server
}
