package algorithm

import (
	"gandi.icu/go-algorithm/utils"
)

// get the intersection of two linked list, both list could have loop or not, or maybe they don't have intersection
func TwoLinkedListIntersection(h1, h2 *utils.ListNode) *utils.ListNode {
	if h1 == nil || h2 == nil {
		return nil
	}

	h1Loop := GetLoopNode(h1)
	h2Loop := GetLoopNode(h2)

	// both have no loop
	if h1Loop == nil && h2Loop == nil {
		return noLoop(h1, h2)
	}

	// both have loop
	if h1Loop != nil && h2Loop != nil {
		return bothLoop(h1, h2, h1Loop, h2Loop)
	}

	// one with loop, one without loop, they must doesn't have intersection
	return nil
}

// both head1 and head2 have no loop
func noLoop(h1, h2 *utils.ListNode) *utils.ListNode {
	p1, p2 := h1, h2
	var gap int
	for p1 != nil {
		gap++
		p1 = p1.Next
	}

	for p2 != nil {
		gap--
		p2 = p2.Next
	}

	// see if h1 and h2 have same end node
	if p1 != p2 {
		return nil
	}

	// find the longer list, set p1 as longer one
	if gap < 0 {
		p1, p2 = h2, h1
	} else {
		p1, p2 = h1, h2
	}

	// get abs of gap
	if gap < 0 {
		gap = -gap
	}

	// p1 go gap step first
	for gap != 0 {
		p1 = p1.Next
		gap -= 1
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

func bothLoop(h1, h2, h1Loop, h2Loop *utils.ListNode) *utils.ListNode {
	// 1. no intersection
	// 2. same loop enter point
	// 3. different loop enter point

	if h1Loop != h2Loop {
		p := h1Loop.Next
		for p != h1Loop {
			// situation 3
			if p == h2Loop {
				return h2Loop // return h1Loop or h2Loop, both are intersection
			}
			p = p.Next
		}
		// situation 1
		return nil
	}

	// situation 2
	p1, p2 := h1, h2
	var gap int
	for p1 != h1Loop {
		gap++
		p1 = p1.Next
	}

	for p2 != h1Loop {
		gap--
		p2 = p2.Next
	}

	// find the longer list, set p1 as longer one
	if gap < 0 {
		p1, p2 = h2, h1
	} else {
		p1, p2 = h1, h2
	}

	// get abs of gap
	if gap < 0 {
		gap = -gap
	}

	// p1 go gap step first
	for gap != 0 {
		p1 = p1.Next
		gap -= 1
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
