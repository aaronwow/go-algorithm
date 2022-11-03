package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestQuickSort(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	QuickSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("QuickSort failed")
	}
}

func TestQuickSort3Way(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	QuickSort3Way(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("QuickSort3Way failed")
	}
}
