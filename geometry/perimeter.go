package integers

import "math"

//Rectangle Struct
type Rectangle struct {
	Width, Height float64
}

//Circle Struct
type Circle struct {
	Radius float64
}

//Triangle struct
type Triangle struct {
	Base, Height float64
}

//Shape interface
type Shape interface {
	Area() float64
}

//Perimeter Function
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area method of Rectangle type
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Area method of Circle type
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Area method for Triangle type
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
