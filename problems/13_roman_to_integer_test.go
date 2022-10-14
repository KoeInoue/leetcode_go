package problem

import "testing"

func TestRomanToInt3(t *testing.T) {
	got := RomanToInt("III")
	want := 3

	if got != want {
		t.Errorf("got %d want %d given, %q", got, want, "III")
	}
}

func TestRomanToInt58(t *testing.T) {
	got := RomanToInt("LVIII")
	want := 58

	if got != want {
		t.Errorf("got %d want %d given, %q", got, want, "LVIII")
	}
}

func TestRomanToInt1994(t *testing.T) {
	got := RomanToInt("MCMXCIV")
	want := 1994

	if got != want {
		t.Errorf("got %d want %d given, %q", got, want, "MCMXCIV")
	}
}
