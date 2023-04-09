package algorithm

func moveZeroes(nums []int) []int {
	validIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[validIndex], nums[i] = nums[i], nums[validIndex]
			validIndex++
		}
	}

	return nums
}
