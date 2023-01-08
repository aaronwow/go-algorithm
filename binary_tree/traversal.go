package algorithm

import (
	"fmt"

	"gandi.icu/go-algorithm/utils"
)

func PreOrderUnrecur(head *utils.BinaryTreeNode) {
	if head == nil {
		return
	}
	var res []int
	stack := utils.BinaryTreeNodeStack{}
	stack.Push(head)
	for len(stack) > 0 {
		node, err := stack.Pop()
		if err != nil {
			panic(err)
		}
		res = append(res, node.Val)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	for _, v := range res {
		fmt.Printf("%d ", v)
	}
}

func InOrderUnrecur(head *utils.BinaryTreeNode) {
	if head == nil {
		return
	}
	var res []int
	stack := utils.BinaryTreeNodeStack{}
	for len(stack) > 0 || head != nil {
		if head != nil {
			stack.Push(head)
			head = head.Left
		} else {
			node, err := stack.Pop()
			if err != nil {
				panic(err)
			}
			res = append(res, node.Val)
			head = node.Right
		}
	}

	for _, v := range res {
		fmt.Printf("%d ", v)
	}
}

func PostOrderUnrecur(head *utils.BinaryTreeNode) {
	if head == nil {
		return
	}
	stack1 := utils.BinaryTreeNodeStack{}
	stack2 := utils.BinaryTreeNodeStack{}
	stack1.Push(head)
	for len(stack1) > 0 {
		node, err := stack1.Pop()
		if err != nil {
			panic(err)
		}
		stack2.Push(node)
		if node.Left != nil {
			stack1.Push(node.Left)
		}
		if node.Right != nil {
			stack1.Push(node.Right)
		}
	}
	for len(stack2) > 0 {
		node, err := stack2.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", node.Val)
	}
}

// level order traversal
func LevelOrder(head *utils.BinaryTreeNode) {

}
