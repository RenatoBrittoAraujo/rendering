package pixel_rendermaster

import (
	"fmt"
	_ "image/png"
	"math"

	"github.com/faiface/pixel/pixelgl"
	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/rendermaster/pixel_rendermaster/pixel_core"
	"github.com/renatobrittoaraujo/rendering/internal/shared"
	"go.uber.org/zap"
	"golang.org/x/image/colornames"
)

const (
	PATH_BOARD  = "internal/assets/black_square.png"
	PATH_CROSS  = "internal/assets/cross.png"
	PATH_CIRCLE = "internal/assets/circle.png"
	PATH_BUTTON = "internal/assets/button.png"
)

type PixelRM struct {
	cfg    config.Config
	game   *game
	logger *zap.Logger

	pixelCore *pixel_core.PixelCore

	board  shared.Object
	button shared.Object

	cross  shared.Image
	circle shared.Image
}

// Data structure for multi user inputs of data loading
// This data should be set by the game, this place only store
// game asset categories
// TODO: interface for the game and this shit here
// TIP: maybe disconnect completely from the game? make it a function you call once and
// shit just happens in the layer below? syncing problems with this. Maybe keep only assets
// here and everything that happens above they will feed everything. We just do as they say.

func NewRenderMaster(cfg config.Config, logger *zap.Logger) *PixelRM {
	return &PixelRM{
		cfg:    cfg,
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

		err = p.pixelCore.MainLoop()
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
	pixelCore, err := pixel_core.NewPixelCore(p.cfg, p.logger, p)
	if err != nil {
		return err
	}
	p.pixelCore = pixelCore

	return nil
}

func (p *PixelRM) load() error {
	p.board = shared.Object{
		Image: shared.Image{
			Path: PATH_BOARD,
			Name: "board",
		},
		Pos: shared.Rect{
			Min: shared.Point{
				X: p.pixelCore.GetMinX(),
				Y: p.pixelCore.GetMinY(),
			},
			Max: shared.Point{
				X: p.pixelCore.GetMaxX(),
				Y: p.pixelCore.GetMaxY(),
			},
		},
	}

	p.button = shared.Object{
		Image: shared.Image{
			Path: PATH_BUTTON,
			Name: "button",
		},
		Pos: shared.Rect{
			Min: shared.Point{
				X: p.pixelCore.GetCenterX() - 300,
				Y: p.pixelCore.GetCenterY() + 350,
			},
			Max: shared.Point{
				X: p.pixelCore.GetCenterX() + 300,
				Y: p.pixelCore.GetCenterY() + 400,
			},
		},
	}

	p.cross = shared.Image{
		Path: PATH_CROSS,
		Name: "cross",
	}
	p.circle = shared.Image{
		Path: PATH_CIRCLE,
		Name: "circle",
	}

	images := []shared.Image{
		p.cross,
		p.circle,
		p.board.Image,
		p.button.Image,
	}

	err := p.pixelCore.LoadImages(images)
	if err != nil {
		return err
	}

	return nil
}

func (p *PixelRM) Draw() error {

	// paint screen white
	p.pixelCore.Win.Clear(colornames.White)

	// Draw board in center
	err := p.pixelCore.DrawSpriteRect(&p.board)
	if err != nil {
		return err
	}

	// for each board position
	// 	draw symbol inside
	// for i := 0; i < 9; {
	// 	for j := i + 3; i < j; i++ {
	// 		var img *obj

	// 		switch p.game.board[int(i)] {
	// 		case '+':
	// 			img = &p.ds.cross
	// 		case 'o':
	// 			img = &p.ds.circle
	// 		default:
	// 			continue
	// 		}

	// 		var x float64 = float64((i % 3) * 100)
	// 		var y float64 = float64((i / 3) * 100)

	// 		err := p.pixelCore.DrawSpriteRect(img, pixel.IM.Moved(pixel.Vec{X: x, Y: y}))
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	// }

	// draw bottom button
	// if !p.game.running {

	// 	// center :=
	// 	err := p.drawSpriteRect(&p.ds.button, &shared.Rect{
	// 		Min: {
	// 			X: ,
	// 		},
	// 		Max: {},
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}

	// }

	return nil
}

var i float64

func (p *PixelRM) Update() error {
	// get input

	// apply input over game state

	// check winning conditions

	// if end, prompt restart and display button

	/// ========== tests

	i++
	p.board.Pos.Min.X = p.pixelCore.GetMinX() + math.Abs(math.Sin(i/50))*100
	p.board.Pos.Min.Y = p.pixelCore.GetMinY() + math.Abs(math.Sin(i/50))*100

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

// scale images to the size user wants
// process mouse click inputs (preferebly easy)
