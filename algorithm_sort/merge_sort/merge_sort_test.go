package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestMergeSortUpToDown(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	MergeSortUpToDown(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("MergeSortUpToDown failed")
	}
}

func TestMergeSortDownToUp(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	MergeSortUpToDown(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("MergeSortDownToUp failed")
	}
}
