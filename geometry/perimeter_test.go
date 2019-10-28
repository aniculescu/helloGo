package integers

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	var width float64 = 20
	var height float64 = 40
	var rectangle = Rectangle{width, height}

	want := Perimeter(rectangle)
	var got float64 = 120
	if got != want {
		t.Errorf("For width = %.2f and height = %.2f, perimeter should be %.2f, but got %.2f", width, height, want, got)
	}
}

func TestArea(t *testing.T) {

	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, want: 100},
		{name: "Circle", shape: Circle{Radius: 5}, want: math.Pi * (5 * 5)},
		{name: "Triangle", shape: Triangle{Base: 15, Height: 4}, want: 0.5 * 5 * 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.shape.Area()
			if got != c.want {
				t.Errorf("Shape: %#v - Calculated Area: %g, Actual Area %g", c.shape, got, c.want)
			}
		})

	}

}
