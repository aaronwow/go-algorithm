// a bit different from leetcode 662, null node are not include
package algorithm

import (
	"gandi.icu/go-algorithm/utils"
)

func MaximumWidthOfBinaryTree(head *utils.BinaryTreeNode) int {
	if head == nil {
		return 0
	}

	queue, maxWidth := []*utils.BinaryTreeNode{}, 0
	queue = append(queue, head)
	for len(queue) > 0 {
		qLen := len(queue)
		if qLen > maxWidth {
			maxWidth = qLen
		}
		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return maxWidth
}
