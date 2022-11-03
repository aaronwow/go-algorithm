package utils

import (
	"fmt"
	"testing"
)

func TestGenerateRandomInts(t *testing.T) {
	ints := GenerateRandomInts(101)
	fmt.Println(ints)
}
