package algorithm

import (
	"fmt"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestTraversal(t *testing.T) {
	// create a binary tree
	head := &utils.BinaryTreeNode{Val: 1}
	head.Left = &utils.BinaryTreeNode{Val: 2}
	head.Right = &utils.BinaryTreeNode{Val: 3}
	head.Left.Left = &utils.BinaryTreeNode{Val: 4}
	head.Left.Right = &utils.BinaryTreeNode{Val: 5}
	head.Right.Left = &utils.BinaryTreeNode{Val: 6}
	head.Right.Right = &utils.BinaryTreeNode{Val: 7}

	// test preorder traversal
	PreOrderUnrecur(head)
	fmt.Println()
	// test inorder traversal
	InOrderUnrecur(head)
	fmt.Println()
	// test postorder traversal
	PostOrderUnrecur(head)
	fmt.Println()
}
