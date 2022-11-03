package algorithm

import "gandi.icu/go-algorithm/utils"

func QuickSort(arr []int) {
	process(arr, 0, len(arr)-1)
}

func process(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	j := partition(arr, lo, hi)
	process(arr, lo, j-1)
	process(arr, j+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	i, j := lo+1, hi
	for {
		for arr[i] < arr[lo] && i < hi {
			i++
		}
		for arr[j] > arr[lo] && j > lo {
			j--
		}
		if i >= j {
			break
		}
		utils.SwapInt(arr, i, j)
	}

	utils.SwapInt(arr, lo, j)
	return j
}

func QuickSort3Way(arr []int) {
	process2(arr, 0, len(arr)-1)
}

func process2(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	lt, gt := partition2(arr, lo, hi)
	process2(arr, lo, lt-1)
	process2(arr, gt+1, hi)
}

func partition2(arr []int, lo int, hi int) (int, int) {
	lt, gt := lo, hi
	i := lo + 1
	v := arr[lo]
	for i <= gt {
		if arr[i] < v {
			utils.SwapInt(arr, i, lt)
			i++
			lt++
		} else if arr[i] > v {
			utils.SwapInt(arr, i, gt)
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}
