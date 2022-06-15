package pixel_core

import (
	"fmt"
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/renatobrittoaraujo/rendering/internal/shared"
)

func loadImage(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (p *PixelCore) LoadImages(images []shared.Image) error {
	var final_err error
	for _, img_data := range images {
		img, err := loadImage(img_data.Path)
		if err != nil {
			final_err = fmt.Errorf("%w; %v", final_err, err)
		}
		p.ds.Images[img_data.Name] = img
	}
	return final_err
}
