package pixel_core

import "github.com/faiface/pixel"

type obj struct {
	Label string
	Pic   pixel.Picture
}

type PixelRMDS struct {
	Images map[string]pixel.Picture
}

func NewRDMS() *PixelRMDS {
	return &PixelRMDS{
		Images: make(map[string]pixel.Picture, 0),
	}
}
