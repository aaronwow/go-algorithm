package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestBubbleSort(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	BubbleSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("BubbleSort failed")
	}
}
