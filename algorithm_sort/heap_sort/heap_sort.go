package algorithm

import (
	"gandi.icu/go-algorithm/utils"
)

func HeapSortMax(arr []int) {
	l := len(arr)
	// for i := 1; i < l; i++ {
	// 	swim(arr, i) // logN
	// }
	// 优化initial过程
	for i := l - 1; i >= 0; i-- { // N
		sink(arr, i, l)
	}

	for i := 0; i < l-1; i++ { // N
		sink(arr, 0, l-i)            // logN
		utils.SwapInt(arr, 0, l-i-1) // 1
	}
}

func swim(arr []int, i int) {
	for i > 1 && arr[i] > arr[(i-1)/2] {
		utils.SwapInt(arr, i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func sink(arr []int, i int, l int) {
	// 左子节点存在
	for 2*i+1 < l {
		j := 2*i + 1
		// 右子节点存在且大于左子节点
		if j+1 < l && arr[j] < arr[j+1] {
			j++
		}
		if arr[i] > arr[j] {
			break
		}
		utils.SwapInt(arr, i, j)
		i = j
	}
}
