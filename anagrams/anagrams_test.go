package main

import (
	"reflect"
	"testing"
)

/*
Tests: 2 or more word sets(arrays)
*/

func TestAnagrams(t *testing.T) {
	cases := []struct {
		name   string
		set1   []string
		set2   []string
		assert []int
	}{
		{
			name:   "Test 1",
			set1:   []string{"elbow", "taste", "dust"},
			set2:   []string{"below", "state", "stut"},
			assert: []int{1, 1, 0},
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := Anagrams(test.set1, test.set2)
			t.Logf("\nSet1: %v\nSet2: %v\nOutput: %v\nAssertion: %v", test.set1, test.set2, got, test.assert)
			if !reflect.DeepEqual(got, test.assert) {
				t.Errorf("\nAnagrams Found: %v\nAsserted: %v", got, test.assert)
			}
		})
	}
}
