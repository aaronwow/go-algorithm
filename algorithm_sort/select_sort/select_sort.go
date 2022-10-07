package algorithm

import "gandi.icu/go-algorithm/utils"

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if min != i {
			utils.SwapInt(arr, min, i)
		}
	}
}
