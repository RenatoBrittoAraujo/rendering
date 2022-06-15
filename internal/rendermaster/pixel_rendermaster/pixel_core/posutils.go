package pixel_core

func (p *PixelCore) GetCenterX() float64 {
	return p.Win.Bounds().Center().X
}

func (p *PixelCore) GetCenterY() float64 {
	return p.Win.Bounds().Center().Y
}

func (p *PixelCore) GetMaxX() float64 {
	return p.Win.Bounds().Max.X
}

func (p *PixelCore) GetMaxY() float64 {
	return p.Win.Bounds().Max.Y
}

func (p *PixelCore) GetMinX() float64 {
	return p.Win.Bounds().Min.X
}

func (p *PixelCore) GetMinY() float64 {
	return p.Win.Bounds().Min.Y
}
