package problem

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		Name      string
		GivenArr  []int
		GivenVal  int
		Expect    []int
		ExpectVal int
	}{
		{
			"Test with array {3, 2, 2, 3}",
			[]int{3, 2, 2, 3},
			3,
			[]int{2, 2, 3, 3},
			2,
		},
		{
			"Test with array {0,1,2,2,3,0,4,2}",
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			[]int{0, 1, 4, 0, 3, 2, 2, 2},
			5,
		},
		{
			"Test with array {1}",
			[]int{1},
			1,
			[]int{1},
			0,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()

			got := removeElement(test.GivenArr, test.GivenVal)

			if got != test.ExpectVal {
				t.Errorf("got %d, want %d", got, len(test.Expect))
			}
		})
	}
}
