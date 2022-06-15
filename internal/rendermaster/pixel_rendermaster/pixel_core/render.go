package pixel_core

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/renatobrittoaraujo/rendering/internal/shared"
)

// Draws a picture as a sprite, starting out from point Min to point Max
// with rotation applied after. Image will be made to fit the size passed.
func (p *PixelCore) DrawSpriteRect(obj *shared.Object) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sprite '%s' draw failed; %w", obj.Image.Name, r.(error))
			p.logger.Sugar().Errorw("merda", err)
		}
	}()

	mtrx := pixel.IM
	img := p.ds.Images[obj.Image.Name]
	imgBounds := img.Bounds()
	imgFrame := imgBounds.Resized(imgBounds.Center(), pixel.V(obj.W(), obj.H()))

	xScale := obj.W() / imgBounds.W()
	yScale := obj.H() / imgBounds.H()
	mtrx = mtrx.ScaledXY(imgBounds.Min, pixel.V(xScale, yScale))

	// Correct for center justification of positions
	x := obj.Center().X
	y := obj.Center().Y
	mtrx = mtrx.Moved(pixel.V(x, y))

	sprite := pixel.NewSprite(img, imgFrame)
	sprite.Draw(p.Win, mtrx)

	return err
}
