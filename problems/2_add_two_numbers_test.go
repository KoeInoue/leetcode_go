package problem

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{Val: 3}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{Val: 4}}}
	got := AddTwoNumbers(l1, l2)
	want := &ListNode{7, &ListNode{0, &ListNode{Val: 8}}}

	for {
		if got.Val != want.Val {
			t.Errorf("got %+v want %+v", got.Val, want.Val)
		}

		// As soon as either of the  runs out we're done with this loop
		if got.Next == nil {
			break
		} else if want.Next == nil {
			l1.Next = l2.Next
			break
		}

		// Advance!
		got = got.Next
		want = want.Next
	}
}
