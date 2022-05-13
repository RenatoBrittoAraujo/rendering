package driver

import (
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/ebiten_helper"
	"github.com/renatobrittoaraujo/rendering/internal/game"
)

type Driver interface {
	Setup(config *config.Config) err
	Run() err
}

type driver struct {
	renderMaster ebiten_helper.EbitenHelper
	game         game.Game
}

func NewDriver() Driver {
	newDriver := &driver{}

	return newDriver
}

func (d *driver) HasFrame() {

}

func (d *driver) GetLastFrame() {

}

func (d *driver) GenerateFrame() {

}

func (d *driver) UpdateGame() {

}

func (d *driver) Start() {

}
