package problem

import (
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		Name   string
		Given  int
		Expect int
	}{
		{
			"Test with 5",
			5,
			4,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := firstBadVersion(test.Given)

			if got != test.Expect {
				t.Errorf("got %v want %v", got, test.Expect)
			}
		})
	}
}
