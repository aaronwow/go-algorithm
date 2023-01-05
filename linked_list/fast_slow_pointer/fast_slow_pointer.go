package main

import (
	"fmt"

	"gandi.icu/go-algorithm/utils"
)

func main() {
	// create a even list 1 2 3 4
	evenList := utils.GenerateRandomLinkedList(4)
	// create a odd list 1 2 3 4 5
	oddList := utils.GenerateRandomLinkedList(5)

	// 奇数中间，偶数前一个
	fast, slow := evenList, evenList
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	for fast.Next != nil {
		fast = fast.Next
	}
	fmt.Println(fast.Val, slow.Val)
	fast, slow = oddList, oddList
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	for fast.Next != nil {
		fast = fast.Next
	}
	fmt.Println(fast.Val, slow.Val)

	// 奇数中间，偶数后一个
	fast, slow = evenList, evenList
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fmt.Println(fast.Val, slow.Val)
	fast, slow = oddList, oddList
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fmt.Println(fast.Val, slow.Val)

	// 奇数中间前一个，偶数前一个再前一个
	fast, slow = evenList, evenList
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fmt.Println(fast.Val, slow.Val)
	fast, slow = oddList, oddList
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fmt.Println(fast.Val, slow.Val)
}
