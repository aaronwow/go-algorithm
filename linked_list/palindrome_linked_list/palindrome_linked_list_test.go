package algorithm

import (
	"fmt"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestPalindromeLinkedList(t *testing.T) {
	// generate odd palindrome linked list
	oddPalindrome := &utils.ListNode{Val: 1}
	oddPalindrome.Next = &utils.ListNode{Val: 2}
	oddPalindrome.Next.Next = &utils.ListNode{Val: 3}
	oddPalindrome.Next.Next.Next = &utils.ListNode{Val: 2}
	oddPalindrome.Next.Next.Next.Next = &utils.ListNode{Val: 1}
	fmt.Println(PalindromeLinkedList(oddPalindrome))
	// generate even palindrome linked list
	evenPalindrome := &utils.ListNode{Val: 1}
	evenPalindrome.Next = &utils.ListNode{Val: 2}
	evenPalindrome.Next.Next = &utils.ListNode{Val: 3}
	evenPalindrome.Next.Next.Next = &utils.ListNode{Val: 3}
	evenPalindrome.Next.Next.Next.Next = &utils.ListNode{Val: 2}
	evenPalindrome.Next.Next.Next.Next.Next = &utils.ListNode{Val: 1}
	fmt.Println(PalindromeLinkedList(evenPalindrome))
	// generate non palindrome linked list
	nonPalindrome := utils.GenerateRandomLinkedList(10)
	fmt.Println(PalindromeLinkedList(nonPalindrome))
}
