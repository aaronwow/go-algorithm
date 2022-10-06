package algorithm

func BubbleSort(ints []int) {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-i-1; j++ {
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
}
