package rendermaster

import (
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/rendermaster/pixel_rendermaster"
)

// Is the owner of main thread.
type Rendermaster interface {
	// init
	Run() error

	// input
	// keys pressed
	// mouse pressed
	// game close call

	// output
	// LoadImage(path string)
	// MakeSprite([]Image) (Sprite, error)

	// draw them somewhere
}

func NewRendermaster(cfg config.Config) (Rendermaster, error) {
	return pixel_rendermaster.NewRenderMaster(cfg), nil
}
