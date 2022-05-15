package rendermaster

import "github.com/renatobrittoaraujo/rendering/internal/config"

type RenderMaster interface {
}

type pixel struct{}

func NewRenderMaster(config *config.Config) RenderMaster {
	return &pixel{}
}
