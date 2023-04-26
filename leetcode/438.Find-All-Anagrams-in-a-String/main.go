package algorithm

func findAnagrams(s string, p string) []int {
	needs, window := make(map[byte]int), make(map[byte]int)
	res := []int{}
	for i := 0; i < len(p); i++ {
		needs[p[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s) {
		cur := s[right]
		right++

		if _, ok := needs[cur]; ok {
			window[cur]++
			if window[cur] == needs[cur] {
				valid++
			}
		}

		if right-left >= len(p) {
			if valid == len(needs) {
				res = append(res, left)
			}

			d := s[left]
			left++

			if _, ok := needs[cur]; ok {
				if window[d] == needs[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
