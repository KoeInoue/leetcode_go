package problem

import (
	"testing"
)

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		Name   string
		Given  []int
		Expect []int
	}{
		{
			"Test 1",
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			"Test 2",
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := sortedSquares(test.Given)

			if len(got) != len(test.Expect) {
				t.Errorf("got %v want %v", got, test.Expect)
			}
			for i, v := range got {
				if v != test.Expect[i] {
					t.Errorf("got %v want %v", got, test.Expect)
				}
			}
		})
	}
}
