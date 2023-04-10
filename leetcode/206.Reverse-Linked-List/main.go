package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, temp *ListNode
	cur := head

	for cur != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}
