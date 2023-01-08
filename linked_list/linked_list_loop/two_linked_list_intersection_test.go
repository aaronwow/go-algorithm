package algorithm

import (
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestBothNoLoop(t *testing.T) {
	// create head1 and head2 without loop and also without intersection
	head1 := utils.GenerateRandomLinkedList(10)
	head2 := utils.GenerateRandomLinkedList(10)
	if TwoLinkedListIntersection(head1, head2) != nil {
		t.Errorf("both no loop failed")
	}
	// create head1 and head2 without loop and but have intersection
	head1 = &utils.ListNode{Val: 1}
	head1.Next = &utils.ListNode{Val: 2}
	head1.Next.Next = &utils.ListNode{Val: 3}
	head1.Next.Next.Next = &utils.ListNode{Val: 4}
	head1.Next.Next.Next.Next = &utils.ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next = &utils.ListNode{Val: 6}
	head2 = &utils.ListNode{Val: 7}
	head2.Next = &utils.ListNode{Val: 8}
	head2.Next.Next = &utils.ListNode{Val: 9}
	head2.Next.Next.Next = head1.Next.Next.Next
	if TwoLinkedListIntersection(head1, head2) != head1.Next.Next.Next {
		t.Errorf("both no loop failed")
	}
}

func TestBothLoop(t *testing.T) {
	// create head1 and head2 with loop and also without intersection
	head1 := &utils.ListNode{Val: 1}
	head1.Next = &utils.ListNode{Val: 2}
	head1.Next.Next = &utils.ListNode{Val: 3}
	head1.Next.Next.Next = &utils.ListNode{Val: 4}
	head1.Next.Next.Next.Next = &utils.ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next = &utils.ListNode{Val: 6}
	head1.Next.Next.Next.Next.Next.Next = head1.Next.Next
	head2 := &utils.ListNode{Val: 7}
	head2.Next = &utils.ListNode{Val: 8}
	head2.Next.Next = &utils.ListNode{Val: 9}
	head2.Next.Next.Next = &utils.ListNode{Val: 10}
	head2.Next.Next.Next.Next = &utils.ListNode{Val: 11}
	head2.Next.Next.Next.Next.Next = &utils.ListNode{Val: 12}
	head2.Next.Next.Next.Next.Next.Next = head2.Next.Next
	if TwoLinkedListIntersection(head1, head2) != nil {
		t.Errorf("both loop failed")
	}
	// create head1 and head2 with loop and also with intersection but different loop enter point
	head1 = &utils.ListNode{Val: 1}
	head1.Next = &utils.ListNode{Val: 2}
	head1.Next.Next = &utils.ListNode{Val: 3}
	head1.Next.Next.Next = &utils.ListNode{Val: 4}
	head1.Next.Next.Next.Next = &utils.ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next = &utils.ListNode{Val: 6}
	head1.Next.Next.Next.Next.Next.Next = head1.Next.Next
	head2 = &utils.ListNode{Val: 7}
	head2.Next = &utils.ListNode{Val: 8}
	head2.Next.Next = &utils.ListNode{Val: 9}
	head2.Next.Next.Next = &utils.ListNode{Val: 10}
	head2.Next.Next.Next.Next = &utils.ListNode{Val: 11}
	head2.Next.Next.Next.Next.Next = &utils.ListNode{Val: 12}
	head2.Next.Next.Next.Next.Next.Next = head1.Next.Next.Next
	if TwoLinkedListIntersection(head1, head2) != head1.Next.Next.Next || TwoLinkedListIntersection(head2, head1) != head1.Next.Next {
		t.Errorf("both loop failed")
	}

	// create head1 and head2 with loop and also with intersection and same loop enter point
	head1 = &utils.ListNode{Val: 1}
	head1.Next = &utils.ListNode{Val: 2}
	head1.Next.Next = &utils.ListNode{Val: 3}
	head1.Next.Next.Next = &utils.ListNode{Val: 4}
	head1.Next.Next.Next.Next = &utils.ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next = &utils.ListNode{Val: 6}
	head1.Next.Next.Next.Next.Next.Next = head1.Next.Next.Next.Next
	head2 = &utils.ListNode{Val: 7}
	head2.Next = &utils.ListNode{Val: 8}
	head2.Next.Next = &utils.ListNode{Val: 9}
	head2.Next.Next.Next = &utils.ListNode{Val: 10}
	head2.Next.Next.Next.Next = &utils.ListNode{Val: 11}
	head2.Next.Next.Next.Next.Next = &utils.ListNode{Val: 12}
	head2.Next.Next.Next.Next.Next.Next = head1.Next.Next
	if TwoLinkedListIntersection(head1, head2) != head1.Next.Next {
		t.Errorf("both loop failed")
	}
}
