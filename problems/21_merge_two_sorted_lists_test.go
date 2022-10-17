package problem

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{Val: 4}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{Val: 4}}}
	got := mergeTwoLists(list1, list2)
	want := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{Val: 4}}}}}}

	for {
		if got.Val != want.Val {
			t.Errorf("got %+v want %+v", got.Val, want.Val)
		}

		// As soon as either of the  runs out we're done with this loop
		if got.Next == nil {
			break
		} else if want.Next == nil {
			list1.Next = list2.Next
			break
		}

		// Advance!
		got = got.Next
		want = want.Next
	}
}

func TestMergeTwoListsEmpty(t *testing.T) {
	list1 := &ListNode{}
	list2 := &ListNode{}
	got := mergeTwoLists(list1, list2)
	want := &ListNode{}

	for {
		// As soon as either of the  runs out we're done with this loop
		if got.Next == nil {
			break
		} else if want.Next == nil {
			list1.Next = list2.Next
			break
		}

		if got.Val != want.Val {
			t.Errorf("got %+v want %+v", got.Val, want.Val)
		}

		// Advance!
		got = got.Next
		want = want.Next
	}
}
