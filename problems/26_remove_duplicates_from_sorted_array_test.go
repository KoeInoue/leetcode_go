package problem

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		Name   string
		Given  []int
		Expect []int
	}{
		{
			"Test with array {1,1,2}",
			[]int{1, 1, 2},
			[]int{1, 2},
		},
		{
			"Test with array {1,2,2, 3, 4, 4, 5,5}",
			[]int{1, 2, 2, 3, 4, 4, 5, 5},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range cases {
		t.Run(test.Name+" length", func(t *testing.T) {
			assertInt(t, test.Expect, test.Given)
		})
	}
}

func assertInt(t testing.TB, want, given []int) {
	t.Helper()

	got := removeDuplicates(given)

	if got != len(want) {
		t.Errorf("got %d, want %d", got, len(want))
	}
}
