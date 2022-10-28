package problem

import "leetcode/common"

func AddTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	// We'll use l1 to write the results in, so we need to store the head
	carry, head := 0, l1
	for {
		// Add the two nodes' values and the carry, write to l1
		l1.Val += l2.Val + carry

		// Check for carry and update l1's value
		carry = int(l1.Val / 10)
		l1.Val = l1.Val % 10 // or l1.Next.Val - (carry * 10)

		// As soon as either of the  runs out we're done with this loop
		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}

		// Advance!
		l1 = l1.Next
		l2 = l2.Next
	}

	// We'll now apply the carry to l1, adding nodes as needed
	for carry != 0 {
		// If there is a carry and l1.Next is null, add another node to the list
		if l1.Next == nil {
			l1.Next = &common.ListNode{0, nil}
		}

		// Same as the first loop, just using one list and the carry
		l1.Next.Val += carry

		carry = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10

		// Advance!
		l1 = l1.Next
	}

	return head
}
