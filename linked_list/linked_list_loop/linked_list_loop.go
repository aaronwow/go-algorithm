package algorithm

import "gandi.icu/go-algorithm/utils"

// 1. determine whether a linked list has a cycle
func GetLoopNode(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow := head.Next
	fast := head.Next.Next

	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
