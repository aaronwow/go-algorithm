package algorithm

func mySqrt(x int) int {
	lo := 1
	hi := x

	for lo <= hi {
		mid := lo + (hi-lo)>>2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return hi
}
