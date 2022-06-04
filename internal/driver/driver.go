package driver

import (
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/game"
	"github.com/renatobrittoaraujo/rendering/internal/rendermaster"
	"github.com/renatobrittoaraujo/rendering/internal/util"
	"go.uber.org/zap"
)

// Driver manages the setup of game and renderer
// it provides communication abstractions between both
// Game --frames--> Rendermaster
// Rendermaster --inputs--> Game

type Driver interface {
	Run() error
}

type driver struct {
	inputs chan Inputs
	frames chan Frames

	game         game.Game
	rendermaster rendermaster.Rendermaster
}

// Call rendermaster
// give it main thread
// delegate all non-render tasks to other threads
// keep communications in sync via channels, should be seamless to routines
func NewDriver(config config.Config, logger *zap.Logger) (Driver, error) {
	game, err := game.NewGame(config, logger)
	if err != nil {
		return nil, util.FowardingError("could not intialize game", err)
	}

	rendermaster, err := rendermaster.NewRendermaster(config)
	if err != nil {
		return nil, util.FowardingError("could not intialize rendermaster", err)
	}

	newDriver := &driver{
		game:         game,
		rendermaster: rendermaster,
	}

	return newDriver, nil
}

func (d *driver) Run() error {
	err := d.rendermaster.Run()
	return err
}

type Frames struct {
}

type Inputs struct {
}
