package pixel_util

func (p *PixelUtil) GetCenterX() float64 {
	return p.Win.Bounds().Center().X
}

func (p *PixelUtil) GetCenterY() float64 {
	return p.Win.Bounds().Center().Y
}

func (p *PixelUtil) GetMaxX() float64 {
	return p.Win.Bounds().Max.X
}

func (p *PixelUtil) GetMaxY() float64 {
	return p.Win.Bounds().Max.Y
}

func (p *PixelUtil) GetMinX() float64 {
	return p.Win.Bounds().Min.X
}

func (p *PixelUtil) GetMinY() float64 {
	return p.Win.Bounds().Min.Y
}
