package algorithm

import (
	"fmt"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestPartitionList(t *testing.T) {
	testList := &utils.ListNode{Val: 1}
	testList.Next = &utils.ListNode{Val: 4}
	testList.Next.Next = &utils.ListNode{Val: 3}
	testList.Next.Next.Next = &utils.ListNode{Val: 2}
	testList.Next.Next.Next.Next = &utils.ListNode{Val: 5}
	testList.Next.Next.Next.Next.Next = &utils.ListNode{Val: 2}

	res := PartitionList(testList, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
