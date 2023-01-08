package utils

import "errors"

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTreeNodeStack []*BinaryTreeNode

func (s *BinaryTreeNodeStack) Push(v *BinaryTreeNode) {
	*s = append(*s, v)
}

func (s *BinaryTreeNodeStack) Pop() (*BinaryTreeNode, error) {
	if len(*s) == 0 {
		return nil, errors.New("stack is empty")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}
