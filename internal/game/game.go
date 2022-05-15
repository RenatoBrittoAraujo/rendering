package game

import (
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/util"
)

type Game interface {
}

type tictactoe struct {
}

func NewGame(config config.Config) (game Game, err error) {
	gameName, err := config.Get("game")
	if err != nil {
		return nil, util.FowardingError("failed to find game in config", err)
	}

	switch gameName {
	case "tictactoe":
		game = &tictactoe{}
	default:
		util.BaseError("failed to find game '%s'", gameName)
	}

	return game, err
}

func RunGameLoop(game Game) {

}
