package problem

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		Name   string
		Given  int
		Expect bool
	}{
		{
			"Test with 121",
			121,
			true,
		},
		{
			"Test with 1000021",
			1000021,
			false,
		},
		{
			"Test with 1000021",
			11,
			true,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := isPalindrome(test.Given)

			if got != test.Expect {
				t.Errorf("got %v want %v", got, test.Expect)
			}
		})
	}
}
