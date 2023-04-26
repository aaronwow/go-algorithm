package algorithm

import "sort"

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func groupAnagrams(strs []string) [][]string {
	record := map[string][]string{}
	res := [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}

	for _, v := range record {
		res = append(res, v)
	}

	return res
}
