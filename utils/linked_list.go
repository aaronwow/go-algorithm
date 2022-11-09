package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

func GenerateRandomLinkedList(length int) *ListNode {
	if length <= 0 {
		return nil
	}

	head := &ListNode{Val: 1}
	cur := head
	for i := 2; i <= length; i++ {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}

	return head
}

func GenerateRandomDoublyLinkedList(length int) *DoublyListNode {
	if length <= 0 {
		return nil
	}

	head := &DoublyListNode{Val: 1}
	cur := head
	for i := 2; i <= length; i++ {
		cur.Next = &DoublyListNode{Val: i}
		cur.Next.Prev = cur
		cur = cur.Next
	}

	return head
}
