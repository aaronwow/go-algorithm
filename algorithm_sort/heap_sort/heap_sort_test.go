package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestHeapSortMax(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	HeapSortMax(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("HeapSortMax failed")
	}
}
