package problem

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		Name   string
		Given  []string
		Expect string
	}{
		{
			"Test with FL",
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			"Test with None",
			[]string{"dog", "racecar", "car"},
			"",
		},
		{
			"Test with C",
			[]string{"cir", "car"},
			"c",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := longestCommonPrefix(test.Given)

			if got != test.Expect {
				t.Errorf("got %v want %v", got, test.Expect)
			}
		})
	}
}
