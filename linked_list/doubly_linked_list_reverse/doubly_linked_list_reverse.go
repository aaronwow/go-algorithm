package algorithm

import "gandi.icu/go-algorithm/utils"

func DoublyLinkedListReverse(head *utils.DoublyListNode) *utils.DoublyListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *utils.DoublyListNode
	cur := head

	for cur != nil {
		prev = cur.Prev
		cur.Prev = cur.Next
		cur.Next = prev
		cur = cur.Prev
	}
	return prev.Prev
}
