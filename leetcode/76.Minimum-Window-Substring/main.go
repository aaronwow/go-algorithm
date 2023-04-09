package algorithm

func minWindow(s string, t string) string {
	need := map[byte]int{}
	window := map[byte]int{}

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right, valid := 0, 0, 0
	start, length := 0, len(s)+1

	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			last := s[left]
			left++
			if need[last] > 0 {
				if window[last] == need[last] {
					valid--
				}
				window[last]--
			}
		}
	}

	if length == len(s)+1 {
		return ""
	}

	return s[start : start+length]
}
