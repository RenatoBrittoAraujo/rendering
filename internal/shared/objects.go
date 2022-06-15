package shared

import "math"

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

func (o *Object) Center() Point {
	return Point{
		X: o.CenterX(),
		Y: o.CenterY(),
	}
}

func (o *Object) W() float64 {
	return math.Abs(o.Pos.Min.X - o.Pos.Max.X)
}

func (o *Object) H() float64 {
	return math.Abs(o.Pos.Min.Y - o.Pos.Max.Y)
}

func (o *Object) MinX() float64 {
	return o.Pos.Min.X
}

func (o *Object) MinY() float64 {
	return o.Pos.Min.Y
}

func (o *Object) MaxX() float64 {
	return o.Pos.Max.X
}

func (o *Object) MaxY() float64 {
	return o.Pos.Max.Y
}

func (o *Object) CenterX() float64 {
	return (o.Pos.Max.X - o.Pos.Min.X) / 2.0
}

func (o *Object) CenterY() float64 {
	return (o.Pos.Max.Y - o.Pos.Min.Y) / 2.0
}
