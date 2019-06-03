package mathutil

import "testing"

func TestAdding(t *testing.T) {
	cases := []struct {
		num1, num2, want int
	}{
		{2, 3, 5},
		{1, 2, 3},
		{3, 5, 8},
	}
	for _, c := range cases {
		got := Adding(c.num1, c.num2)
		if got != c.want {
			t.Errorf("Adding %d + %d, got %d want %d", c.num1, c.num2, got, c.want)
		}
	}
}
