package algorithm

import "gandi.icu/go-algorithm/utils"

func ShellSort(arr []int) {
	arrLength := len(arr)
	gap := 1
	// 这里采用Knuth增量O3/2, Sedgewick增量O4/3更好
	for gap < arrLength/3 {
		gap = gap*3 + 1
	}

	for ; gap > 0; gap /= 3 {
		for i := gap; i < arrLength; i++ {
			for j := i; j >= gap && arr[j] < arr[j-gap]; j -= gap {
				utils.SwapInt(arr, j, j-gap)
			}
		}
	}
}
