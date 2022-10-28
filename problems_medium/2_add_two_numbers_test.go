package problem

import (
	"leetcode/common"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &common.ListNode{2, &common.ListNode{4, &common.ListNode{Val: 3}}}
	l2 := &common.ListNode{5, &common.ListNode{6, &common.ListNode{Val: 4}}}
	got := AddTwoNumbers(l1, l2)
	want := &common.ListNode{7, &common.ListNode{0, &common.ListNode{Val: 8}}}

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
