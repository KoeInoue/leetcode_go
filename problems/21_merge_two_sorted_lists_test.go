package problem

import (
	"leetcode/common"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		Name   string
		Given1 *common.ListNode
		Given2 *common.ListNode
		Expect *common.ListNode
	}{
		{
			"Test with lists",
			&common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 4}}},
			&common.ListNode{Val: 1, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4}}},
			&common.ListNode{Val: 1, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 4}}}}}},
		},
		{
			"Test with empty lists",
			&common.ListNode{},
			&common.ListNode{},
			&common.ListNode{},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := mergeTwoLists(test.Given1, test.Given2)

			for {
				if got.Val != test.Expect.Val {
					t.Errorf("got %+v want %+v", got.Val, test.Expect.Val)
				}

				// As soon as either of the  runs out we're done with this loop
				if got.Next == nil {
					break
				} else if test.Expect.Next == nil {
					test.Given1.Next = test.Given2.Next
					break
				}

				got = got.Next
				test.Expect = test.Expect.Next
			}
		})
	}

}
