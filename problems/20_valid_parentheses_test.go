package problem

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	given := "()"
	got := isValid(given)
	want := true

	if got != want {
		t.Errorf("got %v want %v given %v", got, want, given)
	}
}

func TestIsValidParentheses3(t *testing.T) {
	given := "()[]{}"
	got := isValid(given)
	want := true

	if got != want {
		t.Errorf("got %v want %v given %v", got, want, given)
	}
}

func TestIsValidParentheses2(t *testing.T) {
	given := "(]"
	got := isValid(given)
	want := false

	if got != want {
		t.Errorf("got %v want %v given %v", got, want, given)
	}
}

func TestIsValidParenthesesNest(t *testing.T) {
	given := "([{}])"
	got := isValid(given)
	want := true

	if got != want {
		t.Errorf("got %v want %v given %v", got, want, given)
	}
}

func TestIsValidClose(t *testing.T) {
	given := "]"
	got := isValid(given)
	want := false

	if got != want {
		t.Errorf("got %v want %v given %v", got, want, given)
	}
}

func TestIsValidMixed(t *testing.T) {
	given := "(])"
	got := isValid(given)
	want := false

	if got != want {
		t.Errorf("got %v want %v given %v", got, want, given)
	}
}
