package algorithm

func twoSum(nums []int, target int) []int {
	cache := map[int]bool{}

	for _, v := range nums {
		if cache[target-v] {
			return []int{target - v, v}
		}
		cache[v] = true
	}
	return []int{}
}
