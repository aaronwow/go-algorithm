package algorithm

import "gandi.icu/go-algorithm/utils"

// stack
func PalindromeLinkedList(head *utils.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	n1, n2 := head, head
	// 偶数慢指针在左侧, n1慢指针获取中间位置
	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	n2 = n1.Next
	n1.Next = nil
	var n3 *utils.ListNode
	for n2 != nil {
		n3 = n2.Next
		n2.Next = n1
		n1 = n2
		n2 = n3
	}

	n3 = n1 // last node
	n2 = head
	res := true // 因为需要恢复链表所以没有直接返回
	for n1 != nil && n2 != nil {
		if n1.Val != n2.Val {
			res = false
			break
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	n1 = n3.Next
	n3.Next = nil
	for n1 != nil {
		n2 = n1.Next
		n1.Next = n3
		n3 = n1
		n1 = n2
	}
	return res
}
