package algorithm

import (
	"fmt"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestLinkedListReverse(t *testing.T) {
	head := utils.GenerateRandomDoublyLinkedList(10)
	head = DoublyLinkedListReverse(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
