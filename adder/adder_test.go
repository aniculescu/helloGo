package integers

import (
	"fmt"
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
