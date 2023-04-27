package algorithm

func reverseStr(s string, k int) string {
	if k > len(s) {
		return reverse(s)
	}

	for i := 0; i < len(s); i += 2 * k {
		if len(s)-i >= k {
			s = s[:i] + reverse(s[i:i+k]) + s[i+k:]
		} else {
			s = s[:i] + reverse(s[i:])
		}
	}

	return s
}

func reverse(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
