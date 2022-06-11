package shared

type Sprite struct {
	Sequence []Image
}

type Image struct {
	Path string
	Name string
}

type Point struct {
	X, Y, Z float64
}

type Rect struct {
	Min Point
	Max Point
}

type Object struct {
	Image Image
	Label string

	Pos Rect
	Rot float64 // radians
}
