package algorithm

func isAnagram(s string, t string) bool {
	seenMap := make(map[rune]int, 0)

	for _, c := range s {
		seenMap[c]++
	}

	for _, c := range t {
		seenMap[c]--
		if seenMap[c] == 0 {
			delete(seenMap, c)
		}
	}

	return len(seenMap) == 0
}
