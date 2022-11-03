package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestRadixSort(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	RadixSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("RadixSort failed")
	}
}
