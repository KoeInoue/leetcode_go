package problem

import (
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		Name        string
		GivenNums   []int
		GivenTarget int
		Expect      int
	}{
		{
			"Test expect index 2",
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			"Test expect index 1",
			[]int{1, 3, 5, 6},
			2,
			-1,
		},
		{
			"Test expect index 4",
			[]int{1, 3, 5, 6},
			7,
			-1,
		},
		{
			"Test expect index 0",
			[]int{1},
			1,
			0,
		},
		{
			"Test expect index 1 with array of ",
			[]int{0},
			1,
			-1,
		},
		{
			"Test expect index 0-0",
			[]int{1, 3, 5, 6},
			0,
			-1,
		},
		{
			"Test expect index 1-3",
			[]int{1, 3},
			3,
			1,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := search(test.GivenNums, test.GivenTarget)

			if got != test.Expect {
				t.Errorf("got %v want %v", got, test.Expect)
			}
		})
	}
}
