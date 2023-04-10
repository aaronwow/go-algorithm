package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	var lenA, lenB int
	for p1 != nil {
		p1 = p1.Next
		lenA++
	}
	for p2 != nil {
		p2 = p2.Next
		lenB++
	}

	gap := abs(lenA - lenB)

	if lenB > lenA {
		p1, p2 = headB, headA
	} else {
		p1, p2 = headA, headB
	}

	for gap != 0 {
		p1 = p1.Next
		gap--
	}

	for p1 != nil && p2 != nil && p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
