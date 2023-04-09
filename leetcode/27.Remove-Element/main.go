package algorithm

func removeElement(nums []int, val int) int {
	validIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[validIndex] = nums[i]
			validIndex++
		}
	}

	return validIndex
}
