package utils

import (
	"math/rand"
)

func GenerateRandomInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(100)
	}
	return ints
}
