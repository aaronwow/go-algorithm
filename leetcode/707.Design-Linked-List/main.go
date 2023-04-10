package main

func main() {
	myLinkedList := MyLinkedList{}
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	myLinkedList.Get(1)           // return 2
	myLinkedList.DeleteAtIndex(1) // now the linked list is 1->3
	myLinkedList.Get(1)           // return 3
}

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	cur := this.Head
	for {
		if cur == nil {
			return -1
		}
		if index == 0 {
			return cur.Val
		}
		cur = cur.Next
		index--
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{Val: val}
	newNode.Next = this.Head
	this.Head = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{Val: val}
	if this.Head == nil {
		this.Head = newNode
		return
	}
	cur := this.Head
	for {
		if cur.Next == nil {
			cur.Next = newNode
			return
		}
		cur = cur.Next
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.Head
	for {
		if cur == nil {
			return
		}
		if index == 1 {
			newNode := &Node{Val: val}
			newNode.Next = cur.Next
			cur.Next = newNode
			return
		}
		cur = cur.Next
		index--
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.Head == nil {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	cur := this.Head
	for {
		if cur == nil {
			return
		}
		if index == 1 {
			if cur.Next == nil {
				return
			}
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
		index--
	}
}
