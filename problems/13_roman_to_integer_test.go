package problem

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		Name   string
		Given  string
		Expect int
	}{
		{
			"Test with III",
			"III",
			3,
		},
		{
			"Test with LVIII",
			"LVIII",
			58,
		},
		{
			"Test with MCMXCIV",
			"MCMXCIV",
			1994,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := RomanToInt(test.Given)

			if got != test.Expect {
				t.Errorf("got %v want %v", got, test.Expect)
			}
		})
	}
}
