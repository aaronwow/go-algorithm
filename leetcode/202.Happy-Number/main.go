package algorithm

func isHappy(n int) bool {
	cache := map[int]bool{}

	for {
		sum := calc(n)
		if sum == 1 {
			return true
		}

		if cache[sum] {
			return false
		}

		cache[sum] = true
		n = sum
	}

}

func calc(num int) int {
	sum := 0
	for num != 0 {
		cur := num % 10
		sum += cur * cur
		num /= 10
	}
	return sum
}
