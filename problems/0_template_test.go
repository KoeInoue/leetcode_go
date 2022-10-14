package problem

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	got := template()
	want := 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
