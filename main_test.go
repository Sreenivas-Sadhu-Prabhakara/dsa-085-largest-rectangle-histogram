package main

import "testing"

func TestLargestRectangle(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{5}, 5},
		{[]int{}, 0},
		{[]int{6, 2, 5, 4, 5, 1, 6}, 12},
		{[]int{0, 0}, 0},
		{[]int{5, 4, 3, 2, 1}, 9},
	}
	for _, c := range cases {
		if got := LargestRectangle(c.in); got != c.want {
			t.Errorf("LargestRectangle(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
