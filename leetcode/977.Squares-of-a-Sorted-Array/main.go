package algorithm

func sortedSquares(nums []int) []int {
	size := len(nums)
	ans := make([]int, size)

	for i, j, k := 0, size-1, size-1; i <= j; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			ans[k] = nums[i] * nums[i]
			i++
		} else {
			ans[k] = nums[j] * nums[j]
			j--
		}
	}
	return ans
}
