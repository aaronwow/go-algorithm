package utils

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestGenerateRandomInts(t *testing.T) {
	ints := GenerateRandomInts(101)
	fmt.Println(ints)
}
