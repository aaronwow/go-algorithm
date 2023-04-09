package algorithm

func searchRange(nums []int, target int) []int {
	low := 0
	high := len(nums) - 1

	startPoint := lowerBound(nums, low, high, target)
	if startPoint == len(nums) || nums[startPoint] != target {
		return []int{-1, -1}
	}
	endPoint := lowerBound(nums, low, high, target+1) - 1

	return []int{startPoint, endPoint}
}

func lowerBound(nums []int, lo int, hi int, target int) int {
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return lo
}
