package algorithm

import (
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestLinkedListCycleDetect(t *testing.T) {
	// create a linked list with cycle
	head := &utils.ListNode{Val: 1}
	head.Next = &utils.ListNode{Val: 2}
	head.Next.Next = &utils.ListNode{Val: 3}
	head.Next.Next.Next = &utils.ListNode{Val: 4}
	head.Next.Next.Next.Next = &utils.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &utils.ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next

	// test cycle detect
	if GetLoopNode(head) != head.Next.Next {
		t.Errorf("cycle detect failed")
	}

	// create a linked list without cycle
	head = utils.GenerateRandomLinkedList(7)

	// test cycle detect
	if GetLoopNode(head) != nil {
		t.Errorf("cycle detect failed")
	}
}
