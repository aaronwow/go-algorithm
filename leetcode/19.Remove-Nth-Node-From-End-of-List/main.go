package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	fast, slow := dummyHead, dummyHead
	for n != 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
}
