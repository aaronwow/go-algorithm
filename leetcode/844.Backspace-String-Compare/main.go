package algorithm

func backspaceCompare(s string, t string) bool {
	s1 := make([]rune, 0)
	for _, v := range s {
		if v == '#' {
			if len(s1) > 0 {
				s1 = s1[:len(s1)-1]
			}
		} else {
			s1 = append(s1, v)
		}
	}

	s2 := make([]rune, 0)
	for _, v := range t {
		if v == '#' {
			if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}
		} else {
			s2 = append(s2, v)
		}
	}

	return string(s1) == string(s2)
}
