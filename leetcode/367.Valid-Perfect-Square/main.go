package algorithm

func isPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)>>2
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}
