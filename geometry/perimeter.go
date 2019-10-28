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

//Area function
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
