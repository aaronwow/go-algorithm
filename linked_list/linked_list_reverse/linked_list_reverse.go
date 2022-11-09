package algorithm

import "gandi.icu/go-algorithm/utils"

// Two Pointer
func LinkedListReverse(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *utils.ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
