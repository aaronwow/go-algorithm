package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestInsertSort(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	InsertSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("InsertSort failed")
	}
}
