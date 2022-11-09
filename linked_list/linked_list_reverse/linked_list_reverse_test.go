package algorithm

import (
	"fmt"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestLinkedListReverse(t *testing.T) {
	head := utils.GenerateRandomLinkedList(10)
	head = LinkedListReverse(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
