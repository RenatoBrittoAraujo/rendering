package commons

type Sprite struct {
	sequence []Image
}

type Image struct {
	path  string
	pos   Point
	scale float64
	rot   float64 // radians
}

type Point struct {
	x, y, z float64
}

type Rect struct {
	p1 Point
	p2 Point
}
