package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestSelectSort(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	SelectSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("SelectSort failed")
	}
}
