package algorithm

func twoSum(nums []int, target int) []int {
	cache := map[int]int{}

	for i, v := range nums {
		if _, ok := cache[target-v]; ok {
			return []int{cache[target-v], i}
		}
		cache[v] = i
	}
	return []int{}
}
