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

type ListWithRandNode struct {
	Val  int
	Next *ListWithRandNode
	Rand *ListWithRandNode
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

func GenerateRandomListWithRandLinedList(length int) *ListWithRandNode {
	if length <= 0 {
		return nil
	}

	var listArr []*ListWithRandNode
	// create a list
	head := &ListWithRandNode{Val: GenerateRandomInt(100)}
	listArr = append(listArr, head)
	cur := head
	for i := 0; i < length-1; i++ {
		cur.Next = &ListWithRandNode{Val: GenerateRandomInt(100)}
		listArr = append(listArr, cur.Next)
		cur = cur.Next
	}
	// 增加nil值
	listArr = append(listArr, nil)

	// set rand pointer
	cur = head
	for cur != nil {
		index := GenerateRandomInt(length + 1) //last index node is now nil
		randomNode := listArr[index]
		cur.Rand = randomNode
		cur = cur.Next
	}

	return head
}
