package problem

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Test with array {1,1,2}", func(t *testing.T) {
		given := []int{1, 1, 2}
		want := []int{1, 2}

		assertInt(t, want, given)
	})
}

func assertInt(t testing.TB, want, given []int) {
	t.Helper()

	got := removeDuplicates(given)

	if got != len(want) {
		t.Errorf("got %d, want %d", got, len(want))
	}
}
