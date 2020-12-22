package main

import (
	"testing"
)

/*
Tests:
Odd number of elements
findMedian([1,6,3,5,8,9,4,10,2]) => 5

Even number of elements
findMedian([1,6,3,5,8,9,4,10,2,7]) => 5.5

*/

func TestFindMedian(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		assert float64
	}{
		{
			name:   "array with odd number of elements",
			arr:    []int{1, 6, 3, 5, 8, 9, 4, 10, 2},
			assert: 5,
		},
		{
			name:   "array with even number of elements",
			arr:    []int{1, 6, 3, 5, 8, 9, 4, 10, 2, 7},
			assert: 5.5,
		},
		{
			name:   "array with 1 element",
			arr:    []int{4},
			assert: 4,
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := FindMedian(test.arr)
			t.Logf("\nArray: %v\nMedian: %v\nAssertion: %v", test.arr, got, test.assert)

			if got != test.assert {
				t.Errorf("\nCalculated Median: %v\nAsserted: %v", got, test.assert)
			}
		})
	}
}
