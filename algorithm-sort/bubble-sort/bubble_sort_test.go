package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestBubbleSort(t *testing.T) {
	ints := utils.GenerateRandomInts(10)
	BubbleSort(ints)
	if !sort.IntsAreSorted(ints) {
		t.Errorf("BubbleSort failed")
	}
}
