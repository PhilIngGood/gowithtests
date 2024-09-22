package geometry

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() (perim float64) {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() (area float64) {
	return r.Width * r.Height
}

func (c Circle) Area() (area float64) {
	return math.Pow(c.Radius, 2) * math.Pi
}

func (t Triangle) Area() (area float64) {
	return t.Width * t.Height / 2
}
