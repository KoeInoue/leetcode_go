package problem

import (
	"testing"
)

func TestRotateArray(t *testing.T) {
	cases := []struct {
		Name      string
		GivenArr  []int
		GivenStep int
		Expect    []int
	}{
		{
			"Test 1",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"Test 2",
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			"Test k=0",
			[]int{-1, -100, 3, 99},
			0,
			[]int{-1, -100, 3, 99},
		},
		{
			"Test length=1",
			[]int{1},
			1,
			[]int{1},
		},
		{
			"Test k>length",
			[]int{1, 2, 3, 4, 5, 6, 7},
			100,
			[]int{6, 7, 1, 2, 3, 4, 5},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			rotate(test.GivenArr, test.GivenStep)
		})
	}
}
