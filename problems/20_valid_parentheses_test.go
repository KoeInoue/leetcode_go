package problem

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		Name   string
		Given  string
		Expect bool
	}{
		{
			"Test with '()'",
			"()",
			true,
		},
		{
			"Test with ()[]{}",
			"()[]{}",
			true,
		},
		{
			"Test with (]",
			"(]",
			false,
		},
		{
			"Test with ([{}])",
			"([{}])",
			true,
		},
		{
			"Test with ]",
			"]",
			false,
		},
		{
			"Test with (])",
			"(])",
			false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := isValid(test.Given)

			if got != test.Expect {
				t.Errorf("got %v want %v", got, test.Expect)
			}
		})
	}
}
