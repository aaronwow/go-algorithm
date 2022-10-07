package algorithm

import "gandi.icu/go-algorithm/utils"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				utils.SwapInt(arr, j, i)
			}
		}
	}
}
