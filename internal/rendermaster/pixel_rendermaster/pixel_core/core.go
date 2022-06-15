package pixel_core

import (
	"errors"
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"go.uber.org/zap"
)

type PixelCore struct {
	cfg    config.Config
	ds     *PixelRMDS
	logger *zap.Logger

	Win    *pixelgl.Window
	Wincfg pixelgl.WindowConfig

	LoopMethods LoopMethods
}

type LoopMethods interface {
	Update() error
	Draw() error
}

func NewPixelCore(cfg config.Config, logger *zap.Logger, loopMethods LoopMethods) (*PixelCore, error) {
	wincfg := pixelgl.WindowConfig{
		Title:  "Whatever",
		Bounds: pixel.R(0, 0, 2000, 2000),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(wincfg)
	if err != nil {
		return nil, errors.New("config failure: failed to create OS window")
	}

	return &PixelCore{
		cfg:         cfg,
		ds:          NewRDMS(),
		Win:         win,
		Wincfg:      wincfg,
		LoopMethods: loopMethods,
	}, nil
}

func (p *PixelCore) MainLoop() error {
	for !p.Win.Closed() {
		err := p.LoopMethods.Update()
		if err != nil {
			return fmt.Errorf("frame update failed; %v", err)
		}

		err = p.LoopMethods.Draw()
		if err != nil {
			return fmt.Errorf("frame draw failed; %v", err)
		}

		p.Win.Update()
		if err != nil {
			return fmt.Errorf("window update failed; %v", err)
		}
	}
	return nil
}
