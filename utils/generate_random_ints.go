package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func GenerateRandomInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(100)
	}
	return ints
}

func GenerateRandomInt(max int) int {
	return rand.Intn(max)
}
