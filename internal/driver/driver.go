package driver

import (
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/game"
	"github.com/renatobrittoaraujo/rendering/internal/rendermaster"
	"github.com/renatobrittoaraujo/rendering/internal/util"
	"go.uber.org/zap"
)

type Driver interface {
	Run() error
}

type driver struct {
	game         game.Game
	rendermaster rendermaster.RenderMaster
}

func NewDriver(config config.Config, logger *zap.Logger) (Driver, error) {
	game, err := game.NewGame(config, logger)
	if err != nil {
		return nil, util.FowardingError("could not intialize game", err)
	}

	newDriver := &driver{
		game: game,
	}

	return newDriver, nil
}

func (d *driver) Run() error {
	return nil
}
