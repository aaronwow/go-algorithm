package algorithm

import (
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestMaximumWidthOfBinaryTree(t *testing.T) {
	// create a root node
	head := &utils.BinaryTreeNode{Val: 0}
	// level 2
	head.Left = &utils.BinaryTreeNode{Val: 0}
	head.Right = &utils.BinaryTreeNode{Val: 0}
	// level 3
	head.Left.Left = &utils.BinaryTreeNode{Val: 0}
	head.Left.Right = &utils.BinaryTreeNode{Val: 0}
	head.Right.Left = &utils.BinaryTreeNode{Val: 0}
	head.Right.Right = &utils.BinaryTreeNode{Val: 0}
	// level 4
	head.Left.Left.Left = &utils.BinaryTreeNode{Val: 0}
	head.Left.Left.Right = &utils.BinaryTreeNode{Val: 0}
	head.Right.Left.Left = &utils.BinaryTreeNode{Val: 0}
	// level 5
	head.Left.Left.Left.Left = &utils.BinaryTreeNode{Val: 0}
	head.Left.Left.Left.Right = &utils.BinaryTreeNode{Val: 0}
	head.Left.Left.Right.Left = &utils.BinaryTreeNode{Val: 0}
	head.Left.Left.Right.Right = &utils.BinaryTreeNode{Val: 0}
	head.Right.Left.Left.Left = &utils.BinaryTreeNode{Val: 0}
	head.Right.Left.Left.Right = &utils.BinaryTreeNode{Val: 0}

	if MaximumWidthOfBinaryTree(head) != 6 {
		t.Errorf("Failed: maximum width of this tree is 6")
	}
}
