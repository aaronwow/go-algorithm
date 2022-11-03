package algorithm

import "math"

func RadixSort(arr []int) {
	process(arr, 0, len(arr)-1, maxbits(arr))
}

func maxbits(arr []int) int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	bits := 0
	for max > 0 {
		max /= 10
		bits++
	}
	return bits
}

func getDigit(x int, d int) int {
	return x / int(math.Pow(10, float64(d-1))) % 10
}

func process(arr []int, l int, r int, digit int) {
	if l >= r {
		return
	}
	const radix = 10
	bucket := make([]int, r-l+1)
	for d := 1; d <= digit; d++ {
		count := [radix]int{}
		for i := l; i <= r; i++ {
			j := getDigit(arr[i], d)
			count[j]++
		}

		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}

		for i := r; i >= l; i-- {
			j := getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}

		for i, j := l, 0; i <= r; i, j = i+1, j+1 {
			arr[i] = bucket[j]
		}
	}

	// if l >= r {
	// 	return
	// }
	// bucket := make([][]int, 10)
	// for i := 0; i < 10; i++ {
	// 	bucket[i] = make([]int, 0)
	// }
	// for i := l; i <= r; i++ {
	// 	bucket[getDigit(arr[i], d)] = append(bucket[getDigit(arr[i], d)], arr[i])
	// }
	// index := l
	// for i := 0; i < 10; i++ {
	// 	for _, v := range bucket[i] {
	// 		arr[index] = v
	// 		index++
	// 	}
	// }
	// for i := 0; i < 10; i++ {
	// 	process(arr, l+i*len(bucket[i]), l+(i+1)*len(bucket[i])-1, d-1)
	// }
}
