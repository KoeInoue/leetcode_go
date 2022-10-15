package problem

import (
	"testing"
)

func TestLongestCommonPrefixFl(t *testing.T) {
	got := longestCommonPrefix([]string{"flower", "flow", "flight"})
	want := "fl"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLongestCommonPrefixNone(t *testing.T) {
	got := longestCommonPrefix([]string{"dog", "racecar", "car"})
	want := ""

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLongestCommonPrefixC(t *testing.T) {
	got := longestCommonPrefix([]string{"cir", "car"})
	want := "c"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
