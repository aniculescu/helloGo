package algos

import (
	"reflect"
	"testing"
)

func TestArrayIncludes(t *testing.T) {
	cases := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{name: "Single num difference", arr1: []int{1, 2}, arr2: []int{2, 3, 1}, want: []int{3}},
		{name: "Multiples num difference", arr1: []int{4, 1, 7}, arr2: []int{4, 2, 9, 7, 1}, want: []int{2, 9}},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := ArrayIncludes(test.arr1, test.arr2)
			t.Logf("Array1: %v", test.arr1)
			t.Logf("Array2: %v", test.arr2)
			t.Logf("Difference: %v", got)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
