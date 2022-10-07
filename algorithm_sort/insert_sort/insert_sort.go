package algorithm

import "gandi.icu/go-algorithm/utils"

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			utils.SwapInt(arr, j, j-1)
		}
	}
}
