package problem

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	got := isPalindrome(121)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestIs(t *testing.T) {
	got := isPalindrome(1000021)
	want := false

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIs11(t *testing.T) {
	got := isPalindrome(11)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
