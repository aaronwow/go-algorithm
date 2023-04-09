package algorithm

func minSubArrayLen(target int, nums []int) int {
	fast, slow := 0, 0
	minLength := len(nums) + 1
	sum := 0
	for fast < len(nums) {
		sum += nums[fast]
		for sum >= target {
			minLength = min(minLength, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
		fast++
	}

	if minLength == len(nums)+1 {
		return 0
	}

	return minLength
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
