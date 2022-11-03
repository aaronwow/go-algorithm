package algorithm

import (
	"gandi.icu/go-algorithm/utils"
)

func MergeSortUpToDown(arr []int) {
	process(arr, 0, len(arr)-1)
}

func process(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	process(arr, l, mid)
	process(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, lo int, mid int, hi int) {
	i := lo
	j := mid + 1
	res := make([]int, hi-lo+1)
	k := 0

	for i <= mid && j <= hi {
		if arr[i] <= arr[j] {
			res[k] = arr[i]
			i++
		} else {
			res[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		res[k] = arr[i]
		i++
		k++
	}

	for j <= hi {
		res[k] = arr[j]
		j++
		k++
	}

	for i := 0; i < k; i++ {
		arr[lo+i] = res[i]
	}
}

func MergeSortDownToUp(arr []int) {
	l := len(arr)
	for sz := 1; sz < l; sz += sz {
		for lo := 0; lo < l-sz; lo += sz * 2 {
			merge(arr, lo, lo+sz-1, utils.Min(lo+sz+sz-1, l-1))
		}
	}
}
