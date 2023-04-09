package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return len(nums)
	}
	validIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[validIndex] {
			validIndex++
			nums[validIndex] = nums[i]
		}
	}

	return validIndex
}
