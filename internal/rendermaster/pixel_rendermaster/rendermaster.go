package pixel_rendermaster

import (
	"errors"
	"fmt"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"golang.org/x/image/colornames"
)

type PixelRM struct {
	cfg config.Config
	ds  *PixelRMDS

	win    *pixelgl.Window
	wincfg pixelgl.WindowConfig
}

// Data structure for multi user inputs of data loading
// This data should be set by the game, this place only store
// game asset categories
// TODO: interface for the game and this shit here
// TIP: maybe disconnect completely from the game? make it a function you call once and
// shit just happens in the layer below? syncing problems with this. Maybe keep only assets
// here and everything that happens above they will feed everything. We just do as they say.
type PixelRMDS struct {
	pic pixel.Picture
}

func NewRenderMaster(cfg config.Config) *PixelRM {
	return &PixelRM{
		cfg: cfg,
		ds:  &PixelRMDS{},
	}
}

func (p *PixelRM) Run() error {
	var err error
	pixelgl.Run(func() {
		err = p.config()
		if err != nil {
			return
		}

		err = p.load()
		if err != nil {
			return
		}

		err = p.mainLoop()
		if err != nil {
			return
		}
	})

	if err != nil {
		return fmt.Errorf("pixel run failure; %v", err)
	}
	return nil
}

func (p *PixelRM) config() error {
	wincfg := pixelgl.WindowConfig{
		Title:  "Whatever",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(wincfg)
	if err != nil {
		return errors.New("config failure: failed to create OS window")
	}

	p.win = win
	p.wincfg = wincfg

	return nil
}

func (p *PixelRM) load() error {
	pic, err := loadPicture("internal/assets/surfing-js.png")
	if err != nil {
		return err
	}
	p.ds.pic = pic

	return nil
}

func (p *PixelRM) mainLoop() error {
	// chan frameData  // how tf we get a frame here while not coupling game/pixel

	for !p.win.Closed() {
		err := p.draw()
		if err != nil {
			return errors.New("frame draw failed")
		}

		p.win.Update()
		if err != nil {
			return errors.New("window update failed")
		}
	}
	return nil
}

func (p *PixelRM) draw() error {
	sprite := pixel.NewSprite(p.ds.pic, p.ds.pic.Bounds())
	p.win.Clear(colornames.Greenyellow)
	sprite.Draw(p.win, pixel.IM.Moved(p.win.Bounds().Center()))

	return nil
}
