package algorithm

func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	left := 0
	ans := 0

	for right, val := range fruits {
		cnt[val]++
		for len(cnt) > 2 {
			last := fruits[left]
			cnt[last]--
			if cnt[last] == 0 {
				delete(cnt, last)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
