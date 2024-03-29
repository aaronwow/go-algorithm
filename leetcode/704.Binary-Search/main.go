package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
