package pixel_util

import (
	"errors"
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/shared"
	"go.uber.org/zap"
)

type PixelUtil struct {
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

func NewPixelUtil(cfg config.Config, logger *zap.Logger, loopMethods LoopMethods) (*PixelUtil, error) {
	wincfg := pixelgl.WindowConfig{
		Title:  "Whatever",
		Bounds: pixel.R(0, 0, 2000, 2000),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(wincfg)
	if err != nil {
		return nil, errors.New("config failure: failed to create OS window")
	}

	return &PixelUtil{
		cfg:         cfg,
		ds:          NewRDMS(),
		Win:         win,
		Wincfg:      wincfg,
		LoopMethods: loopMethods,
	}, nil
}

func (p *PixelUtil) LoadImages(images []shared.Image) error {
	var final_err error
	for _, img_data := range images {
		img, err := loadPicture(img_data.Path)
		if err != nil {
			final_err = fmt.Errorf("%w; %v", final_err, err)
		}
		p.ds.Images[img_data.Name] = img
	}
	return final_err
}

func (p *PixelUtil) MainLoop() error {
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

func (p *PixelUtil) DrawSpriteRect(obj *shared.Object) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sprite '%s' draw failed; %w", obj.Image.Name, r.(error))
			p.logger.Sugar().Errorw("merda", err)
		}
	}()

	mtrx := pixel.IM.
		Moved(pixel.Vec{obj.Pos.Min.X, obj.Pos.Min.Y}).
		Scaled(pixel.Vec{0, 100}, 0.5)

	image := p.ds.Images[obj.Image.Name]
	sprite := pixel.NewSprite(image, image.Bounds())
	sprite.Draw(p.Win, mtrx)
	return err
}
