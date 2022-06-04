package pixel_rendermaster

import (
	"errors"
	"fmt"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"go.uber.org/zap"
	"golang.org/x/image/colornames"
)

const (
	PATH_BOARD  = "internal/assets/tictactoe_background.png"
	PATH_CROSS  = "internal/assets/cross.png"
	PATH_CIRCLE = "internal/assets/circle.png"
	PATH_BUTTON = "internal/assets/button.png"
)

type PixelRM struct {
	cfg    config.Config
	game   *game
	ds     *PixelRMDS
	logger *zap.Logger

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
type obj struct {
	label string
	pic   pixel.Picture
}

type PixelRMDS struct {
	background obj
	circle     obj
	cross      obj
	button     obj
}

func NewRenderMaster(cfg config.Config, logger *zap.Logger) *PixelRM {
	return &PixelRM{
		cfg:    cfg,
		ds:     &PixelRMDS{},
		logger: logger,
		game:   NewGame(),
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
	background, err := loadPicture(PATH_BOARD)
	if err != nil {
		return err
	}

	cross, err := loadPicture(PATH_CROSS)
	if err != nil {
		return err
	}

	circle, err := loadPicture(PATH_CIRCLE)
	if err != nil {
		return err
	}

	button, err := loadPicture(PATH_BUTTON)
	if err != nil {
		return err
	}

	p.ds.background.pic = background
	p.ds.background.label = PATH_BOARD
	p.ds.cross.pic = cross
	p.ds.cross.label = PATH_CROSS
	p.ds.circle.pic = circle
	p.ds.circle.label = PATH_CIRCLE
	p.ds.button.pic = button
	p.ds.button.label = PATH_BUTTON

	return nil
}

func (p *PixelRM) mainLoop() error {
	// chan frameData  // how tf we get a frame here while not coupling game/pixel

	for !p.win.Closed() {
		err := p.update()
		if err != nil {
			return fmt.Errorf("frame update failed; %v", err)
		}

		err = p.draw()
		if err != nil {
			return fmt.Errorf("frame draw failed; %v", err)
		}

		p.win.Update()
		if err != nil {
			return fmt.Errorf("window update failed; %v", err)
		}
	}
	return nil
}

func (p *PixelRM) draw() error {

	// paint screen white
	p.win.Clear(colornames.White)

	// Draw board in center
	err := p.drawSpriteRect(&p.ds.background, pixel.IM.Moved(p.win.Bounds().Center()))
	if err != nil {
		return err
	}

	// for each board position
	// 	draw symbol inside
	for i := 0; i < 9; {
		for j := i + 3; i < j; i++ {
			var img *obj

			switch p.game.board[int(i)] {
			case '+':
				img = &p.ds.cross
			case 'o':
				img = &p.ds.circle
			default:
				continue
			}

			var x float64 = float64((i % 3) * 100)
			var y float64 = float64((i / 3) * 100)

			err := p.drawSpriteRect(img, pixel.IM.Moved(pixel.Vec{X: x, Y: y}))
			if err != nil {
				return err
			}
		}

	}

	// draw bottom button
	if !p.game.running {
		centerBottom := pixel.Vec{
			X: p.win.Bounds().Center().X,
			Y: p.win.Bounds().Max.Y - 10,
		}
		err := p.drawSpriteRect(&p.ds.button, pixel.IM.Moved(centerBottom))
		if err != nil {
			return err
		}

	}

	return nil
}

func (p *PixelRM) update() error {
	// get input

	// apply input over game state

	// check winning conditions

	// if end, prompt restart and display button

	return nil
}

type game struct {
	running bool

	board string
}

func NewGame() *game {
	return &game{
		running: false,
		board:   "---------",
	}
}

type point struct{ x, y int }

type rect struct{ p1, p2 point }

func (p *PixelRM) drawSpriteRect(cobj *obj, mtrx pixel.Matrix) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sprite '%s' draw failed; %w", cobj.label, r.(error))
			p.logger.Sugar().Errorw("merda", err)
		}
	}()
	sprite := pixel.NewSprite(cobj.pic, cobj.pic.Bounds())
	sprite.Draw(p.win, mtrx)
	return err
}
