package problem

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	got := longestCommonPrefix([]string{"flower", "flow", "flight"})
	want := "fl"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
