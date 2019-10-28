package integers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d, received %d", expected, sum)
	}
}

func ExampleAdder() {
	sum := Adder(2, 2)
	fmt.Println(sum)
	//Output: 4
}

func TestSum(t *testing.T) {
	numSlice := []int{1, 2, 4}
	total := Sum(numSlice)
	expected := 7
	if total != expected {
		t.Errorf("Expected %d, received %d in %v", expected, total, numSlice)
	}
}
func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}
	total := SumAll(slice1, slice2)
	expected := []int{3, 9}
	if !reflect.DeepEqual(total, expected) {
		t.Errorf("Expected %v, received %v + %v", expected, slice1, slice2)
	}
}

func TestSumAllTrails(t *testing.T) {
	assertSum := func(t *testing.T, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected %v, received %v", want, got)
		}
	}
	t.Run("Adding slice trails", func(t *testing.T) {
		slice1 := []int{2, 3}
		slice2 := []int{0, 1, 1}
		trailTotal := SumAllTrails(slice1, slice2)
		expected := []int{3, 2}
		assertSum(t, expected, trailTotal)
	})

	t.Run("Adding slice trail with empty slice", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{0, 1, 1}
		trailTotal := SumAllTrails(slice1, slice2)
		expected := []int{0, 2}
		assertSum(t, expected, trailTotal)
	})
}
