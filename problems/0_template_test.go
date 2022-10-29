package problem

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	cases := []struct {
		Name   string
		Given  int
		Expect int
	}{
		{
			"Test with 1",
			1,
			2,
		},
		{
			"Test with 2",
			2,
			4,
		},
		{
			"Test with 100",
			100,
			200,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := template(test.Given)

			if got != test.Expect {
				t.Errorf("got %v want %v", got, test.Expect)
			}
		})
	}
}
