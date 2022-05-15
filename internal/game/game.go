package game

import (
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/util"
	"go.uber.org/zap"
)

type Game interface {
}

type tictactoe struct {
	logger *zap.Logger
}

func NewGame(config config.Config, logger *zap.Logger) (game Game, err error) {
	gameName, err := config.Get("game")
	if err != nil {
		return nil, util.FowardingError("failed to find game in config", err)
	}

	switch gameName {
	case "tictactoe":
		game = &tictactoe{
			logger: logger,
		}
	default:
		util.BaseError(logger, "failed to find game '%s'", gameName)
	}

	return game, err
}

func RunGameLoop(game Game) {

}
