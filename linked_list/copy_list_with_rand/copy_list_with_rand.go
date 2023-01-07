package algorithm

import "gandi.icu/go-algorithm/utils"

// TODO two approach, with hashmap & without hashmap

func CopyListWithRand(head *utils.ListWithRandNode) *utils.ListWithRandNode {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*utils.ListWithRandNode]*utils.ListWithRandNode)

	cur := head
	for cur != nil {
		nodeMap[cur] = &utils.ListWithRandNode{Val: cur.Val}
		cur = cur.Next
	}

	for cur = head; cur != nil; cur = cur.Next {
		nodeMap[cur].Rand = nodeMap[cur.Rand]
		nodeMap[cur].Next = nodeMap[cur.Next]
	}
	return nodeMap[head]
}

func CopyListWithRandWithoutHashMap(head *utils.ListWithRandNode) *utils.ListWithRandNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		node := &utils.ListWithRandNode{Val: cur.Val, Next: cur.Next}
		cur.Next = node
		cur = node.Next
	}

	// link rand
	cur = head
	for cur != nil {
		if cur.Rand != nil {
			cur.Next.Rand = cur.Rand.Next
		}
		cur = cur.Next.Next
	}

	// restore the original link
	newList := head.Next
	cur = head
	for cur != nil {
		newCur := cur.Next
		cur.Next = newCur.Next
		if cur.Next != nil {
			newCur.Next = cur.Next.Next
		} else {
			newCur.Next = nil
		}
		cur = cur.Next
	}
	return newList
}
