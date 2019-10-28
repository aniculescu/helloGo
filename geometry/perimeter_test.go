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
	var width float64 = 10
	var height float64 = 10
	var radius float64 = 5
	var rectangle = Rectangle{width, height}
	var circle = Circle{radius}
	t.Run("Test Area for Rectangle", func(t *testing.T) {
		want := rectangle.Area()
		var got float64 = 100
		if got != want {
			t.Errorf("For width = %.2f and height = %.2f, area should be %.2f, but got %.2f", width, height, want, got)
		}
	})
	t.Run("Test Area for Circle", func(t *testing.T) {
		want := circle.Area()
		var got float64 = math.Pi * (radius * radius)
		if got != want {
			t.Errorf("For width = %.2f and height = %.2f, area should be %.2f, but got %.2f", width, height, want, got)
		}
	})

}
